package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Pedir la ruta de la imagen al usuario
	fmt.Print("Ingrese la ruta de la imagen: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	imagePath := scanner.Text()

	// Leer el archivo de imagen
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println("❌ Error al leer la imagen:", err)
		return
	}

	// Codificar la imagen en Base64
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Mostrar la imagen codificada
	fmt.Println("\n✅ Imagen codificada en Base64:")
	fmt.Println(encodedImage)

	// Guardar la imagen codificada en un archivo de texto
	outputFile := "imagen_base64.txt"
	err = ioutil.WriteFile(outputFile, []byte(encodedImage), 0644)
	if err != nil {
		fmt.Println("❌ Error al guardar la imagen codificada:", err)
		return
	}

	fmt.Println("\n💾 La imagen codificada se ha guardado en:", outputFile)
}