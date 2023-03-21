package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// 生产
const (
	P_USERNAME    string = ""
	P_PASSWORD    string = ""
	P_DBNAME_READ string = ""
	P_DBHOST_READ string = ""
	P_DBPORT_READ int    = 6021

	P_DBHOST_WR    string = ""
	P_DBNAME_WR    string = ""
	P_TABLENAME_WR string = ""
	P_DBPORT_WR    int    = 61106
)

var proCassHosts = []string{""}

// 测试
const (
	D_USERNAME  string = ""
	D_PASSWORD  string = ""
	D_DBNAME    string = ""
	D_TABLENAME string = ""
	D_DBHOST    string = ""
	D_DBPORT    int    = 61106
)

var devCassHosts = []string{}

// 公共参数
const (
	NETWORK    string = ""
	CASSUSER   string = ""
	CASSPASSWD string = ""
	KEYSPACE   string = ""
	DAY        int    = 0
)

// 环境设置
var Env = "pro"

func GetCassandraConn() (*gocql.Session, error) {
	var hosts []string
	if Env == "dev" {
		hosts = devCassHosts
	} else {
		hosts = proCassHosts
	}

	cluster := gocql.NewCluster()
	cluster.Keyspace = KEYSPACE
	cluster.Hosts = hosts
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: CASSUSER,
		Password: CASSPASSWD,
	}
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = 3 * time.Duration(time.Second)
	cluster.Timeout = 10 * time.Duration(time.Second)
	cluster.Consistency = gocql.Quorum
	cluster.NumConns = 10

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetMysqlConn(read bool) (*sql.DB, error) {
	var connInfo string
	if Env == "dev" {
		connInfo = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=3s&charset=utf8mb4&parseTime=true", D_USERNAME, D_PASSWORD, NETWORK, D_DBHOST, D_DBPORT, D_DBNAME)
		fmt.Println(connInfo)
	} else if read && Env == "pro" {
		connInfo = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=3s&charset=utf8mb4&parseTime=true", P_USERNAME, P_PASSWORD, NETWORK, P_DBHOST_READ, P_DBPORT_READ, P_DBNAME_READ)
	} else {
		connInfo = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=3s&charset=utf8mb4&parseTime=true", P_USERNAME, P_PASSWORD, NETWORK, P_DBHOST_WR, P_DBPORT_WR, P_DBNAME_WR)
	}

	conn, err := sql.Open("mysql", connInfo)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
