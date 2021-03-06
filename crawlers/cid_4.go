package crawlers

import (
	"fn-dart/models"
	"fn-dart/utils"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var CID4Pattern = regexp.MustCompile(`투자 ?판단 ?관련 ?주요 ?경영 ?사항`)
var CID4SubPattern1 = regexp.MustCompile(`1\. ?제목`)

type CID4 struct{}

func (c CID4) GetCID() string {
	return "4"
}

func (c CID4) IsTarget(title string) bool {
	matched := CID4Pattern.MatchString(title)
	return matched
}

func (c CID4) GetDetail(item *models.Report) error {
	// Get Table
	table, err := utils.GetMainTable(item)
	if err != nil {
		return err
	}

	item.Values = make([]string, 1)

	// Parse table
	trs := table.Find("tr")
	trs.EachWithBreak(func(idx int, sel *goquery.Selection) bool {
		tds := sel.Find("td")
		if tds.Length() <= 1 {
			return true
		}

		cellTitle, _ := utils.ReadCP949(tds.Text())
		if CID4SubPattern1.MatchString(cellTitle) {
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[0] = utils.TrimAll(value)
			return false
		}
		return true
	})
	return nil
}
