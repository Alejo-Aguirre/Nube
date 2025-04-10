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

	// Formatos de archivo permitidos (solo imágenes en este caso)
	formatosPermitidos := []string{".jpg", ".png", ".jpeg"}

	// Llamar a la función para listar archivos y obtener los nombres en un slice
	nombresArchivos := listarArchivos(directorio, formatosPermitidos)

	// Determinar la cantidad de imágenes disponibles
	cantidadImagenes := len(nombresArchivos)

	// Imprimir los nombres de los archivos y la cantidad de imágenes
	fmt.Println("Nombres de los archivos de imagen:")
	for _, nombre := range nombresArchivos {
		fmt.Println(nombre)
	}
	fmt.Printf("Cantidad de imágenes disponibles: %d\n", cantidadImagenes)

}

func listarArchivos(directorio string, formatosPermitidos []string) []string {
	// Slice para almacenar los nombres de los archivos
	var nombresArchivos []string

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
					// Agregar el nombre del archivo al slice
					nombresArchivos = append(nombresArchivos, info.Name())
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error al recorrer el directorio:", err)
	}
	return nombresArchivos
}