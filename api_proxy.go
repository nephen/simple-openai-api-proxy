package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"flag"
	"strconv"
	"net/http"
	"net/http/httputil"
)

const (
	DAEMON = "daemon"
)

var (
	port int
	damaen bool
)

func init()  {
	flag.IntVar(&port, "port", 8080, "监听端口")
	flag.BoolVar(&damaen, DAEMON, false, "是否后台运行")
}

func ReverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[*] receive a request from %s, request header: %s: \n", r.RemoteAddr, r.Header)
	target := "api.openai.com"
	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = target
		req.Host = target
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
	fmt.Printf("[*] receive the destination website response header: %s\n", w.Header())
}

func StripSlice(slice []string, element string) []string {
	for i := 0; i < len(slice); {
		if slice[i] == element && i != len(slice)-1 {
			slice = append(slice[:i], slice[i+1:]...)
		} else if slice[i] == element && i == len(slice)-1 {
			slice = slice[:i]
		} else {
			i++
		}
	}
	return slice
}

func SubProcess(args []string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] Error: %s\n", err)
	}
	return cmd
}

func main() {
	flag.Parse()
	fmt.Printf("[*] PID: %d PPID: %d ARG: %s\n", os.Getpid(), os.Getppid(), os.Args)
	if damaen {
		SubProcess(StripSlice(os.Args, "-"+DAEMON))
		fmt.Printf("[*] Daemon running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
		os.Exit(0)
	}
	fmt.Printf("[*] Forever running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
	fmt.Printf("[*] Starting server at port %v\n", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), http.HandlerFunc(ReverseProxyHandler)); err != nil {
		log.Fatal(err)
	}
}