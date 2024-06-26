package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Object struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Type     string
	Data     string
	FolderID uint `gorm:"not null"`
}

func randomFileName() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	buf := make([]byte, 512)

	_, err = file.Read(buf)
	if err != nil {
		return
	}
	_, _ = file.Seek(0, 0)

	contentType := http.DetectContentType(buf)

	fmt.Println("contentType", contentType)

	if contentType != "image/png" && contentType != "image/jpeg" && contentType != "application/pdf" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Save file to disk

	filePath := "uploads/" + randomFileName() + fileHeader.Filename
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpObject := Object{Name: fileHeader.Filename, Type: contentType, Data: filePath, FolderID: 1}

	result := db.Create(&tmpObject)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "File uploaded successfully")
}

func createObject(w http.ResponseWriter, r *http.Request) {
	var jsonObj Object
	err := json.NewDecoder(r.Body).Decode(&jsonObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&jsonObj)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getObject(w http.ResponseWriter, r *http.Request) {
	var jsonObj Object
	id, _ := url.PathUnescape(r.PathValue("id"))

	result := db.First(&jsonObj, "id = ?", id)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "JSON object not found", http.StatusNotFound)
		return
	}

	if jsonObj.Type != "MD" {
		jsonObj.Data = ""
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonObj)
	if err != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

}

func getFileFromObject(w http.ResponseWriter, r *http.Request) {
	var jsonObj Object
	id, _ := url.PathUnescape(r.PathValue("id"))
	fmt.Println(id)

	result := db.First(&jsonObj, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "JSON object not found", http.StatusNotFound)
		return
	}

	if jsonObj.Type == "MD" {
		http.Error(w, "You can only get a File from an File lol. U looser", http.StatusBadRequest)
		return
	}

	filePath := filepath.FromSlash(jsonObj.Data)
	http.ServeFile(w, r, filePath)

}

func updateObject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateObject")

	var jsonObj Object
	id, _ := url.PathUnescape(r.PathValue("id"))

	result := db.First(&jsonObj, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "JSON object not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&jsonObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if jsonObj.Type != "MD" {
		http.Error(w, "Wrong Object Type. Only MD-Files can be updated", http.StatusBadRequest)
		return
	}

	fmt.Println(jsonObj)
	fmt.Println(&jsonObj)

	result2 := db.Save(&jsonObj)
	if result2.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateObjectName(w http.ResponseWriter, r *http.Request) {
	var object Object
	id, _ := url.PathUnescape(r.PathValue("id"))

	result := db.First(&object, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "Object not found", http.StatusNotFound)
		return
	}

	var updatedData struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if checkName(updatedData.Name) {
		http.Error(w, "/ are not allowed in Names", http.StatusBadRequest)
		return
	}

	fmt.Println(updatedData.Name)
	//object.Name = updatedData.Name

	//result = result.Update("name", updatedData.Name)
	result = db.Model(&Object{}).Where("id = ?", id).Update("name", updatedData.Name)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateParentFolderID(w http.ResponseWriter, r *http.Request) {
	var jsonObj Object
	id, _ := url.PathUnescape(r.PathValue("id"))

	result := db.First(&jsonObj, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "JSON object not found", http.StatusNotFound)
		return
	}

	var updatedData struct {
		ParentFolderID uint `json:"parentFolderID"`
	}

	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonObj.FolderID = updatedData.ParentFolderID

	result = db.Save(&jsonObj)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	var jsonObj Object
	id, _ := url.PathUnescape(r.PathValue("id"))

	result := db.First(&jsonObj, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "JSON object not found", http.StatusNotFound)
		return
	}

	result = db.Delete(&jsonObj)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if jsonObj.Type != "MD" {
		err := os.Remove(jsonObj.Data)
		if err != nil {
			fmt.Println("Cant find Objekt?")
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "object deleted")
}

func checkName(name string) bool {
	return strings.Contains(name, "/")
}
