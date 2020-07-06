package crawlers

import (
	"fmt"
	"fn-dart/models"
	"fn-dart/utils"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var CID1Pattern = regexp.MustCompile(`단일 ?판매 ?ㆍ ?공급 ?계약 ?체결`)
var CID1SubPattern1 = regexp.MustCompile(`1\. ?판매ㆍ공급 ?계약 ?(내용|구분)`)
var CID1SubPattern11 = regexp.MustCompile(`- ?체결 ?계약명?`)
var CID1SubPattern2 = regexp.MustCompile(`2\. ?계약 ?내역`)
var CID1SubPattern21 = regexp.MustCompile(`확정 ?계약 ?금액`)
var CID1SubPattern22 = regexp.MustCompile(`매출액 ?대비 ?\(\%\)`)
var CID1SubPattern3 = regexp.MustCompile(`3\. ?계약 ?상대방?`)
var CID1SubPattern4 = regexp.MustCompile(`4\. ?판매ㆍ공급 ?지역`)

type CID1 struct{}

func (c CID1) GetCID() string {
	return "1"
}

func (c CID1) IsTarget(title string) bool {
	matched := CID1Pattern.MatchString(title)
	return matched
}

func (c CID1) GetDetail(item *models.Report) error {
	// Get Table
	table, err := utils.GetMainTable(item)
	if err != nil {
		return err
	}

	item.Values = make([]string, 5)

	// Parse table
	trs := table.Find("tr")
	level := 0
	trs.EachWithBreak(func(idx int, sel *goquery.Selection) bool {
		tds := sel.Find("td")
		if tds.Length() <= 1 {
			return true
		}

		cellTitle, _ := utils.ReadCP949(tds.Text())

		if level == 2 {
			if CID1SubPattern21.MatchString(cellTitle) {
				value, _ := utils.ReadCP949(tds.Next().Text())
				item.Values[1] = utils.TrimAll(value)
			} else if CID1SubPattern22.MatchString(cellTitle) {
				value, _ := utils.ReadCP949(tds.Next().Text())
				item.Values[2] = utils.TrimAll(value)
			}
		}

		if CID1SubPattern1.MatchString(cellTitle) {
			level = 1
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[0] = utils.TrimAll(value)
		} else if CID1SubPattern11.MatchString(cellTitle) {
			level = 1
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[1] = utils.TrimAll(value)
		} else if CID1SubPattern3.MatchString(cellTitle) {
			level = 1
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[3] = utils.TrimAll(value)
		} else if CID1SubPattern4.MatchString(cellTitle) {
			level = 1
			value, _ := utils.ReadCP949(tds.Next().Text())
			item.Values[4] = utils.TrimAll(value)
		} else if CID1SubPattern2.MatchString(cellTitle) || level == 2 {
			level = 2
			cellSubTitle, _ := utils.ReadCP949(tds.Next().Text())
			fmt.Println(cellSubTitle)
			if CID1SubPattern21.MatchString(cellSubTitle) {
				value, _ := utils.ReadCP949(tds.Next().Text())
				item.Values[1] = utils.TrimAll(value)
			} else if CID1SubPattern22.MatchString(cellSubTitle) {
				value, _ := utils.ReadCP949(tds.Next().Text())
				item.Values[2] = utils.TrimAll(value)
			}
		} else {
			return false
		}
		return false
	})
	return nil
}
