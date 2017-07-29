// GoTV - watching Weiqi like TV
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dgf/gotv/memory"
	"github.com/dgf/gotv/web"
)

var (
	port = flag.Int("port", 3217, "server port")
	host = flag.String("host", "localhost", "server host")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Fatal(http.ListenAndServe(addr, web.New("GoTV", memory.New())))
}
