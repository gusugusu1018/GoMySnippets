package main

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var mux *http.ServeMux
	mux = http.NewServeMux()
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/", index)
	var server *http.Server
	server = &http.Server{}
	server.Addr = ":11180"
	server.Handler = mux
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	var t *template.Template
	t, _ = template.ParseFiles("template/index.html")
	t.Execute(writer, struct{}{})
}
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Cannot Upload. Only POST Method.")
		return
	}
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var e error
	var uploadedFileName string
	file, fileHeader, e = r.FormFile("image")
	if e != nil {
		fmt.Fprintln(w, "Can't get image. ")
		return
	}
	uploadedFileName = fileHeader.Filename
	var saveImage *os.File
	saveImage, e = os.Create("./" + uploadedFileName)
	if e != nil {
		fmt.Fprintln(w, "Can't os.Create() for save image. ")
		return
	}
	defer saveImage.Close()
	defer file.Close()
	size, e := io.Copy(saveImage, file)
	if e != nil {
		fmt.Println(e)
		fmt.Println("Can't io.Copy(). ")
		os.Exit(1)
	}
	fmt.Println(size)
	fmt.Fprintf(w, "Success")
}
