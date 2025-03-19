package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtener el nombre del host
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("❌ Error al obtener el nombre del host:", err)
		return
	}

	// Imprimir el nombre del host
	fmt.Println("✅ Nombre del host:", hostName)
}