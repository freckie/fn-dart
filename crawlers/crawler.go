package crawlers

import (
	"fn-dart/models"
)

const DartPopupURL = "http://dart.fss.or.kr/dsaf001/main.do?rcpNo=%s"
const DartReportURL = "http://dart.fss.or.kr/report/viewer.do?rcpNo=%s&dcmNo=%s&eleId=0&offset=0&length=0&dtd=HTML"

type Crawler interface {
	GetCID() string                      // Crawler ID
	GetDetail(item *models.Report) error // Crawl detail values of each Reports
}

func GetCID(c Crawler) string {
	return c.GetCID()
}

func GetDetail(c Crawler, item *models.Report) error {
	return c.GetDetail(item)
}
