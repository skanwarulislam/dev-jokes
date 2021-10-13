package models

import (
	"github.com/markbates/goth"
)

type Data struct {
	User        goth.User
	Joke        ChuckNorris
	CompanyName string
	Provider    string
}
type ChuckNorris struct {
	Value   string `json:"value"`
	Id      string `json:"id"`
	IconUrl string `json:"icon_url"`
	Url     string `json:"url"`
}
