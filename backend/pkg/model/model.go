package model

type Content struct {
	Url        string `json:"url"`
	Uid        string `json:"uid"`
	Content_id string `json:"content_id"`
	Date       int    `json:"date"`
	Comment    string `json:"comment"`
}
