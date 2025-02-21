package main

import (
	"net/http"
	"os"

	"github.com/203-cloud/demo-flux-team-a/demo-app/internal/web"
	flag "github.com/spf13/pflag"
)


func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.Parse()
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello, world!"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "` + message + `"}`))
	})
	web.ServeGraceful(port)
}