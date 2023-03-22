/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.pri.ibanyu.com/devops/tkill/imp"
)

var cfgFile string
var killContext = imp.NewKillContext()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "tkill -l 10.106.68.xx [dry_run|execute]\n" +
		"  tkill -l 10.106.68.xx,10.106.70.xx -d eo_osfile [dry_run|execute]",
	Short: "kill tidb long query..",
	Long:  `kill tidb long query..`,
	Args:  cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("hello world..")
		killContext.Operator(args)
		/*		if err != nil {
				fmt.Println(err)
			}*/
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&killContext.Host, "host", "l", "", "host_list")
	rootCmd.MarkFlagRequired("host")
	rootCmd.Flags().StringVarP(&killContext.DbName, "db_name", "d", "", "db_name")
	rootCmd.Flags().IntVarP(&killContext.ExecTime, "query_time", "q", 5, "query_time")
	rootCmd.Flags().IntVarP(&killContext.Port, "db_port", "p", 4000, "db_port")
	rootCmd.Flags().IntVarP(&killContext.ConnCount, "conn_count", "c", 1, "conn_count")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".tkill" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tkill")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
