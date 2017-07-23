// GoTV - watching Weiqi like TV
package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	port = flag.Int("port", 3217, "server port")
	host = flag.String("host", "localhost", "server host")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("start %s\n", addr)
}
