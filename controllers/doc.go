package controllers

import (
	"log"
	"../database"
	"../models"
	"../utils"
)

func CreateDocument(doc models.Document) {
	err := database.CreateDocument(doc)

	if err != nil {
		log.Println("Error while creating the document:", err)
		return
	}

	response, err := utils.RespondWithJson(doc)

	if err != nil {
		log.Println("Error while creating the document:", err)
		return
	}

	log.Println(response)
	log.Println("Document created successfully")
}

func GetDocuments() {
	var docs models.AllDocuments

	docs, err := database.GetDocuments()

	if err != nil {
		log.Println("Error while getting the documents:", err)
		return
	}

	response, err := utils.RespondWithJson(docs)

	if err != nil {
		log.Println("Error while getting the document:", err)
		return
	}

	log.Println(response)
	log.Println("Documents returned successfully")
}

func GetDocumentById(id string) {
	var doc models.Document

	doc, err := database.GetDocumentById(id)

	if err != nil {
		log.Println("Error while getting the document:", err)
		return
	}

	response, err := utils.RespondWithJson(doc)

	if err != nil {
		log.Println("Error while getting the document:", err)
		return
	}

	log.Println(response)
	log.Println("Document returned successfully")
}

func UpdateDocument(id string, doc models.Document) {
	err := database.UpdateDocument(id, doc)

	if err != nil {
		log.Println("Error while updating the document:", err)
		return
	}

	response, err := utils.RespondWithJson(doc)

	if err != nil {
		log.Println("Error while updating the document:", err)
		return
	}

	log.Println(response)
	log.Println("Document updated successfully")
}

func DeleteDocument(id string) {
	err := database.DeleteDocument(id)

	if err != nil {
		log.Println("Error while deleting the document:", err)
		return
	}

	log.Println("Document deleted successfully")
}
