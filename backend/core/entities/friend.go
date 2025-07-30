package entities

type Friends struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ProfilePic string `json:"profile_pic"`
	Validated  string `json:"validated"`
}
