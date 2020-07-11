package main

import (
	"log"
	"os"

	"github.com/Lemo-yxk/lemo"
	"github.com/Lemo-yxk/lemo/http"
	server3 "github.com/Lemo-yxk/lemo/http/server"
	"github.com/Lemo-yxk/lemo/tcp/server"
	server2 "github.com/Lemo-yxk/lemo/websocket/server"
)

func main() {

	// utils.Process.Fork(run, 1)
	//
	// go func() {
	// 	http.HandleFunc("/reload", func(writer http.ResponseWriter, request *http.Request) {
	// 		utils.Process.Reload()
	// 	})
	// 	console.Log(http.ListenAndServe(":12345", nil))
	// }()

	// run()
	// utils.Signal.ListenKill().Done(func(sig os.Signal) {
	// 	console.Info(sig)
	// })

	// var progress = utils.HttpClient.NewProgress()
	// progress.Rate(0.01).OnProgress(func(p []byte, current int64, total int64) {
	// 	console.OneLine("Downloading... %d %d B complete", current, total)
	// })
	//
	// utils.HttpClient.Get("https://www.twle.cn/static/js/jquery.min.js").Progress(progress).Send()

	// console.SetFormatter(console.NewJsonFormatter())

}

func run() {

	var webSocketServer = &server2.Server{Host: "127.0.0.1:8667", Path: "/"}

	var webSocketServerRouter = &server2.Router{IgnoreCase: true}

	webSocketServer.Use(func(next server2.Middle) server2.Middle {
		return func(conn *server2.WebSocket, receive *lemo.ReceivePackage) {
			next(conn, receive)
		}
	})

	webSocketServer.OnMessage = func(conn *server2.WebSocket, messageType int, msg []byte) {
		log.Println(len(msg))
	}

	webSocketServerRouter.Group("/hello").Handler(func(handler *server2.RouteHandler) {
		handler.Route("/world").Handler(func(conn *server2.WebSocket, receive *lemo.Receive) error {
			log.Println(string(receive.Body.Message))
			return conn.Json(lemo.JsonPackage{
				Event: "/hello/world",
				Data:  "i am server",
			})
		})
	})

	go webSocketServer.SetRouter(webSocketServerRouter).Start()

	var httpServer = server3.Server{Host: "127.0.0.1:8666"}

	var httpServerRouter = &server3.Router{}

	httpServer.Use(func(next server3.Middle) server3.Middle {
		return func(stream *http.Stream) {
			// if stream.Request.Header.Get("Upgrade") == "websocket" {
			// 	httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "0.0.0.0:8667"}).ServeHTTP(stream.Response, stream.Request)
			// } else {
			// 	log.Println(1, "start")
			// 	next(stream)
			// 	log.Println(1, "end")
			// }
			next(stream)
		}
	})

	httpServer.Use(func(next server3.Middle) server3.Middle {
		return func(stream *http.Stream) {
			// log.Println(2, "start")
			// next(stream)
			// log.Println(2, "end")
			next(stream)
		}
	})

	httpServerRouter.Route("GET", "/hello").Handler(func(stream *http.Stream) error {
		// log.Println("handler")
		return stream.EndString("hello")
	})

	httpServerRouter.Group("/hello").Handler(func(handler *server3.RouteHandler) {
		handler.Get("/world").Handler(func(t *http.Stream) error {
			return t.JsonFormat("SUCCESS", 200, os.Getpid())
		})
	})

	httpServer.OnSuccess = func() {
		log.Println(httpServer.LocalAddr())
	}

	go httpServer.SetRouter(httpServerRouter).Start()

	log.Println("start success")

	var tcpServer = &server.Server{Host: "127.0.0.1:8888"}

	tcpServer.OnMessage = func(conn *server.Socket, messageType int, msg []byte) {
		log.Println(len(msg))
	}

	var tcpServerRouter = &server.Router{IgnoreCase: true}

	tcpServerRouter.Group("/hello").Handler(func(handler *server.RouteHandler) {
		handler.Route("/world").Handler(func(conn *server.Socket, receive *lemo.Receive) error {
			log.Println(string(receive.Body.Message))
			return nil
		})
	})

	tcpServer.OnSuccess = func() {
		log.Println(tcpServer.LocalAddr())
	}

	tcpServer.SetRouter(tcpServerRouter).Start()

}
