package main

import (
	"flag"
	// "fmt"
	// "log"
	"os"
	"os/signal"
	"syscall"
	// "time"

	"code.sohuno.com/kzapp/push-server/api"
)

var (
	flagSet = flag.NewFlagSet("pushapi", flag.ExitOnError)

	config      = flagSet.String("config", "", "path to config file")
	showVersion = flagSet.Bool("version", false, "print version string")

	httpAddress = flagSet.String("http-address", "0.0.0.0:4171", "<addr>:<port> to listen on for HTTP clients")
)

func init() {
	// flagSet.Var(&nsqlookupdHTTPAddresses, "lookupd-http-address", "lookupd HTTP address (may be given multiple times)")
	// flagSet.Var(&nsqdHTTPAddresses, "nsqd-http-address", "nsqd HTTP address (may be given multiple times)")
}

func main() {
	flagSet.Parse(os.Args[1:])

	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	pushAPI := api.NewPushAPI(httpAddress)
	pushAPI.Main()
	<-exitChan
	pushAPI.Exit()
}
