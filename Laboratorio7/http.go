package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Definir la ruta y el manejador para la página de prueba
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Escribir el mensaje "Hola Mundo!" en la respuesta
		fmt.Fprintf(w, "Hola Mundo!")
	})

	// Iniciar el servidor HTTP en el puerto 5001
	fmt.Println("Servidor HTTP iniciado en http://localhost:5001")
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}