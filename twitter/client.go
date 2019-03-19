package twitter

type WebhookLoad struct {
	UserId string `json:"for_user_id"`
	TweetCreateEvent []Tweet `json:"tweet_create_events"`
}
type Tweet struct {
	Id    int64
	IdStr string `json:"id_str"`
	User  User
	Text  string
}
type User struct {
	Id          int64
	IdStr       string `json:"id_str"`
	Name        string
	DisplayName string `json:"display_name"`
}
