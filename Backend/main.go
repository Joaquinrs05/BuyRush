package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "¡Hola desde Go backend!")
	})

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
