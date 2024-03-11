package main

import (
	"context"
	"haproxyctl/pkg/dataplane-client"
	"log"
)

const (
	username = "admin"
	password = "adminpwd"
)

var dataplaneClient *dataplane.APIClient

var auth context.Context

func main() {
	setup()

	log.Printf("%+v\n", getBackends())
}

func setup() {
	cfg := dataplane.NewConfiguration()
	cfg.BasePath = "http://localhost:15555/v2"
	dataplaneClient = dataplane.NewAPIClient(cfg)

	auth = context.WithValue(context.Background(), dataplane.ContextBasicAuth, dataplane.BasicAuth{
		UserName: username,
		Password: password,
	})
}

func getBackends() []dataplane.Backend {
	results, resp, err := dataplaneClient.BackendApi.GetBackends(auth, nil)
	log.Printf("http response: %v", resp.StatusCode)
	if err != nil {
		log.Printf("error: %v", err)
		return nil
	}

	return results.Data
}
