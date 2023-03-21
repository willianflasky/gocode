package logic

import (
	"cassTransfer/db"
)

func GetTablenum(tableNamePattern string) (int, error) {
	sqlSelect := "SELECT COUNT(*) FROM information_schema.tables WHERE table_name LIKE ?"
	//fmt.Println(sqlSelect)
	connection, err := db.GetMysqlConn(true)
	if err != nil {
		Logger.Errorf("获取数据链接：%v", err)
		return 0, err
	}

	defer connection.Close()

	stmt, err := connection.Prepare(sqlSelect)
	if err != nil {
		Logger.Errorf("prepare SQL失败: %v", err)
		return 0, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(tableNamePattern + "%").Scan(&count)
	if err != nil {
		Logger.Errorf("执行sql失败: %v", err)
		return 0, err
	}

	return count, nil

}
