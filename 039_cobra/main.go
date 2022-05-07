package main

import (
	"9981/rootcmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 注册命令
func init() {
	rootcmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Short: "user",
	Use:   "user ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("./main user ", args)
	},
}

func main() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
