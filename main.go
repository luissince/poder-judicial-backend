package main

import (
	"api-pdf/helper"
	"api-pdf/pdf"
	"fmt"
	_ "image/jpeg" // Importa el formato JPEG
	_ "image/png"  // Importa el formato PNG
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// type Data struct {
// 	nombreSistema  string
// 	versionSistema string
// 	usuarioNombre  string
// 	celularAxeso   string
// 	sede           string
// 	cargo          string
// 	personaReporte string
// 	celularPersona string
// 	fecha          string
// 	descripcion    string
// 	descartes      string
// }

func main() {

	// Cargar las variables de entorno desde el archivo .env
	godotenv.Load()

	// Obtener el valor de la variable de entorno GO_PORT
	var go_port string = os.Getenv("GO_PORT")

	router := gin.Default()

	router.Use(cors.Default())

	// Middleware para manejar el error 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Ruta no encontrada."})
	})

	// Rutas
	router.GET("/", func(c *gin.Context) {
		log.Println("Endpoint ping")

		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG APP MOVIL UPLA",
		})
	})

	router.GET("/base64", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "BASE 64 CONVERTOR",
		})
	})
	router.POST("/pdf", handlePDFRequestGin)

	err := router.Run(go_port)
	if err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}

	// http.HandleFunc("/pdf", handlePDFRequest)
	// fmt.Println("Run server on http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePDFRequestGin(c *gin.Context) {
	var data pdf.Data

	if err := c.BindJSON(&data); err != nil {
		// c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	helper.CrearCarpeta("tmp/")

	var count int
	for _, item := range data.Imagenes {
		base64String := "data:image/" + item.Extension + ";base64," + item.Base64String

		imageData := helper.ExtractImageData(base64String)
		if imageData == nil {
			fmt.Println("No se pudo extraer la imagen base64")
			return
		}

		imageType := helper.ExtractImageType(base64String)
		if imageType == "" {
			fmt.Println("No se pudo determinar el tipo de imagen")
			return
		}

		count++
		err := helper.SaveImage(imageData, imageType, "tmp/"+strconv.Itoa(count)+"output."+item.Extension)
		if err != nil {
			fmt.Println("Error al guardar la imagen:", err)
			return
		}
	}

	pdfBytes, err := pdf.CrearPdf(data)
	if err != nil {
		fmt.Println("Error en crear el pdf:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	pdfMemory, err := ioutil.ReadAll(pdfBytes)
	if err != nil {
		fmt.Println("Error en leer la imagen:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Envía el flujo de bytes del PDF como respuesta
	c.Header("Content-Type", "application/pdf")
	// c.Header("Content-Disposition", "attachment; filename=pdf.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfMemory)
}

/*

func handlePDFRequest(w http.ResponseWriter, r *http.Request) {

	pdfBytes, err := pdf.CrearPdf()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Envía el flujo de bytes del PDF al cliente
	w.Header().Set("Content-Type", "application/pdf")
	// w.Header().Set("Content-Disposition", "attachment; filename=pdf.pdf")
	_, err = w.Write([]byte(pdfBytes.String()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
*/
