package tcp

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func Server() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Serer is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				req, err := http.ReadRequest(bufio.NewReader((conn)))
				if err != nil {
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				dump, err := httputil.DumpRequest(req, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello World"
				res := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				res.Write(conn)
			}
		}()
	}
}

func Client() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	for {
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		req, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		err = req.Write(conn)
		if err != nil {
			panic(err)
		}
		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
