package models

type User struct {
	github_id         int    `json:"github_id"`
	login             string `json:"login"`
	url_profile       string `json:"url_profile"`
	name              string `json:"name"`
	avatar            string `json:"avatar"`
	company           string `json:"company"`
	github_type       string `json:"github_type"`
	github_blog       string `json:"github_blog"`
	github_location   string `json:"github_location"`
	github_bio        string `json:"github_bio"`
	github_created_at string `json:"github_created_at"`
	github_updated_at string `json:"github_updated_at"`
}
