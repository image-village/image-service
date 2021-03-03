package data

import (
	"fmt"
	"github.com/upper/db/v4"
)

// Image represenets a record from the "images" table.
type Image struct {
	Email  string  `db:"email"`
	GcpID  string  `db:"gcp_id"`
	ID     uint    `db:"id,omitempty"`
	Price  float64 `db:"price"`
	Title  string  `db:"title"`
	URL    string  `db:"url"`
	UserID string  `db:"user_id"`
}

// ImagesStore represents a store for images
type ImagesStore struct {
	db.Collection
}

// GetImageByID - get one image by its ID
func (images *ImagesStore) GetImageByID(id int64) (*Image, error) {
	var image Image
	if err := images.Find(db.Cond{"id": id}).One(&image); err != nil {
		return nil, err
	}
	return &image, nil
}

// GetImagesByUser - get all user images
func (images *ImagesStore) GetImagesByUser(userID string) (*[]Image, error) {
	var imgs []Image
	res := images.Find(db.Cond{"user_id": userID}).OrderBy("-title")
	if err := res.All(&imgs); err != nil {
		return nil, err
	}
	return &imgs, nil
}

// GetAllImages - get all images
func (images *ImagesStore) GetAllImages(userID string) (*[]Image, error) {
	var imgs []Image
	res := images.Find(db.Cond{"user_id": userID}).OrderBy("-title")
	if err := res.All(&imgs); err != nil {
		return nil, err
	}
	return &imgs, nil
}

// AddImage - get all images
func (images *ImagesStore) AddImage(image Image) (*Image, error) {
	err := images.InsertReturning(&image)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

// Images initializes an ImageStore
func Images(sess db.Session) *ImagesStore {
	return &ImagesStore{sess.Collection("images")}
}

// Store returns images collection
func (image *Image) Store(sess db.Session) db.Store {
	return Images(sess)
}

// BeforeUpdate event driven hook
func (image *Image) BeforeUpdate(sess db.Session) error {
	fmt.Println("**** BeforeUpdate was called ****")
	return nil
}

// AfterUpdate event driven hook
func (image *Image) AfterUpdate(sess db.Session) error {
	fmt.Println("**** AfterUpdate was called ****")
	return nil
}

// Compile-time interface checks
// i.e. ImageStore satisfies db.Store interface
var _ = interface{ db.Store }(&ImagesStore{})

var _ = interface{ db.Record }(&Image{})
