package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

// nombre del flag, valor por defecto, descripcion
var site = flag.String("site", "scanme.nmap.org", "URL del sitio a escannear")

//go run net/port.go --site=scanme.webscantest.com

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//checking ports
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", *site, port),
			)

			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(i)

	}

	wg.Wait()
}
