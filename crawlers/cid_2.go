package crawlers

import (
	"fn-dart/models"
	"fn-dart/utils"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var CID2Pattern = regexp.MustCompile(`기타 ?경영 ?사항 ?\(특허권 ?취득\)`)
var CID2SubPattern1 = regexp.MustCompile(`1\. ?(특허)? ?명칭`)
var CID2SubPattern2 = regexp.MustCompile(`5\. ?특허 ?활용 ?계획`)

type CID2 struct{}

func (c CID2) GetCID() string {
	return "2"
}

func (c CID2) IsTarget(title string) bool {
	matched := CID2Pattern.MatchString(title)
	return matched
}

func (c CID2) GetDetail(item *models.Report) error {
	// Get Table
	table, err := utils.GetMainTable(item)
	if err != nil {
		return err
	}

	item.Values = make([]string, 2)

	// Parse table
	trs := table.Find("tr")
	trs.EachWithBreak(func(idx int, sel *goquery.Selection) bool {
		tds := sel.Find("td")
		if tds.Length() <= 1 {
			return true
		}

		cellTitle, _ := utils.ReadCP949(tds.Text())
		if CID2SubPattern1.MatchString(cellTitle) {
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[0] = utils.TrimAll(value)
			return true
		} else if CID2SubPattern2.MatchString(cellTitle) {
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[1] = utils.TrimAll(value)
			return false
		}
		return true
	})
	return nil
}
