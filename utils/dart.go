package utils

import (
	"encoding/json"
	"fmt"
	"fn-dart/models"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const apiListURL = "https://opendart.fss.or.kr/api/list.json"
const dartPopupURL = "http://dart.fss.or.kr/dsaf001/main.do?rcpNo=%s"
const dartReportURL = "http://dart.fss.or.kr/report/viewer.do?rcpNo=%s&dcmNo=%s&eleId=0&offset=0&length=0&dtd=HTML"

var numberPattern = regexp.MustCompile(`[0-9]+`)

func GetRecentReports(cfg models.Config, startDate string, page string, maxItems string) ([]models.APIResultListItem, error) {
	req, err := http.NewRequest("GET", apiListURL, nil)
	if err != nil {
		return nil, err
	}

	// Make query
	query := req.URL.Query()
	query.Add("crtfc_key", cfg.Crawler.DartAPIKey)
	query.Add("bgn_de", startDate)
	query.Add("end_de", "")
	query.Add("page_no", page)
	query.Add("page_count", maxItems)
	req.URL.RawQuery = query.Encode()

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] Request Status Code: %d, %s", resp.StatusCode, resp.Status)
	}

	// Load JSON
	jsonData := models.APIResult{}
	err = json.NewDecoder(resp.Body).Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData.List, nil
}

func GetMainTable(item *models.Report) (*goquery.Selection, error) {
	// Popup request
	popupURL := fmt.Sprintf(dartPopupURL, item.RceptNo)
	req, err := http.Get(popupURL)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] Request Status Code : %d, %s", req.StatusCode, req.Status)
	}

	// Load HTML
	html, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		return nil, err
	}

	aTag := html.Find("div.view_search").Find("a")
	onclick, isFound := aTag.Attr("onclick")
	if !isFound {
		return nil, err
	}
	dcmNo := numberPattern.FindAllString(onclick, -1)[1]
	item.DcmNo = dcmNo

	// Report request
	reportURL := fmt.Sprintf(dartReportURL, item.RceptNo, dcmNo)
	item.ReportURL = reportURL
	req2, err := http.Get(reportURL)
	if err != nil {
		return nil, err
	}
	defer req2.Body.Close()

	if req2.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] Request Status Code : %d, %s", req2.StatusCode, req2.Status)
	}

	// Load HTML
	html2, err := goquery.NewDocumentFromReader(req2.Body)
	if err != nil {
		return nil, err
	}

	// Parsing
	table := html2.Find("table > tbody")

	return table, nil
}
