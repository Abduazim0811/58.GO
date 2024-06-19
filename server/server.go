package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func Qrcode(c *gin.Context){
	content := c.PostForm("content")
	sizeStr := c.PostForm("size")
	if content==""{
		c.JSON(http.StatusBadRequest, gin.H{"error" :"contentni ichi bo'sh"})
	}

	if sizeStr == ""{
		sizeStr="256"
	}

	size, err := strconv.Atoi(sizeStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"size butun son bolishi shart"})
		return
	}

	qr, err := qrcode.New(content, qrcode.Medium)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Qrcode yaratishda xatolik!!!!!!!!!"})
		return
	}

	png, err := qr.PNG(size)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Qrcode ni png ga o'tqazishda xatolik bor!!!!!!!!!!!!!!"})
		return
	}
	
	c.Data(http.StatusOK, "image/png", png)
}

func main(){
	r := gin.Default()
	r.POST("/generete", Qrcode)
	r.Run(":1211")
}
