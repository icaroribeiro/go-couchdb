package main

import (
  "log"
  "./controllers"
  "./database"
  "./models"
)

func init() {
    err := database.EstablishConnection()

  if err != nil {
    log.Fatal("Error while establishing database connection: ", err)
  }
}

func main() {
  //var id string
  //id = "5b269f93d4f487ed90ed93673203472d"

  var name string
  name = "Doc 15"

  var doc = models.Document{ 
      Name: name,
    }

  // Create Document
  controllers.CreateDocument(doc)

  // Get Documents
  //controllers.GetDocuments()

  // Get Document by Id
  //controllers.GetDocumentById(id)

  // Update Document
  //controllers.UpdateDocument(id, doc)

  // Delete Document
  //controllers.DeleteDocument(id)
}
