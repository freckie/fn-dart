package crawlers

import (
	"fmt"
	"fn-dart/models"
	"net/http"
	"regexp"
)

const CID4Pattern = regexp.MustCompile(`투자 ?판단 ?관련 주요 ?경영 ?사항`)

type CID4 struct{}

func (c CID4) GetCID() string {
	return "4"
}

func (c CID4) IsTarget(item *models.Report) bool {
	matched, _ := CID4Pattern.MatchString(item.Title)
	return matched
}

func (c CID4) GetDetail(item *models.Report) error {
	// Popup request
	popupURL := fmt.Sprintf(DartPopupURL, item.RceptNo)
	req, err := http.Get(popupURL)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return result, fmt.Errorf("[ERROR] Request Status Code : %d, %s", req.StatusCode, req.Status)
	}

	return nil
}
