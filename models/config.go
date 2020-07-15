package models

type Config struct {
	Telegram ConfigItemTelegram `json:"telegram"`
	Crawler  ConfigItemCrawler  `json:"crawler"`
	Sound    ConfigItemSound    `json:"sound"`
	Targets  ConfigItemTarget   `json:"targets"`
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

type ConfigItemSound struct {
	On       bool   `json:"on"`
	FilePath string `json:"file_path"`
}

type ConfigItemTarget struct {
	CID1 ConfigItemTargetItem `json:"1"`
	CID2 ConfigItemTargetItem `json:"2"`
	CID3 ConfigItemTargetItem `json:"3"`
	CID4 ConfigItemTargetItem `json:"4"`
}

type ConfigItemTargetItem struct {
	Description   string `json:"description"`
	ValuesCount   int    `json:"values_count"`
	MessageFormat string `json:"message_format"`
}
