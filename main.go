package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	address := flag.String("a", "127.0.0.1", "address")
	port := flag.Int("p", 8080, "port")
	flag.Parse()

	httpAddress := fmt.Sprintf("%s:%v", *address, *port)

	srv := &http.Server{Addr: httpAddress}

	dirPath, _ := os.Getwd()
	fmt.Println("Work dir:", dirPath)
	http.Handle("/", http.FileServer(http.Dir(dirPath)))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	fmt.Println("created httpServer success!")
	fmt.Println("server listening at: http://" + httpAddress)

	fmt.Println("if you want to stop the server, please press any key !")

	var command string
	fmt.Scanf("%s", &command)

	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	fmt.Println("httpServer stop!")
}
