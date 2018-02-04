package user

import (
	"fmt"
	utils "manageiq-exchange/models/utils"
)

type User struct {
	GithubId        int    `json:"github_id"`
	Login           string `json:"login"`
	Url_profile     string `json:"url_profile"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Company         string `json:"company"`
	GithubType      string `json:"github_type"`
	GithubBlog      string `json:"github_blog"`
	GithubLocation  string `json:"github_location"`
	GithubBio       string `json:"github_bio"`
	GithubCreatedAt string `json:"github_created_at"`
	GithubUpdatedAt string `json:"github_updated_at"`
}

type UserCollection struct {
	Users []User
	Total int
}

func (c *UserCollection) Init(data []interface{}) {
	for _, v := range data {
		var u User
		u.Init(v.(map[string]interface{}))
		c.Users = append(c.Users, u)
	}
	c.Total = len(data)
}

func (u *User) Init(data map[string]interface{}) {
	utils.CreateFromMap(data, &u)
}

func (u *User) Print() string {
	var result string
	result = fmt.Sprintf("%s: %s (%d)\n\n", utils.PrintColor("User", "Red"), u.Login, u.GithubId)
	result += utils.PrintValues(u, "    ", []string{"Login", "GithubId"})
	return result
}

func (c *UserCollection) Print(total int) string {
	var result string
	result = fmt.Sprintf("\n\n%s (%d/%d)\n\n------------------------------------------\n\n", utils.PrintColor("Users", "Red"), c.Total, total)
	for _, user := range c.Users {
		result += fmt.Sprintf("%s\n\n", user.Print())
	}
	return result
}
