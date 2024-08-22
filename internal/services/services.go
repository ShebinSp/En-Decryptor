package services

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
//	"os"
	"strconv"
)

type imageData struct {
	id       int
	fileName string
	format   string
	data     []byte
	aecKey   []byte
	hash     string
}

type datas struct {
	images []imageData
}

// Type constraint `newType` is used as go generic function
type newType interface {
	int | string
}

func (d *datas) addData(img imageData) {
	d.images = append(d.images, img)
//	fmt.Println("Size of datas: ", len(d.images)) //
}

var data datas

func EncodeImage(w http.ResponseWriter, r *http.Request) {
	// to store the image data
	var imgData imageData

	// Parse the multipart form data
	r.ParseMultipartForm(int64(10 << 20)) // limit size to 10 MB

	// Retrive the file from form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "error retriving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Get the file name
	imgData.fileName = handler.Filename

	img, format, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode the image", http.StatusBadRequest)
		return
	}

	imgData.format = format

	// buffer to store the bytes
	var buf bytes.Buffer

	// Encode the image to bytes
	switch format {
	case "jpeg":
		err = jpeg.Encode(&buf, img, nil)
	case "png":
		err = png.Encode(&buf, img)
	default:
		http.Error(w, "unsupported image format", http.StatusBadRequest)
	}

	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}

	imgData.id = generateID()
	imgData.encryptAES(buf.Bytes())

	// Adding to image data
	data.addData(imgData)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(imgData.hash))

}

func DecodeImage(w http.ResponseWriter, r *http.Request) {
	var imgData imageData

	//key := r.PostFormValue("key")
	key := r.URL.Query().Get("key")
	if len(key) < 64 {
		http.Error(w, "Invalid key length", http.StatusBadRequest)
		return
	}
	fmt.Println("Key: ", key)

	hash := key[:64]
	idStr := key[64:]
	fmt.Println("idStr: ", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id convertion failed", http.StatusBadRequest)
	}

	// Get the datas of the image using id or hash by utlizing type conversion in go
	imgData, err = getImageData(id)
	if err != nil {
		// If id lookup fails, hash is used
		imgData, err = getImageData(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
	fmt.Println("ID: ", id)
	fmt.Println("ImageData.Name: ", imgData.fileName, "ID: ", imgData.id)

	imgByte, err := imgData.decryptAES()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	// ---Uncomment to save locally---
	// err = generateImage(imgData, imgByte)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/"+imgData.format)
	w.Write(imgByte)
}

// getImageData takes of newType as parameter, the `T` can be either int or string(go generics)
func getImageData[T newType](nT T) (imageData, error) {
	var imgData imageData

	// Check if T is int
	if id, ok := any(nT).(int); ok {
		for i, image := range data.images {
			if image.id == id {
				imgData = data.images[i]
				return imgData, nil
			}
		}
	}

	// Check if T is string(hash)
	if hash, ok := any(nT).(string); ok {
		for i, image := range data.images {
			if image.hash == hash {
				imgData = data.images[i]
				return imgData, nil
			}
		}
	}
	return imgData, fmt.Errorf("no record found with id %v or hash %v", nT, nT)
}

//    <----UNCOMMENT TO SAVE LOCALLY----->

// func decode(data []byte) (image.Image, error) {
// 	img, _, err := image.Decode(bytes.NewBuffer(data))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to decode the image: %w", err)
// 	}
// 	return img, err
// }

// func generateImage(imgData imageData, imgByte []byte) error {

// 	switch imgData.format {
// 	case "jpeg":
// 		image, err := decode(imgByte)
// 		if err != nil {
// 			return err
// 		}
// 		err = createImage(imgData, image)
// 		if err != nil {
// 			return err
// 		}

// 	case "png":
// 		image, err := decode(imgByte)
// 		if err != nil {
// 			return err
// 		}
// 		err = createImage(imgData, image)
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func createImage(data imageData, image image.Image) error {
// 	file, err := os.OpenFile(
// 		"./internal/images/"+data.fileName, // set file path and name accordingly
// 		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
// 		0644,
// 	)
// 	if err != nil {
// 		fmt.Println("error")
// 		return err
// 	}

// 	defer file.Close()

// 	switch data.format {
// 	case "jpeg":
// 		err := jpeg.Encode(file, image, nil)
// 		if err != nil {
// 			return err
// 		}

// 	case "png":
// 		err := png.Encode(file, image)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
