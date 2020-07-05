package main

import (
	"fmt"
	"os"

	"fn-dart/config"
	"fn-dart/utils"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	data, err := utils.GetRecentReports(cfg, "20200703", "1", "30")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(data)
}
