package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Directorio que se va a listar
	var directorio string
	fmt.Print("Ingrese la ruta del directorio: ")
	fmt.Scanln(&directorio)

	// Formatos de archivo permitidos
	formatosPermitidos := []string{".jpg", ".png", ".jpeg"}

	// Llamar a la función para listar archivos
	listarArchivosNombre(directorio, formatosPermitidos)

}

func listarArchivosNombre(directorio string, formatosPermitidos []string) {
	// Recorrer el directorio usando filepath.Walk
	err := filepath.Walk(directorio, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error al acceder al directorio:", err)
			return nil
		}

		// Verificar si es un archivo (no un directorio)
		if !info.IsDir() {
			// Obtener la extensión del archivo
			ext := strings.ToLower(filepath.Ext(info.Name()))

			// Verificar si la extensión está en la lista de formatos permitidos
			for _, formato := range formatosPermitidos {
				if ext == formato {
					fmt.Println(info.Name()) // Imprimir solo el nombre del archivo
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error al recorrer el directorio:", err)
	}
}
