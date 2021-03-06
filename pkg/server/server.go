package server

import (
	// "crypto/tls"
	"os"
	"log"
	"net/http"
	"time"

	"github.com/kangbojk/go-react-fullstack/pkg/server/router"
	"github.com/kangbojk/go-react-fullstack/pkg/usecase"
)

func NewServer(service usecase.Service) *http.Server {
	handler := router.NewRouter(service)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("start server on port: ", port)

	return s
}

func Run() {
}

// func tlsConfig() *tls.Config {
// 	crt, err := ioutil.ReadFile("./public.crt")
// 	if err != nil {
// 			log.Fatal(err)
// 	}

// 	key, err := ioutil.ReadFile("./private.key")
// 	if err != nil {
// 			log.Fatal(err)
// 	}

// 	cert, err := tls.X509KeyPair(crt, key)
// 	if err != nil {
// 			log.Fatal(err)
// 	}

// 	return &tls.Config{
// 			Certificates: []tls.Certificate{cert},
// 			ServerName:   "localhost",
// 	}
// }

// func NewTLSServer() {
// 	server := &http.Server{
// 			Addr:         ":8443",
// 			ReadTimeout:  5 * time.Second,
// 			WriteTimeout: 10 * time.Second,
// 			TLSConfig:    tlsConfig(),
// 	}

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte(fmt.Sprintf("Protocol: %s", r.Proto)))
// 	})

// 	if err := server.ListenAndServeTLS("", ""); err != nil {
// 			log.Fatal(err)
// 	}
// }
