package model

//User struct
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	Age       string `json:"age"`
	Email     string `json:"email"`
}

//ResponseResult struct
type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
