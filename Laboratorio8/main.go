package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageData struct {
	Name   string
	Base64 string
}

type PageData struct {
	Title       string
	Hostname    string
	Images      []ImageData
	Directorio1 []ImageData
	Directorio2 []ImageData
	CourseInfo  string
	Team        string
	Port        string
	ImageDir    string
	ImageDir2   string
}

func main() {
	port := "5001"
	imageDir := `C:\Users\lenovo\Desktop\nube\Goolang\Nube\archivos`
	imageDir2 := `C:\Users\lenovo\Desktop\nube\Goolang\Nube\archivos2`

	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	if len(os.Args) > 2 {
		imageDir = os.Args[2]
	}

	rand.Seed(time.Now().UnixNano())

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "No identificado"
	}

	printServerInfo(hostname, port, imageDir, imageDir2)

	// Servir imágenes estáticas
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(imageDir))))
	http.Handle("/static2/", http.StripPrefix("/static2/", http.FileServer(http.Dir(imageDir2))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleHome(w, r, hostname, port, imageDir, imageDir2)
	})

	startServer(port)
}

func printServerInfo(hostname, port, imageDir, imageDir2 string) {
	fmt.Println("Servidor de imágenes iniciado en:")
	fmt.Printf("Host: %s\nPuerto: %s\nDirectorio 1: %s\nDirectorio 2: %s\n", hostname, port, imageDir, imageDir2)
}

func handleHome(w http.ResponseWriter, r *http.Request, hostname, port, imageDir, imageDir2 string) {
	imageFiles, err := getRandomImages(imageDir, 4)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allImages, err := getFixedImages(imageDir, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allImages2, err := getFixedImages(imageDir2, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error cargando la plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:       "Servidor de Imágenes",
		Hostname:    hostname,
		Images:      imageFiles,
		Directorio1: allImages,
		Directorio2: allImages2,
		CourseInfo:  "Computación en la nube 2025 - 1",
		Team:        "Alejandro Aguirre. -Kevin Buitron. -Kevin soto.",
		Port:        port,
		ImageDir:    imageDir,
		ImageDir2:   imageDir2,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla: "+err.Error(), http.StatusInternalServerError)
	}
}

func getFixedImages(dir string, count int) ([]ImageData, error) {
	var allImages []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("No se pudo leer el directorio: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				allImages = append(allImages, file.Name())
			}
		}
	}

	if len(allImages) == 0 {
		return nil, fmt.Errorf("No se encontraron imágenes en el directorio")
	}

	if len(allImages) < count {
		count = len(allImages)
	}

	var result []ImageData
	for i := 0; i < count; i++ {
		result = append(result, ImageData{
			Name:   allImages[i],
			Base64: "/static/" + allImages[i],
		})
	}

	return result, nil
}

func getRandomImages(dir string, count int) ([]ImageData, error) {
	var allImages []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("No se pudo leer el directorio: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				allImages = append(allImages, file.Name())
			}
		}
	}

	if len(allImages) == 0 {
		return nil, fmt.Errorf("No se encontraron imágenes en el directorio")
	}

	if len(allImages) < count {
		count = len(allImages)
	}

	rand.Shuffle(len(allImages), func(i, j int) {
		allImages[i], allImages[j] = allImages[j], allImages[i]
	})

	var result []ImageData
	for i := 0; i < count; i++ {
		result = append(result, ImageData{
			Name:   allImages[i],
			Base64: "/static/" + allImages[i],
		})
	}

	return result, nil
}

func startServer(port string) {
	fmt.Printf("Servidor en ejecución en: http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}
