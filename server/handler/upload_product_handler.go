package handler

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadHandler(c *gin.Context) {
	var upload models.ProductUpload
	err := c.ShouldBind(&upload)
	tools.CheckError(err, "failed to bind upload product data")

	imageFile, err := upload.Image.Open()
	tools.CheckError(err, "failed to open image product")

	defer func(imageFile multipart.File) {
		err = imageFile.Close()
		tools.CheckError(err, "failed to close image product")
	}(imageFile)

	techSpec, err := upload.TechnicalSpec.Open()
	tools.CheckError(err, "failed to open technical spec product")

	defer func(techSpec multipart.File) {
		err = techSpec.Close()
		tools.CheckError(err, "failed to close technical spec product")
	}(techSpec)

	installGuide, err := upload.InstallationGuide.Open()
	tools.CheckError(err, "failed to open installation guide product")

	defer func(installGuide multipart.File) {
		err = installGuide.Close()
		tools.CheckError(err, "failed to close installation guide product")
	}(installGuide)

	path := createPath(makeHash())
	imageFilePath := path + "/" + getInfoFile(c, "image")
	techSpecFilePath := path + "/" + getInfoFile(c, "technical_spec")
	installGuideFilePath := path + "/" + getInfoFile(c, "installation_guide")

	err = c.SaveUploadedFile(upload.Image, imageFilePath)
	tools.CheckError(err, "failed to save image upload in product")

	err = c.SaveUploadedFile(upload.TechnicalSpec, techSpecFilePath)
	tools.CheckError(err, "failed to save image upload in product")

	err = c.SaveUploadedFile(upload.InstallationGuide, installGuideFilePath)
	tools.CheckError(err, "failed to save image upload in product")
}

func getInfoFile(c *gin.Context, name string) string {
	file, err := c.FormFile(name)
	tools.CheckError(err, "failed to get name and ext file in product")
	fileName := file.Filename

	return fileName
}

func createPath(hash string) string {
	baseDir := "uploads/products"

	productDir := filepath.Join(baseDir, hash)

	err := os.MkdirAll(productDir, os.ModePerm)
	tools.CheckError(err, "failed to create products folder")

	return productDir
}

func makeHash() string {
	currentTime := time.Now()
	timeString := currentTime.String()

	layout := "2006-01-02 15:04:05"

	formattedTime := currentTime.Format(layout)

	hash := sha256.Sum256([]byte(timeString))
	hashString := hex.EncodeToString(hash[:])

	return formattedTime + " | " + hashString
}
