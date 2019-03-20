package models

type UrbanDictionaryResponse struct{
	Id string
	Term string
	URL string `json:"url"`
	Definition string
}
