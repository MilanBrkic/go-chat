package server

import (
	"fmt"
	"go-chat/internal/config"
	"net/http"
)

func Listen() {
	port := config.SERVER_PORT
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     w.Write([]byte("Hello, World!"))
	// })

	server := &http.Server{Addr: ":" + port}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Server listening on port %s...\n", port)
}
