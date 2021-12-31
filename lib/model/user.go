package model

// User structure
type User struct {
	User_Name    string `json:"userName"`
	First_Name   string `json:"firstName"`
	Last_Name    string `json:"lastName"`
	Phone_Number string `json:"phoneNumber,omitempty"`
	DateCreated  string `json:"dateCreated,omitempty"`
	DateModified string `json:"dateModified,omitempty"`
}
