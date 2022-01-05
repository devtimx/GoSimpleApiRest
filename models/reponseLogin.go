package models

/*ResponseLogin stores the token obtained with the login*/
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
