package controllers

type AdminAddUser struct {
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
}

type AdminAddCompany struct {
	CompanyName    string `json:"companyName"`
	HrEmail        string `json:"hrEmail"`
	CompanyAbout   string `json:"companyAbout"`
	CompanyCareers string `json:"companyCareers"`
	HandlerID      uint   `json:"handlerId"`
}
