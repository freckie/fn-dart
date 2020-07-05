package models

type Config struct {
	Telegram ConfigItemTelegram `json:"telegram"`
	Crawler  ConfigItemCrawler  `json:"crawler"`
	Targets  []ConfigItemTarget `json:"targets"`
}

type ConfigItemTelegram struct {
	BotToken string  `json:"bot_token"`
	Channels []int64 `json:"channels"`
}

type ConfigItemCrawler struct {
	DartAPIKey string `json:"dart_api_key"`
	DelayTimer int    `json:"delay_timer"`
	MaxProcs   int    `json:"max_procs"`
}

type ConfigItemTarget struct {
	CrawlerID     int    `json:"crawler_id"`
	On            bool   `json:"on"`
	Keyword       string `json:"keyword"`
	MessageFormat string `json:"message_format"`
}
