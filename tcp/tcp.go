package tcp

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
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
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			req, err := http.ReadRequest(bufio.NewReader((conn)))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
			res := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World")),
			}
			res.Write(conn)
			conn.Close()
		}()
	}
}

func Client() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
