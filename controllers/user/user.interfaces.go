package controllers

type PasswordInput struct {
	NewPassword string
}

type UpdateCompanyInput struct {
	CompanyId    uint `json:"companyId,string,omitempty"`
	MailVerified bool
}

type CompanyMailingDetails struct {
	CompanyId      uint `json:"companyId,string,omitempty"`
	TemplateNumber int  `json:"templateNumber,string,omitempty"`
}

type CompanyTemplateUpdateDetails struct {
	CompanyId       uint `json:"companyId,string,omitempty"`
	TemplateNumber  int  `json:"templateNumber,string,omitempty"`
	TemplateContent string
}
