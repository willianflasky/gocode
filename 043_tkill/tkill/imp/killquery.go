package imp

//import "C"
import (
	"database/sql"
	"log"

	//"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME string = "tkill"
	PASSWORD string = "xxxx"
	SYSTEMDB string = "information_schema"
	NETWORK  string = "tcp"
)

type KillContext struct {
	Host      string
	Port      int
	DbName    string
	TableName string
	ConnCount int
	ExecTime  int
	Match     string
}

type QueryInfo struct {
	ID          int
	DB          string
	ExecuteTime int
	Info        string
}

func (k *KillContext) GetConn(host string) (*sql.DB, error) {
	serverInfo := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=2s&readTimeout=3s&charset=utf8", USERNAME, PASSWORD, NETWORK, host, k.Port, SYSTEMDB)
	conn, err := sql.Open("mysql", serverInfo)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (k *KillContext) Operator(args []string) {
	var hostStr string
	var ips []string

	hostStr = k.Host
	matchComma, _ := regexp.MatchString(",", hostStr)
	if matchComma {
		separator := ","
		ips = strings.Split(hostStr, separator)
		for _, ip := range ips {
			pattern := "[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+"
			matchIp, _ := regexp.MatchString(pattern, ip)
			if !matchIp {
				fmt.Println("ip list not meet requirement, [10.10.1.1 or 10.10.1.1,10.10.2.2]")
			}
		}
	} else {
		pattern := "[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]"
		matchIp, _ := regexp.MatchString(pattern, hostStr)
		if !matchIp {
			fmt.Println("ip not meet requirement, [10.xx.xx.xx]")
		}
		ips = append(ips, hostStr)
	}

	if len(args) == 1 && args[0] == "dry_run" {
		wg := sync.WaitGroup{}
		wg.Add(len(ips))

		for _, host := range ips {
			go func(h string) {
				results, err := k.GetLongQuery(h)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Printf("==> Query ip: %s | Query sum: %d\n", h, len(results))
				if len(results) == 0 {
					fmt.Println("No query to list")
				}

				for _, v := range results {
					fmt.Printf("ip: %s, %+v\n", h, v)
				}
				wg.Done()
			}(host)
		}
		wg.Wait()
	} else if len(args) == 1 && args[0] == "execute" {
		wg := sync.WaitGroup{}
		wg.Add(len(ips))

		for _, host := range ips {
			go func(h string) {
				results, err := k.GetLongQuery(h)
				if err != nil {
					fmt.Printf("%v\n", err)
				}

				if len(results) > 0 {
					fmt.Printf("==> Kill ip: %s,  kill query sum:%d\n", h, len(results))
					err := k.KillLongQuery(results, h)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Printf("==> Kill ip: %s\n", h)
					fmt.Println("No query to kill")
				}
				wg.Done()
			}(host)
		}
		wg.Wait()
	} else {
		fmt.Println("Usage:")
		fmt.Print("  tkill -l 10.106.68.xx [dry_run|execute]\n")
		fmt.Println("  tkill -l 10.106.68.xx,10.106.70.xx -d rtc [dry_run|execute] [flags]")
	}
}

func (k *KillContext) GetLongQuery(host string) ([]QueryInfo, error) {
	var sqlStr string

	connection, err := k.GetConn(host)
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	var rows *sql.Rows
	if k.DbName != "" {
		sqlStr = "select id, db, time, info from information_schema.processlist where db = ? and time >= ? " +
			"and info is not null and (info like '%select%' or info like '%SELECT%') and db not in ('INFORMATION_SCHEMA', 'PERFORMANCE_SCHEMA') order by time desc"
		stmt, err := connection.Prepare(sqlStr)
        // log.Println(sqlStr)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		rows, err = stmt.Query(k.DbName, k.ExecTime)
		if err != nil {
			return nil, err
		}
	} else {
		sqlStr = "select id, db, time, info from information_schema.processlist where time >= ? " +
			"and info is not null and (info like '%select%' or info like '%SELECT%') and db not in ('INFORMATION_SCHEMA', 'PERFORMANCE_SCHEMA') order by time desc"
		stmt, err := connection.Prepare(sqlStr)
        // log.Println(sqlStr)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err = stmt.Query(k.ExecTime)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	var results []QueryInfo
	for rows.Next() {
		var q QueryInfo
		err := rows.Scan(&q.ID, &q.DB, &q.ExecuteTime, &q.Info)
		if err != nil {
			fmt.Printf("debug: %s", "xx")
			return nil, err
		}
		results = append(results, q)
	}

	return results, err
}

func NewKillContext() *KillContext {
	return &KillContext{}
}

func (k *KillContext) KillLongQuery(queryList []QueryInfo, host string) (err error) {
	connection, err := k.GetConn(host)
	if err != nil {
		return err
	}
	defer connection.Close()

	wg := sync.WaitGroup{}
	wg.Add(len(queryList))
	for _, ql := range queryList {
		go func(q QueryInfo) error {
			log.Printf("IP: %s\tID: %d\tDB: %s\tExecuteTime: %d\tSQL: %s\n", host, q.ID, q.DB, q.ExecuteTime, q.Info)
			log.Printf("kill tidb %d;", q.ID)
			cmd := fmt.Sprintf("kill tidb %d;", q.ID)
			_, err := connection.Exec(cmd)
			if err != nil {
				wg.Done()
				return err
			}
			wg.Done()
			return nil
		}(ql)
	}
	wg.Wait()

	return nil
}
