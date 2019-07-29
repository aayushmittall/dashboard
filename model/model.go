package model

import "time"

//UserProfile struct
type UserProfile struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	Age       string `json:"age"`
	Email     string `json:"email"`
}

//UserAuth struct
type UserAuth struct {
	UID           int       `json:"uid"`
	Token         string    `json:"token"`
	TimeGenerated time.Time `json:"timegenerated"`
}
