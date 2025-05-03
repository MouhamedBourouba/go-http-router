package main

import (
	"log"
	"net"
	"net/http"

	"github.com/mouhamedBourouba/go-http-router"
)

func main() {
  println("the best router")
  lis, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatal("could not listen", err)
  }
  log.Print("Server open listing on port 8080")

  router := gohttprouter.New()

  router.Get("/", func(w http.ResponseWriter, r *gohttprouter.Request) {
    w.Write([]byte("gg brother"))
  })

  router.Get("/{thebest}", func(w http.ResponseWriter, r *gohttprouter.Request) {
    w.Write([]byte("gg brother"))
  })

  log.Fatal(http.Serve(lis, router))
}
