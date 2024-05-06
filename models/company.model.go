package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName    string `gorm:"not null" json:"companyName"`
	HrEmail        string `gorm:"not null" json:"hrEmail"`
	CompanyAbout   string `json:"companyAbout"`
	CompanyCareers string `json:"companyCareers"`
	Template1      string `json:"template1"`
	Template2      string `json:"template2"`
	Template3      string `json:"template3"`
	HandlerID      uint   `json:"handlerID"`                           // this will store the ID of the User
	Handler        User   `gorm:"foreignKey:HandlerID" json:"handler"` // corrected foreign key definition
}