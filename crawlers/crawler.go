package crawlers

import (
	"fn-dart/models"
)

type Crawler interface {
	GetCID() string                              // Crawler ID
	GetList(number int) ([]models.Report, error) // Crawl list of Reports
	GetDetail(item *models.Report) error         // Crawl detail values of each Reports
}

func GetCID(c Crawler) string {
	return c.GetCID()
}

func GetList(c Crawler, number int) ([]models.Report, error) {
	return c.GetList(number)
}

func GetDetail(c Crawler, item *models.Report) error {
	return c.GetDetail(item)
}
