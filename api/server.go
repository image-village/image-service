package api

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Run - launch api
func Run() {
	router := httprouter.New()
	// httprouter.Handle()
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	router.POST("/image", uploadImage)
	http.ListenAndServe(":8080", router)
}

func uploadImage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	mf, fh, errUpload := req.FormFile("upfile")
	if errUpload != nil {
		fmt.Println(errUpload)
	}
	defer mf.Close()

	ext := strings.Split(fh.Filename, ".")[1]
	// Create a new sha which we'll use to create a unique file name
	hash := sha256.New()
	// Read through the multi-part file to the end and copy into hash
	io.Copy(hash, mf)
	fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext

	// Get working directory
	wd, errWd := os.Getwd()
	if errWd != nil {
		fmt.Println(errWd)
	}
	// define file path
	path := filepath.Join(wd, "api", "temp", fname)
	// create new file at defined path
	nf, errNF := os.Create(path)
	if errNF != nil {
		fmt.Println(errNF)
	}
	// defer new file close
	defer nf.Close()
	// Reset read-write head to the beginning of the multi-part file
	mf.Seek(0, 0)
	// Copy multi-part file into the new file created on the os
	io.Copy(nf, mf)

	JSON(res, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "You have successfully uploaded the image",
	})
}

// HandleError - Custom API error handler
func HandleError(resW http.ResponseWriter, err error) {
	if err != nil {
		http.Error(resW, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

// JSON - Shape response data
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR shape error response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	// var errors [] struct {}

	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

/** 
* { errors }
*   errors = [ { message: "string"} ]
*
*/