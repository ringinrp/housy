package middleware

// Import required packages here ...
import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Create UploadFile function here ...
func UploadFile(next http.HandlerFunc, formImage string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile(formImage)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB

		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max sizes are 10MB"}
			json.NewEncoder(w).Encode(response)
			return
		}

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)

		data := tempFile.Name()
		// filename := data[8:]

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
