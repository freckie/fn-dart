package models

type Report struct {
	CrawlerID string
	Title     string
	RceptNo   string
	DcmNo     string
	CorpName  string
	ReportURL string
	Values    []string
}
