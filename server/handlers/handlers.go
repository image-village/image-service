package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/lagbana/images/config"
	"github.com/lagbana/images/server/responses"
	// "github.com/lagbana/images/storage"
	// "github.com/lagbana/images/db/data"
	"io"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MaxMemory = 10MB
const MaxMemory = 10 * 1024 * 1024

// UserSession - user session object
type UserSession struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// UploadImage -
func UploadImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := r.Header.Get("currentUser")
	u := []byte(t)

	var user UserSession
	json.Unmarshal(u, &user)

	if err := r.ParseMultipartForm(MaxMemory); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	data := r.MultipartForm.Value
	title := data["title"][0]
	price, err := strconv.ParseFloat(data["price"][0], 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(price)
	fmt.Println(title)
	fmt.Printf("🚀 %T\n", price)

	mf, fh, errUpload := r.FormFile("upfile")
	if errUpload != nil {
		fmt.Println(errUpload)
	}
	defer mf.Close()

	// Create file name
	fname := createFileName(mf, fh)
	// Get working directory
	path := createFilePath(fname)


	createFile(mf, path)

	// // Create GCP bucket
	// storage.CreateBucket(user.Email)
	// // Upload to GCP
	// err = storage.UploadFile(w, user.Email, fname, path)
	// if err != nil {
	// 	log.Println(err)
	// }
	// // Url to file in GCP, i.e. "https://storage.googleapis.com/BUCKET_NAME/OBJECT_NAME"
	// url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", user.Email, fname)


}

func createFilePath(fname string) string {
	var path string
	env := config.EnvSetup()

	// Get working directory
	wd, errWd := os.Getwd()
	if errWd != nil {
		fmt.Println(errWd)
	}

	// define file path
	if env.AppEnv == "development" {
		path = filepath.Join(wd, "server", "temp", fname)
	} else {
		path = filepath.Join(wd, fname)
	}

	return path
}

func createFileName(mf multipart.File, fh *multipart.FileHeader) string {
	ext := strings.Split(fh.Filename, ".")[1]
	// Create a new sha which we'll use to create a unique file name
	hash := sha256.New()
	// Read through the multi-part file to the end and copy into hash
	io.Copy(hash, mf)
	fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext

	return fname
}

func createFile(mf multipart.File, path string) {
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
}

// ! REMOVE
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responses.JSON(w, http.StatusOK, "Welcome to the image API")
}
