package models

type APIResult struct {
	Status     string              `json:"status"`
	Message    string              `json:"message"`
	PageNo     int                 `json:"page_no"`
	PageCount  int                 `json:"page_count"`
	TotelCount int                 `json:"total_count"`
	TotalPage  int                 `json:"total_page"`
	List       []APIResultListItem `json:"list"`
}

type APIResultListItem struct {
	CorpCode  string `json:"corp_code"`
	CorpName  string `json:"corp_name"`
	StockCode string `json:"stock_code"`
	CorpCls   string `json:"corp_cls"`
	ReportNM  string `json:"report_nm"`
	RceptNo   string `json:"rcept_no"`
	FirNM     string `json:"fir_nm"`
	RceptDt   string `json:"rcept_dt"`
	RM        string `json:"rm"`
}
