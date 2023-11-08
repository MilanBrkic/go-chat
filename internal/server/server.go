package server

import (
	"fmt"
	"net/http"
	"os"
)

func Listen() {
	port := os.Getenv("PORT")
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
