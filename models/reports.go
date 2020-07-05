package models

type Report struct {
	CrawlerID int
	Title     string
	CorpName  string
	Values    []string
}
