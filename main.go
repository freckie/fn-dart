package main

import (
	"fmt"
	"os"
	"strings"

	"fn-dart/config"
	"fn-dart/crawlers"
	"fn-dart/engine"
	"fn-dart/models"
	"fn-dart/utils"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	data, err := utils.GetRecentReports(*cfg, "", "1", "100")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var c crawlers.Crawler
	cid1 := crawlers.CID1{}
	cid2 := crawlers.CID2{}
	cid3 := crawlers.CID3{}
	cid4 := crawlers.CID4{}

	tg := engine.TGEngine{}
	tg.Cfg = cfg
	err = tg.GenerateBot()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for _, item := range data {
		report := models.Report{}
		report.Title = item.ReportNM
		report.RceptNo = item.RceptNo
		report.CorpName = item.CorpName

		if strings.Contains(item.ReportNM, "[기재정정]") {
			continue
		}

		if cid1.IsTarget(item.ReportNM) {
			report.CrawlerID = "1"
			c = cid1
		} else if cid2.IsTarget(item.ReportNM) {
			report.CrawlerID = "2"
			c = cid2
		} else if cid3.IsTarget(item.ReportNM) {
			report.CrawlerID = "3"
			c = cid3
		} else if cid4.IsTarget(item.ReportNM) {
			report.CrawlerID = "4"
			c = cid4
		} else {
			continue
		}

		err := c.GetDetail(&report)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = tg.SendMessage(report)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(report)
	}
}
