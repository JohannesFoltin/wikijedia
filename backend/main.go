package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error

	db, err = gorm.Open(sqlite.Open("./Data/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Folder{}, &Object{})

	var zeroFolder = Folder{
		Name:     "Workspace",
		ParentID: 0,
		Children: []Folder{},
		Objects:  []Object{},
	}

	var count int64
	db.Model(&Folder{}).Where("id = ?", 0).Count(&count)
	if count == 0 {
		db.Create(&zeroFolder)
	}

	// Create a new folder
	http.Handle("POST /folder", enableCORS(http.HandlerFunc(createFolder)))
	http.Handle("OPTIONS /folder", http.HandlerFunc(handleCorsRequest))

	// Get a folder by ID
	http.Handle("GET /folder/{id}", enableCORS(http.HandlerFunc(getFolder)))

	// Update a folder by ID
	http.Handle("PUT /folder/{id}", enableCORS(http.HandlerFunc(updateFolder)))
	http.Handle("OPTIONS /folder/{id}", http.HandlerFunc(handleCorsRequest))

	http.Handle("PUT /folder/{id}/parent", enableCORS(http.HandlerFunc(updateFolderParentFolderID)))
	http.Handle("OPTIONS /folder/{id}/parent", http.HandlerFunc(handleCorsRequest))

	// Delete a folder by ID
	http.Handle("DELETE /folder/{id}", enableCORS(http.HandlerFunc(deleteFolder)))

	// Create a new JSON object
	http.Handle("POST /object", enableCORS(http.HandlerFunc(createObject)))
	http.Handle("OPTIONS /object", http.HandlerFunc(handleCorsRequest))

	// Get a JSON object by ID
	http.Handle("GET /object/{id}", enableCORS(http.HandlerFunc(getObject)))
	http.Handle("GET /object/data/{id}", enableCORS(http.HandlerFunc(getFileFromObject)))

	// Update a JSON object by ID
	http.Handle("PUT /object/{id}", enableCORS(http.HandlerFunc(updateObject)))
	http.Handle("OPTIONS /object/{id}", http.HandlerFunc(handleCorsRequest))

	// Update a JSON object's name by ID
	http.Handle("PUT /object/{id}/name", enableCORS(http.HandlerFunc(updateObjectName)))
	http.Handle("OPTIONS /object/{id}/name", http.HandlerFunc(handleCorsRequest))

	http.Handle("PUT /object/{id}/parent", enableCORS(http.HandlerFunc(updateParentFolderID)))
	http.Handle("OPTIONS /object/{id}/parent", http.HandlerFunc(handleCorsRequest))

	// Delete a JSON object by ID
	http.Handle("DELETE /object/{id}", enableCORS(http.HandlerFunc(deleteObject)))

	// Get the folder structure
	http.Handle("GET /structure", enableCORS(http.HandlerFunc(getStructure)))

	// File upload
	http.Handle("POST /upload", enableCORS(http.HandlerFunc(handleFileUpload)))
	http.Handle("OPTIONS /upload", http.HandlerFunc(handleCorsRequest))

	// Start the server
	fmt.Println("Server started")
	errh := http.ListenAndServe(":8080", nil)
	if errh != nil {
		fmt.Println("http Server error")
		return
	}

}

func getStructure(w http.ResponseWriter, r *http.Request) {
	var folders []Folder

	// Retrieve all folders from the database
	result := db.Preload("Objects").Find(&folders)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Filter out the "Data" column from all objects
	for i := range folders {
		for j := range folders[i].Objects {
			folders[i].Objects[j].Data = ""
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}
