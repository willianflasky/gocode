package logic

import (
	"cassTransfer/db"
	"errors"
)

type CassMsgInfo struct {
	ClusterId   int64
	ClusterType int
	MsgBucketId int64
	MsgId       int64
	MsgCmd      int64
	MsgData     []byte
	ReplyMsgId  int64
	SourceUid   int64
	TargetUids  []byte
	TimeTag     int64
	TimeFormat  int64
}

var cassMsgInfo CassMsgInfo

func NewCassMsgInfo() *CassMsgInfo {
	cassMsgInfo := &CassMsgInfo{}
	return cassMsgInfo
}

func (cm *CassMsgInfo) GetMsgFromCass(ci []ClusterInfo, timeToday, timeYesterday int64) (cmi []CassMsgInfo, err error) {
	session, err := db.GetCassandraConn()
	if err != nil {
		return nil, err
	}

	var cassMsgInfoList []CassMsgInfo
	var minTime, maxTime int64
	for _, v := range ci {

		bucketId := v.TodayLastMsgId >> 17 //这个是今天从MySQL取的LastMsgId
		yesterdayBucketId := bucketId - 1
		if yesterdayBucketId < 0 {
			yesterdayBucketId = 0
		}
		minTime = timeYesterday << 24
		maxTime = timeToday << 24

		for tmpBucketId := int64(yesterdayBucketId); tmpBucketId <= bucketId; tmpBucketId++ {
			sqlStr := "select * from xx.xx where clusterid=?  and clustertype=? and msgbucketid=? and msgid<=? and timetag>? and timetag<= ? and msgcmd=50724904 ALLOW FILTERING"
			iter := session.Query(sqlStr, v.ClusterId, v.Type, tmpBucketId, v.TodayLastMsgId, minTime, maxTime).Iter().Scanner()

			for iter.Next() {
				var cassMsgInfo CassMsgInfo
				iter.Scan(&cassMsgInfo.ClusterId, &cassMsgInfo.ClusterType, &cassMsgInfo.MsgBucketId, &cassMsgInfo.MsgId, &cassMsgInfo.MsgCmd,
					&cassMsgInfo.MsgData, &cassMsgInfo.ReplyMsgId, &cassMsgInfo.SourceUid, &cassMsgInfo.TargetUids, &cassMsgInfo.TimeTag)
				cassMsgInfo.TimeFormat = cassMsgInfo.TimeTag >> 24
				cassMsgInfoList = append(cassMsgInfoList, cassMsgInfo)
			}
		}
	}
	//for i, info := range cassMsgInfoList {
	//	intolistparsedmsg, _ := UnPackAndGetMessageData(info.MsgData)
	//	fmt.Println("end: ", i, info.ClusterId, info.MsgId, intolistparsedmsg)
	//}

	defer session.Close()
	return cassMsgInfoList, nil
}

func UnPackAndGetMessageData(msgData []byte) (string, error) {
	if len(msgData) <= 12 {
		return "", errors.New("invalid input")
	}
	c := msgData[12:]
	pos, err := findPos(c, "\x00")
	if pos == 0 {
		return "", errors.New("can not found message data")
	}
	if err != nil {
		return "", err
	}
	c = c[:pos]
	chatMessage := string(c)
	//fmt.Println(chatMessage, len(chatMessage))
	return chatMessage, nil
}

func findPos(byes []byte, f string) (int, error) {
	for pos, byt := range byes {
		if string(byt) == f {
			return pos, nil
		}
	}
	return 0, errors.New("pos not find")
}
