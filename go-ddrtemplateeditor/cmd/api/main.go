package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-ddrtemplateeditor/pkg/db"
	"image"
	"image/png"
	"log"
	"models"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type templateDto struct {
	ID    int              `json:"id"`
	Name  string           `json:"name"`
	Links []hypermediaLink `json:"links"`
}

type hypermediaLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

var templateFolder = "../../samples/templates"

func main() {
	router := mux.NewRouter()

	db.Create()
	defer db.DropDb()

	errChan := make(chan error)
	go func() {
		errChan <- loadTemplates()
	}()

	// Define API endpoints
	router.HandleFunc("/templates", getTemplates).Methods("GET")
	router.HandleFunc("/templates/{id}", getTemplate).Methods("GET")
	router.HandleFunc("/templates/{id}/image", getTemplateImage).Methods("GET")
	// router.HandleFunc("/templates", createTemplate).Methods("POST")
	// router.HandleFunc("/templates/{id}", updateTemplate).Methods("PUT")
	// router.HandleFunc("/templates/{id}", deleteTemplate).Methods("DELETE")
	router.Handle("/", http.FileServer(http.Dir(templateFolder)))

	// Wait for templates to load
	err := <-errChan
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":1337", router))
}

func loadTemplates() error {
	files, err := os.ReadDir(templateFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		f, err := os.Open(templateFolder + "/" + file.Name())
		if err != nil {
			return err
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}

		var buffer bytes.Buffer
		err = png.Encode(&buffer, img)
		if err != nil {
			return err
		}

		db.Insert(file.Name(), buffer.Bytes())
	}
	return nil
}

func mapToDto(t ...*models.Template) []templateDto {
	r := []templateDto{}
	for _, t := range t {
		r = append(r, templateDto{ID: t.ID, Name: t.Name, Links: []hypermediaLink{
			{Rel: "self", Href: fmt.Sprintf("/templates/%d", t.ID)},
			{Rel: "self", Href: fmt.Sprintf("/templates/%d/image", t.ID)}}})
	}

	return r
}

func getTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := db.QueryTemplates()
	if err != nil {
		log.Fatal(err)
	}

	rsp := mapToDto(templates...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}

func getTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	t, err := db.QueryTemplate(id)
	if err != nil {
		log.Fatal(err)
	}

	rsp := mapToDto(t)

	json.NewEncoder(w).Encode(rsp)
}

func getTemplateImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	t, err := db.QueryTemplate(id)
	if err != nil {
		log.Fatal(err)
	}

	writeImage(w, &t.Img)

	w.Header().Set("Content-Type", "image/png")
	json.NewEncoder(w).Encode(nil)
}

// func createTemplate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var template Template
// 	_ = json.NewDecoder(r.Body).Decode(&template)
// 	templates = append(templates, template)
// 	json.NewEncoder(w).Encode(template)
// }

// func updateTemplate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, template := range templates {
// 		if template.ID == params["id"] {
// 			templates = append(templates[:index], templates[index+1:]...)
// 			var updatedTemplate Template
// 			_ = json.NewDecoder(r.Body).Decode(&updatedTemplate)
// 			updatedTemplate.ID = params["id"]
// 			templates = append(templates, updatedTemplate)
// 			json.NewEncoder(w).Encode(updatedTemplate)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(nil)
// }

// func deleteTemplate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, template := range templates {
// 		if template.ID == params["id"] {
// 			templates = append(templates[:index], templates[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(nil)
// }

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
