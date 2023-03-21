package logic

import (
	"cassTransfer/db"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

type ClusterInfo struct {
	ClusterId       int64
	Type            int
	TodayLastMsgId  int64
	TodayUpdateTime time.Time
}

var clusterInfo ClusterInfo

func NewClusterInfo() *ClusterInfo {
	clusterInfo := &ClusterInfo{}
	return clusterInfo
}

func (c ClusterInfo) GetClusterInfo(tableName string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	yesterday := time.Now().AddDate(0, 0, db.DAY-1)
	toYesterday := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location()).Unix()

	//mysql
	mysqlTimeYesterdayFormat := time.Unix(toYesterday, 0).Format("2006-01-02 00:00:00")
	Logger.Infof("从表：%s, 开始查询大于时间：%v 一天内的LastMsgId", tableName, mysqlTimeYesterdayFormat)

	tt := time.Now().AddDate(0, 0, db.DAY)
	cassandraTToday := time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, tt.Location())
	//cassandra
	cassandraTimeToday := cassandraTToday.Unix()

	//整型毫秒时间戳，当天-1
	yt := time.Now().AddDate(0, 0, db.DAY-1)
	cassandraTYesterday := time.Date(yt.Year(), yt.Month(), yt.Day(), 0, 0, 0, 0, yt.Location())
	//cassandra
	cassandraTimeYesterday := cassandraTYesterday.Unix()

	// 从数据库查询最近的LastMsgId
	var rows *sql.Rows
	var results []ClusterInfo
	var sqlSelect string

	if db.Env == "dev" {
		sqlSelect = fmt.Sprintf("select ClusterID, Type, LastMsgID, UpdateTime from %s.%s where UpdateTime>=?", db.D_DBNAME, tableName)
	} else {
		sqlSelect = fmt.Sprintf("select distinct(ClusterID), Type, LastMsgID, UpdateTime from %s.%s where UpdateTime>=? ", db.P_DBNAME_READ, tableName)
	}

	connection, err := db.GetMysqlConn(true)
	if err != nil {
		Logger.Errorf("获取数据链接：%v", err)
		return err
	}

	stmt, err := connection.Prepare(sqlSelect)
	if err != nil {
		Logger.Errorf("查询表：%s, %v", tableName, err)
		return err
	}

	rows, err = stmt.Query(mysqlTimeYesterdayFormat)
	if err != nil {
		Logger.Errorf("查询表：%s, %v", tableName, err)
		return err
	}

	for rows.Next() {
		err := rows.Scan(&c.ClusterId, &c.Type, &c.TodayLastMsgId, &c.TodayUpdateTime)
		if err != nil {
			return err
		}
		results = append(results, c)
	}
	Logger.Infof("从表：%s, 查询到新增的集群id数：%d", tableName, len(results))
	if len(results) == 0 {
		return nil
	}

	Logger.Infof("从表：eo_chatmsg, 查询新增的消息")
	cmi, err := NewCassMsgInfo().GetMsgFromCass(results, cassandraTimeToday, cassandraTimeYesterday)
	if err != nil {
		Logger.Errorf("从表：%s 获取数据：%v", "eo_chatmsg", err)
		return err
	}

	Logger.Infof("写表：user_chat_msg, 写入新增的消息")
	err = WriteMsg(cmi)
	if err != nil {
		Logger.Errorf("写表：user_chat_msg, %v", err)
		return err
	}

	defer stmt.Close()
	defer connection.Close()

	return nil
}

func WriteMsg(cmi []CassMsgInfo) error {
	var connInfo, sqlInsert string
	if db.Env == "dev" {
		connInfo = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=3s&readTimeout=20s&writeTimeout=10s&charset=utf8mb4&parseTime=true", db.D_USERNAME, db.D_PASSWORD, db.NETWORK, db.D_DBHOST, db.D_DBPORT, db.D_DBNAME)
		sqlInsert = fmt.Sprintf("replace into %s.%s(clusterid, clustertype, msgbucketid, msgid, msgcmd, msgdata, replymsgid, sourceuid, targetuids, timetag, timeformat)"+
			"values (?,?,?,?,?,?,?,?,?,?,?)", db.D_DBNAME, db.D_TABLENAME)
	} else {
		connInfo = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=3s&charset=utf8mb4&parseTime=true", db.P_USERNAME, db.P_PASSWORD, db.NETWORK, db.P_DBHOST_WR, db.P_DBPORT_WR, db.P_DBNAME_WR)
		sqlInsert = fmt.Sprintf("replace into %s.%s(clusterid, clustertype, msgbucketid, msgid, msgcmd, msgdata, replymsgid, sourceuid, targetuids, timetag, timeformat)"+
			"values (?,?,?,?,?,?,?,?,?,?,?)", db.P_DBNAME_WR, db.P_TABLENAME_WR)
	}

	conn, err := sql.Open("mysql", connInfo)
	if err != nil {
		return err
	}

	stmt, err := conn.Prepare(sqlInsert)
	if err != nil {
		return err
	}

	for _, v := range cmi {
		chatMessage, err := UnPackAndGetMessageData(v.MsgData)

		if err != nil {
			Logger.Errorf("MsgData: %v 数据解析报错：%v", v.MsgData, err)
			continue
		}

		_, err = stmt.Exec(v.ClusterId, v.ClusterType, v.MsgBucketId, v.MsgId, v.MsgCmd, chatMessage, v.ReplyMsgId,
			v.SourceUid, v.TargetUids, v.TimeTag, v.TimeFormat)
		if err != nil {
			Logger.Errorf("数据写入数据库报错：%v", err)
			continue
		}

	}

	defer stmt.Close()
	defer conn.Close()
	return nil
}
