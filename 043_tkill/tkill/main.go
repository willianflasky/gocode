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
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.pri.ibanyu.com/devops/tkill/cmd"
	"gitlab.pri.ibanyu.com/devops/tkill/logger"
)

var path, _ = filepath.Abs(os.Args[0])
var base_dir = filepath.Dir(path)

func main() {
	logspath := filepath.Join(base_dir, "logs")
	if _, err := os.Stat(logspath); os.IsNotExist(err) {
		os.MkdirAll(logspath, os.ModePerm)
	}

	if err := logger.Init(base_dir); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
    
	cmd.Execute()
}
