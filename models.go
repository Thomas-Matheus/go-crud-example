package main

// Main model of the service
type Article struct {
	Uuid          *string `json:"uuid"`
	Title         *string `json:"title"`
	Text          *string `json:"text"`
	Publisher     *string `json:"publisher"`
	Editor        *string `json:"editor"`
	DatePublished *string `json:"datePublished"`
	DateUpdated   *string `json:"dateUpdated"`
	IsDeleted     *bool   `json:"isDeleted"`
}

// Used for bulk requests
type MultipleResponse struct {
	Responses *[]DataResponse `json:"responses"`
}

// Basic model-result of an action over some article
type DataResponse struct {
	Id     *string   `json:"id"`
	Errors *[]string `json:"errors"`
	Trace  *string   `json:"trace"`
}
