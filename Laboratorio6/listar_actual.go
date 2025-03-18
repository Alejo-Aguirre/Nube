package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtener el directorio actual
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error al obtener el directorio actual:", err)
		return
	}

	// Leer el contenido del directorio
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	// Listar los archivos y directorios
	fmt.Println("Archivos en el directorio actual:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}