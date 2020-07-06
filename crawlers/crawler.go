package crawlers

import (
	"fn-dart/models"
)

type Crawler interface {
	GetCID() string // Crawler ID
	IsTarget(title string) bool
	GetDetail(item *models.Report) error // Crawl detail values of each Reports
}

func GetCID(c Crawler) string {
	return c.GetCID()
}

func IsTarget(c Crawler, title string) bool {
	return c.IsTarget(title)
}

func GetDetail(c Crawler, item *models.Report) error {
	return c.GetDetail(item)
}
