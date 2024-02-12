package main

import (
	"bytes"
	"encoding/json"
	"go-ddrtemplateeditor/internal/assets"
	"go-ddrtemplateeditor/pkg/db"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type replacementDto struct {
	TemplateID int `json:"templateId"`
	Items      []struct {
		ID int `json:"id"`
	} `json:"items"`
}

var templateFolder = "../../samples/templates"

var itemsSetFile = "../../samples/items/items.yml"

var dbInstance *db.DB

func main() {
	router := mux.NewRouter()

	d, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer d.DropDb()
	dbInstance = d

	err = loadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	err = loadItems()
	if err != nil {
		log.Fatal(err)
	}

	// Define API endpoints
	router.HandleFunc("/api/templates", getTemplates).Methods("GET")
	router.HandleFunc("/api/templates/{id}", getTemplate).Methods("GET")
	router.HandleFunc("/api/templates/{id}/image", getTemplateImage).Methods("GET")
	router.HandleFunc("/api/templates/{id}/replace", updateTemplate).Methods("PUT")
	router.HandleFunc("/api/items", getItems).Methods("GET")

	log.Fatal(http.ListenAndServe(":1337", router))
}

func loadItems() error {
	file, err := os.Open(itemsSetFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var items []db.Item
	err = yaml.NewDecoder(file).Decode(&items)
	if err != nil {
		return err
	}

	for _, item := range items {
		dbInstance.CreateItem(&item)
	}

	return nil
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

		dbInstance.CreateTemplate(&db.Template{Name: file.Name(), Img: buffer.Bytes()})
	}
	return nil
}

func getTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := dbInstance.QueryTemplates()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templates)
}

func getTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	t, err := dbInstance.QueryTemplate(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(t)
}

func getTemplateImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	t, err := dbInstance.QueryTemplate(id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(t.Img)))
	if _, err := w.Write(t.Img); err != nil {
		log.Println("unable to write image.")
	}
	json.NewEncoder(w).Encode(nil)
}

func updateTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	t1, err := dbInstance.QueryTemplate(id)
	if err != nil {
		log.Fatal(err)
	}

	var replacement replacementDto
	_ = json.NewDecoder(r.Body).Decode(&replacement)
	t2, err := dbInstance.QueryTemplate(replacement.TemplateID)
	if err != nil {
		log.Fatal(err)
	}

	// Load the template images
	dstT := assets.NewTemplate(createImage(t1))
	srcT := assets.NewTemplate(createImage(t2))

	// Load the items
	items := []*assets.Item{}
	for _, item := range replacement.Items {
		i, err := dbInstance.QueryItem(item.ID)
		if err != nil {
			log.Fatal(err)
		}
		item := srcT.NewItem(i.X, i.Y, i.Width, i.Height)
		items = append(items, item)
	}

	// Replace the items
	dstT.ReplaceItems(items...)

	// Save the new template
	err = dbInstance.UpdateTemplateImage(id, dstT.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(nil)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.QueryItems()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createImage(t *db.Template) image.Image {
	img, _, err := image.Decode(bytes.NewReader(t.Img))
	if err != nil {
		log.Fatal(err)
	}

	return img
}
