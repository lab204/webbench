package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"net/http"
	"os"
	"time"
)

type Options struct {
	Url     string `short:"u" long:"url" description:"URL" required:"true"`
	Fore    bool   `short:"f" long:"fore" description:"Don't wait for reply from server."`
	Reload  bool   `short:"r" long:"reload" description:"Send reload request - Pragma: no-cache."`
	Time    int    `short:"t" long:"time" description:"Run benchmark for <sec> seconds. Default 30."`
	Client  int    `short:"c" long:"client" description:"Run <n> HTTP clients at once. Default one."`
	Http09  bool   `short:"9" long:"http09" description:"Use HTTP/0.9 style requests."`
	Http10  bool   `short:"1" long:"http10" description:"Use HTTP/1.0 protocol."`
	Http11  bool   `short:"2" long:"http11" description:"Use HTTP/1.1 protocol."`
	Proxy   string `short:"p" long:"proxy" description:"Use proxy server for request."`
	Get     bool   `long:"get" description:"Use GET request method."`
	Head    bool   `long:"head" description:"Use HEAD request method."`
	Options bool   `long:"options" description:"Use OPTIONS request method."`
	Trace   bool   `long:"trace" description:"Use TRACE request method."`
	Version bool   `short:"v" long:"version" description:"Display program version."`
}

const (
	DEFAULT_RUN_TIME     = 10 * time.Second
	DEFAULT_CLIENT_COUNT = 1
)

var (
	parser *flags.Parser
	op     Options
	c      = make(chan int)
	after  = time.After(30 * time.Second)
	count  = 0
	failed = 0
)

func init() {
	parser = flags.NewParser(&op, flags.Default)
	parser.ParseArgs(os.Args)
}

func Usage() {
	parser.WriteHelp(os.Stderr)
}

func Get(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		failed += 1
	}
	count += 1

	if op.Trace {
		fmt.Println("Status:" + resp.Status)
	}

}

func Keep() {
	for {
		select {
		case <-c:
		case <-after:
			fmt.Printf("Speed=%d pages/min\n", count*2)
			fmt.Printf("Request: %d successed, %d failed", count, failed)
			fmt.Println("done !")
			return
		}

	}
}

func Run() {
	for {
		select {
		case c <- 1:
			Get(op.Url)
		}
	}
}

func main() {
	if op.Url != "" {
		go Run()
		Keep()
	} else {
		fmt.Println(op.Client)
		Usage()
	}
}
