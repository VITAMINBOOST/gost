package gostaudio

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	UploadPath = "./mediafiles/audio/upload/"
)

func SingleUploadFromClient(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	log.Println(file.Filename)

	filename := filepath.Base(file.Filename)

	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		os.MkdirAll(UploadPath, 0755)
	}

	uploadPath := UploadPath + filename
	log.Println(filename)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"status":    "posted",
		"file name": file.Filename,
	})
}

func MultiUploadFromClient(c *gin.Context) {
	firstName := c.PostForm("first_name")
	familyName := c.PostForm("family_name")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["file"]
	log.Println(form)

	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		os.MkdirAll(UploadPath, 0755)
	}

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		filename := filepath.Base(file.Filename)
		uploadPath := UploadPath + filename

		// Upload the file to specific dst.
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	c.JSON(200, gin.H{
		"status":      "posted",
		"file lise":   files,
		"file count":  len(files),
		"first name":  firstName,
		"family name": familyName,
	})
}
