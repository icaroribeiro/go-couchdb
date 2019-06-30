package database

import (
    "fmt"
    "net/url"
    cdb "github.com/cabify/go-couchdb"
    "../utils"
)

var client *cdb.Client
var couchdb *cdb.DB

const (
    path = "./database/cfg.json"
)

type Configuration struct {
    Ip   string
    Port string
    Database string
}

func EstablishConnection() (error) {
    cfg := Configuration{}

    err := utils.ReadJsonObjects(path, &cfg)

    if err != nil {
        return err
    }

    dsn := fmt.Sprintf("http://%s:%s/", cfg.Ip, cfg.Port)

    u, err := url.Parse(dsn)

    if err != nil {
        return err
    }

    client := cdb.NewClient(u, nil, nil)

    if client == nil {
        return fmt.Errorf("a remote CouchDB server creation failed")
    }

    // Check whether the couchdb already exists or not.
    // If so, the db instance will be returned. Otherwise, the db will be created.
    couchdb, err = client.EnsureDB(cfg.Database)

    if err != nil {
        return err
    }

    return nil
}
