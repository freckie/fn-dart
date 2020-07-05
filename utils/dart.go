package utils

import (
	"encoding/json"
	"fmt"
	"fn-dart/models"
	"net/http"
)

const apiListURL = "https://opendart.fss.or.kr/api/list.json"

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
	fmt.Println(req.URL)

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
