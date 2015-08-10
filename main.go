package main

import (
	"fmt"
	"os"
)

func usage() {
	tmp := "webbench [option]... URL\n"
	tmp += "  -f|--force               Don't wait for reply from server.\n"
	tmp += "  -r|--reload              Send reload request - Pragma: no-cache.\n"
	tmp += "  -t|--time <sec>          Run benchmark for <sec> seconds. Default 30.\n"
	tmp += "  -p|--proxy <server:port> Use proxy server for request.\n"
	tmp += "  -c|--clients <n>         Run <n> HTTP clients at once. Default one.\n"
	tmp += "  -9|--http09              Use HTTP/0.9 style requests.\n"
	tmp += "  -1|--http10              Use HTTP/1.0 protocol.\n"
	tmp += "  -2|--http11              Use HTTP/1.1 protocol.\n"
	tmp += "  --get                    Use GET request method.\n"
	tmp += "  --head                   Use HEAD request method.\n"
	tmp += "  --options                Use OPTIONS request method.\n"
	tmp += "  --trace                  Use TRACE request method.\n"
	tmp += "  -?|-h|--help             This information.\n"
	tmp += "  -V|--version             Display program version.\n"
	fmt.Fprintf(os.Stderr, tmp)
}

func main() {
	usage()
}
