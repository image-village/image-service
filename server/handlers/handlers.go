package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/lagbana/images/config"
	"github.com/lagbana/images/data"
	"github.com/lagbana/images/server/responses"
	"github.com/lagbana/images/storage"
	"io"
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

	env := config.EnvSetup()
	bucketName := env.BucketName

	var user UserSession
	json.Unmarshal(u, &user)

	if err := r.ParseMultipartForm(MaxMemory); err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusForbidden, err)
	}

	mfMap := r.MultipartForm.Value

	price, err := strconv.ParseFloat(mfMap["price"][0], 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	title := mfMap["title"][0]

	mf, fh, err := r.FormFile("upfile")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	defer mf.Close()

	// Create file name
	fname, gcpName := createFileName(mf, fh, user.ID)
	// Create path to temp file
	path, err := createFilePath(fname)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	// Create temp file
	err = createFile(mf, path)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// Upload to GCP
	err = storage.UploadFile(bucketName, gcpName, path)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// Delete temp file
	err = deleteFile(path)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// Url to file on GCP, i.e. "https://storage.googleapis.com/BUCKET_NAME/OBJECT_NAME"
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, gcpName)
	record := data.Image{
		Email:  user.Email,
		GcpID:  gcpName,
		Price:  price,
		Title:  title,
		URL:    url,
		UserID: user.ID,
	}
	img, err := data.DB.Images.AddImage(record)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusCreated, fmt.Sprintf("Successfully uploaded new image: %s", img.Title))
}

func createFilePath(fname string) (string, error) {
	var path string
	env := config.EnvSetup()

	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// define file path
	if env.AppEnv == "development" {
		path = filepath.Join(wd, "server", "temp", fname)
	} else {
		path = filepath.Join(wd, fname)
	}

	return path, nil
}

func createFileName(mf multipart.File, fh *multipart.FileHeader, u string) (fileName string, gcpName string) {
	ext := strings.Split(fh.Filename, ".")[1]
	// Create a new sha which we'll use to create a unique file name
	hash := sha256.New()
	// Read through the multi-part file to the end and copy into hash
	io.Copy(hash, mf)
	fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
	// File name on Google Cloud Storage
	gname := fmt.Sprintf("%s/%s", u, fname)

	return fname, gname
}

func createFile(mf multipart.File, path string) error {
	// create new file at defined path
	nf, err := os.Create(path)
	if err != nil {
		return err
	}
	// defer new file close
	defer nf.Close()
	// Reset read-write head to the beginning of the multi-part file
	mf.Seek(0, 0)
	// Copy multi-part file into the new file created on the os
	io.Copy(nf, mf)

	return nil
}

func deleteFile(fname string) error {
	err := os.Remove(fname)
	if err != nil {
		return err
	}
	return nil
}

// ! REMOVE
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responses.JSON(w, http.StatusOK, "Welcome to the image API")
}
