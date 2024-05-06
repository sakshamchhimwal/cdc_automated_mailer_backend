package controllers

type LoginInput struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
