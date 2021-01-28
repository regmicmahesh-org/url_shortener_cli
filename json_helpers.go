package main

var IdLinkStruct = struct {
	CreatedAt      string        `json:"created_at"`
	ID             string        `json:"id"`
	Link           string        `json:"link"`
	CustomBitlinks []interface{} `json:"custom_bitlinks"`
	LongURL        string        `json:"long_url"`
	Archived       bool          `json:"archived"`
	Tags           []interface{} `json:"tags"`
	Deeplinks      []interface{} `json:"deeplinks"`
	References     struct {
		Group string `json:"group"`
	} `json:"references"`
}{}

type Format struct {
	GroupGUID string `json:"group_guid"`
	Domain    string `json:"domain"`
	LongURL   string `json:"long_url"`
}
