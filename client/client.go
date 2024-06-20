package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	content := "https://www.youtube.com/watch?v=wRjxtcSaFTc"
	size := "256"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("content", content)
	writer.WriteField("size", size)

	writer.Close()

	resp, err := http.Post("http://localhost:1211/generate", writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println("Sorovni yoborishda xatolik bor!!!!")
		return
	}

	defer resp.Body.Close()

	outFile, err := os.Create("qrcode.png")
	if err != nil {
		fmt.Println("Fayl yaratishda xatolik:", err)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Javobni o'qishda xatolik:", err)
		return
	}
	fmt.Println("Qrcode png faylda saqlandi")
}
