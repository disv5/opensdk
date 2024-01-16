package account

type GetChildrenAccountByAccountIdResponse struct {
	ChildrenAccounts []ChildrenAccount `json:"childrenAccounts"`
}

type ChildrenAccount struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CompanyName string `json:"companyName"`
	SiteName    string `json:"siteName"`
	SiteURL     string `json:"siteUrl"`
}