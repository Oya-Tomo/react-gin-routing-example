package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8000"
	engin := gin.Default()
	engin.LoadHTMLGlob("./build/*.html")
	engin.Static("/static", "./build/static")

	engin.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	engin.GET("/:file", func(ctx *gin.Context) {
		filename := ctx.Param("file")
		exists := FileExists("./build/" + filename)
		if !exists {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		} else {
			fmt.Println("[INFO] " + "./build/" + filename)
			ctx.File("./build/" + filename);
		}
	})

	fmt.Println("http://localhost" + port)
	engin.Run(port)
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}