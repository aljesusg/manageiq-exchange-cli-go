package main


import (
  "fmt"
  "flag"
)

func menu(){

  banner()
  var host string
  var port int
  flag.StringVar(&host, "host", "localhost", "specify host to use.  defaults to localhost.")
  flag.IntVar(&port, "port", 3000, "specify port to use.  defaults to 3000")
  flag.Parse()

  server, err := GetServer()

  if err != nil {
    fmt.Print(err)
  }
  var miq_exchange Api;
  miq_exchange.server= server
  miq_exchange.port = port
  fmt.Printf("Version : %s",miq_exchange.GetVersion())

}

func banner(){
  fmt.Println("\033[0;31m",BANNER,"\033[0m")
}
