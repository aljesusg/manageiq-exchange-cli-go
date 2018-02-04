package menu

import (
	"errors"
	"flag"
	"fmt"
	"manageiq-exchange/api"
	"manageiq-exchange/constants"
	"os"
)

func Menu() {

	Banner()
	var host string
	var port int
	var version bool
	var providers bool
	var users bool
	var expand bool
	flag.StringVar(&host, "host", "localhost", "specify host to use.  defaults to localhost.")
	flag.IntVar(&port, "port", 3000, "specify port to use.  defaults to 3000")
	flag.BoolVar(&version, "version", false, "About version")
	flag.BoolVar(&providers, "providers", false, "About providers")
	flag.BoolVar(&users, "users", false, "About users")
	flag.BoolVar(&expand, "expand", false, "Expand information")
	flag.Parse()

	server, err := GetServer()

	if err != nil {
		fmt.Print(err)
	}
	var miq_exchange api.Api
	miq_exchange.Init(server, port)

	if version {
		info := miq_exchange.GetInfo()
		fmt.Printf(info.Print())
	}

	if users {
		users := miq_exchange.GetUsers(expand)
		fmt.Printf(users.Print(miq_exchange.Data.Meta.TotalCount))
	}

}

func Banner() {
	fmt.Println("\033[0;31m", constants.BANNER, "\033[0m")
}

func GetOsEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func SetOsEnv(env string, key string) {
	os.Setenv(env, key)
}

func GetServer() (string, error) {
	server := GetOsEnv("EXCHANGE_SERVER", "")
	if len(server) == 0 {
		return "", errors.New("You need to set the environment EXCHANGE_SERVER (ex: localhost)")
	}
	return server, nil
}
