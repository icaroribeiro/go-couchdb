package database

import (
	"../models"
)

func CreateDocument(doc models.Document) (error) {
	_, _, err := couchdb.Post(doc)

	return err
}

func GetDocuments() (models.AllDocuments, error) {
	var docs models.AllDocuments

	err := couchdb.AllDocs(&docs, nil)

	if err != nil {
    	return docs, err
    }

	return docs, nil
}

func GetDocumentById(id string) (models.Document, error) {
	var doc models.Document

	err := couchdb.Get(id, &doc, nil)

	if err != nil {
    	return doc, err
    }

	return doc, nil
}

func UpdateDocument(id string, doc models.Document) error {
	rev, err := couchdb.Rev(id)

	if err != nil {
    	return err
    }

	_, err = couchdb.Put(id, doc, rev)

	if err != nil {
    	return err
    }

    return nil
}

func DeleteDocument(id string) error {
	rev, err := couchdb.Rev(id)

	if err != nil {
    	return err
    }

	_, err = couchdb.Delete(id, rev)

	if err != nil {
    	return err
    }

    return nil
}
