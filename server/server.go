package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func generateQRCode(c *gin.Context) {
	content := c.PostForm("content")
	sizeStr := c.PostForm("size")

	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content parametri kiritilmadi"})
		return
	}

	if sizeStr == "" {
		sizeStr = "256"
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size parametri butun son bo'lishi kerak"})
		return
	}

	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "QR kod yaratishda xatolik: " + err.Error()})
		return
	}

	png, err := qr.PNG(size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "QR kodni PNG formatida yaratishda xatolik: " + err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}

func main() {
	r := gin.Default()
	r.POST("/generate", generateQRCode)

	r.Run(":1211")
}
