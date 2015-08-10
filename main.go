package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"net/http"
	"os"
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

var parser *flags.Parser
var op Options

func init() {
	parser = flags.NewParser(&op, flags.Default)
	parser.ParseArgs(os.Args)
}

func Usage() {
	parser.WriteHelp(os.Stderr)
}

func Start(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	if op.Http09 == true {
		req.Header.Set("", "")
	} else if op.Http10 == true {
		req.Header.Set("", "")
	} else if op.Http11 == true {
		req.Header.Set("", "")
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if op.Trace {
		fmt.Println("Status:" + resp.Status)
	}

}

func main() {

	if op.Url != "" {
		Start(op.Url)
	} else {
		Usage()
	}
}
