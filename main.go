package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"net"
)

func getValidIP() string {
	netInterfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("net.Interfaces failed, err:", err.Error())
        return ""
	}
 
    for i := 0; i < len(netInterfaces); i++ {
        if (netInterfaces[i].Flags & net.FlagUp) != 0 {
            addrs, _ := netInterfaces[i].Addrs()
 
            for _, address := range addrs {
                if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                    if ipnet.IP.To4() != nil {
                        return ipnet.IP.String()
                    }
                }
            }
        }
	}
	return ""
}

func  main() {

	ip:=getValidIP()
	if ip == "" {
		return
	}
 
	address := flag.String("a", "[::]", "address")
	port := flag.Int("p", 8080, "port")
	flag.Parse()

	var httpAddress string 
	httpAddress= fmt.Sprintf("%s:%v",ip, *port)

	if *address !="[::]"{
		httpAddress= fmt.Sprintf("%s:%v",*address, *port)
	}

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
