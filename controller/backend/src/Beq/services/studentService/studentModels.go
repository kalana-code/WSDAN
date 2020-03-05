package studentservice

// Struct of the json object for student. when creating a student the datas should be pass in this struct.
type studentModal struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Gender    string `json:"Gender"`
	BirthDay  string `json:"BirthDay"`
}

type studentLoginModal struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}
