package controllers

type PasswordInput struct {
	NewPassword string
}

type UpdateCompanyInput struct {
	CompanyId    uint
	MailVerified bool
}

type CompanyMailingDetails struct {
	CompanyId      uint
	TemplateNumber int
}
