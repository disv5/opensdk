package account

type GetAccountResponse struct {
	AccountInfoType AccountInfoType `json:"accountInfoType"`
}

type AccountInfoType struct {
	AccountID    int64   `json:"accountId"`
	UserID       int64   `json:"userId"`
	UserName     string  `json:"userName"`
	CompanyName  string  `json:"companyName"`
	SiteName     string  `json:"siteName"`
	SiteURL      string  `json:"siteUrl"`
	Industry1    string  `json:"industry1"`
	Industry2    string  `json:"industry2"`
	AuditState   int     `json:"auditState"`
	Balance      string  `json:"balance"`
	Budget       string  `json:"budget"`
	FinanceState int64   `json:"financeState"`
	DeliveryMode []int64 `json:"deliveryMode"`
	MinPrice     int64   `json:"minPrice"`
	PayAmount    string  `json:"payAmount"`
}
