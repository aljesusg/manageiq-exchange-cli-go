package user

import (
	"fmt"
	"manageiq-exchange/models/utils"
	"reflect"
	"testing"
)

func TestUserInit(t *testing.T) {

	want := User{
		Login:    "aljesusg",
		Name:     "Alberto",
		GithubId: 1,
	}
	var data = map[string]interface{}{
		"login":     "aljesusg",
		"github_id": 1,
		"name":      "Alberto",
	}
	var user User
	user.Init(data)
	if !reflect.DeepEqual(user, want) {
		t.Errorf("UserInit returned %+v, want %+v", user, want)
	}
}

func TestUserPrint(t *testing.T) {
	want := fmt.Sprintf("%s: aljesusg (1)\n\n", utils.PrintColor("User", "Red"))
	want += fmt.Sprintf("    Name : Alberto\n")

	user := User{
		Login:    "aljesusg",
		Name:     "Alberto",
		GithubId: 1,
	}
	if !reflect.DeepEqual(user.Print(), want) {
		t.Errorf("User Print returned -%+v-, want -%+v-", user.Print(), want)
	}
}
