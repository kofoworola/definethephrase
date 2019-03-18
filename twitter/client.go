package twitter

type Tweet struct {
	Id int64
	IdStr string `json:"id_str"`
	User User
	Text string
}
type User struct {
	Id int64
	IdStr string `json:"id_str"`
	Name string
	DisplayName string `json:"display_name"`
}
