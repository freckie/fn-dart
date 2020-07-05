package crawlers

import (
	"fn-dart/models"
)

type CID4 struct {
	Cfg models.Config
}

func (c CID4) GetCID() string {
	return "4"
}

func (c CID4) GetDetail(item *models.Report) error {
	return nil
}
