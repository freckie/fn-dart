package engine

import (
	"fn-dart/config"
	"fn-dart/crawlers"
	"fn-dart/models"
	"fn-dart/utils"
	"log"
	"strings"
	"time"
)

// Engine : Crawling Main Engine
type Engine struct {
	Logger   *log.Logger
	Cfg      *models.Config
	TG       *TGEngine
	SE       *SoundEngine
	Crawlers []crawlers.Crawler
}

// Init : Initialize Engine Driver
func (c *Engine) Init(logger *log.Logger, filePath string) error {
	var err error

	c.Logger = logger

	// Config
	c.Cfg, err = config.LoadConfig(filePath)
	if err != nil {
		return err
	}

	// TG Engine
	c.TG = &TGEngine{}
	c.TG.Cfg = c.Cfg
	err = c.TG.GenerateBot()
	if err != nil {
		return err
	}
	c.Logger.Println("Telegram Engine 세팅 완료!")

	// Sound Engine
	c.SE = &SoundEngine{}
	err = c.SE.Init(c.Cfg.Sound.On, c.Cfg.Sound.FilePath)
	if err != nil {
		return err
	}
	c.Logger.Println("Sound Engine 세팅 완료! 테스트용으로 사운드를 한 번 재생합니다.")
	c.SE.Play()

	return nil
}

// Run : Main Func
func (c *Engine) Run() {
	// var
	isFirst := true
	errorCount := 0
	var prevData []string

	// Set
	var targetCrawler crawlers.Crawler
	cids := make([]crawlers.Crawler, 4)
	cids[0] = crawlers.CID1{}
	cids[1] = crawlers.CID2{}
	cids[2] = crawlers.CID3{}
	cids[3] = crawlers.CID4{}

	// Main Loop
	c.Logger.Println("메인 루프 시작.")
	for {
		if errorCount >= 5 {
			c.Logger.Printf("[ERROR] 에러 다발, 서버 문제로 추정, 당분간 대기.")
			time.Sleep(time.Millisecond * time.Duration(1000*60*5))
			errorCount = 0
			continue
		}

		// Get New Items
		data, err := utils.GetRecentReports(*c.Cfg, "", "1", "30")
		if err != nil {
			c.Logger.Printf("[ERROR] GetRecentReports() : %s", err)
			errorCount++
			continue
		}

		// first time
		if isFirst {
			prevData = utils.MakePrevData(data)
			for idx := range prevData {
				c.TG.AddMessageWithRceptNo(prevData[idx])
			}
			isFirst = false
			c.Logger.Printf("첫 수집 완료.\n")
			time.Sleep(time.Millisecond * time.Duration(c.Cfg.Crawler.DelayTimer))
			continue
		}

		// Detect New Item
		for idx := range data {
			if !utils.IsContain(data[idx].RceptNo, prevData) { // New Item

				if strings.Contains(data[idx].ReportNM, "[기재정정]") {
					continue
				}

				if cids[0].IsTarget(data[idx].ReportNM) {
					targetCrawler = cids[0]
				} else if cids[1].IsTarget(data[idx].ReportNM) {
					targetCrawler = cids[1]
				} else if cids[2].IsTarget(data[idx].ReportNM) {
					targetCrawler = cids[2]
				} else if cids[3].IsTarget(data[idx].ReportNM) {
					targetCrawler = cids[3]
				} else {
					continue
				}

				report := models.Report{}
				report.Title = data[idx].ReportNM
				report.RceptNo = data[idx].RceptNo
				report.CorpName = data[idx].CorpName
				report.CrawlerID = targetCrawler.GetCID()

				err = targetCrawler.GetDetail(&report)
				if err != nil {
					c.Logger.Printf("[ERROR] crawler.GetDetail() : crawler(%s) : idx(%d) : %s", targetCrawler.GetCID(), idx, err)
					errorCount++
					continue
				}
				// Messaging Goroutine
				go func() {
					err = c.TG.SendMessage(report)
					if err == nil {
						go c.SE.Play()
					}
				}()
			}
		}

		prevData = utils.MakePrevData(data)
		time.Sleep(time.Millisecond * time.Duration(c.Cfg.Crawler.DelayTimer))
	} // main loop end
}
