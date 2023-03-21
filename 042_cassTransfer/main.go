package main

import (
	"cassTransfer/logic"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	logic.InitLog("./log/info.log", "./log/error.log", zap.InfoLevel)

	var tableSuffix string
	var wg sync.WaitGroup
	wg = sync.WaitGroup{}

	sum, _ := logic.GetTablenum("ClusterInfo")
	fmt.Println("总表数", sum)

	for i := 0; i < sum; i++ {
		tableSuffix = fmt.Sprintf("ClusterInfo_%04d", i)
		i += 1
		wg.Add(1)

		logic.NewClusterInfo().GetClusterInfo(tableSuffix, &wg)

		if i%5 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	wg.Wait()
}
