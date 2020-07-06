package crawlers

import (
	"fn-dart/models"
	"fn-dart/utils"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var CID3Pattern = regexp.MustCompile(`기타 ?경영 ?사항 ?\(자율 ?공시\)`)
var CID3SubPattern1 = regexp.MustCompile(`1\. ?제목`)

type CID3 struct{}

func (c CID3) GetCID() string {
	return "3"
}

func (c CID3) IsTarget(title string) bool {
	matched := CID3Pattern.MatchString(title)
	return matched
}

func (c CID3) GetDetail(item *models.Report) error {
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
		if CID3SubPattern1.MatchString(cellTitle) {
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[0] = utils.TrimAll(value)
			return false
		}
		return true
	})
	return nil
}
