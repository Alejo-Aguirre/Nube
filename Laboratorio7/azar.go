package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Directorio que se va a listar
	//directorio := "C:/Users/lenovo/Desktop/nube/Goolang/Nube"

	var directorio string
	fmt.Print("Ingrese la ruta del directorio: ")
	fmt.Scanln(&directorio)
	
	// Formatos de archivo permitidos (solo imágenes en este caso)
	formatosPermitidos := []string{".jpg", ".png"}

	// Llamar a la función para listar archivos y obtener los nombres en un slice
	nombresArchivos := listarArchivosAzar(directorio, formatosPermitidos)

	// Verificar si hay archivos de imagen
	if len(nombresArchivos) == 0 {
		fmt.Println("No se encontraron archivos de imagen en el directorio.")
		return
	}

	// Seleccionar un nombre de archivo al azar
	rand.Seed(time.Now().UnixNano()) // Inicializar la semilla aleatoria
	indiceAleatorio := rand.Intn(len(nombresArchivos)) // Generar un índice aleatorio
	archivoSeleccionado := nombresArchivos[indiceAleatorio]

	// Imprimir solo el nombre del archivo seleccionado al azar
	fmt.Println(archivoSeleccionado)
}

func listarArchivosAzar(directorio string, formatosPermitidos []string) []string {
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

	// Retornar el slice con los nombres de los archivos
	return nombresArchivos
}