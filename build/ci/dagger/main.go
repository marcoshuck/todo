package main

import (
	"context"
	"dagger.io/dagger"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	pipeline = map[string][]Job{
		"format": {
			{
				Image:   "golang:1.21",
				Name:    "fmt",
				Command: "go fmt ./...",
			},
			{
				Image:   "golang:1.21",
				Name:    "vet",
				Command: "go vet -v ./...",
			},
			{
				Image:   "golangci/golangci-lint:v1.54.1",
				Name:    "lint",
				Command: "golangci-lint run -v",
			},
		},
		"test": {
			{
				Image:   "golang:1.21",
				Name:    "race",
				Command: "go test -race ./...",
			},
			{
				Image:   "golang:1.21",
				Name:    "test",
				Command: "go test -covermode=atomic -coverprofile=coverage.tx -v ./...",
			},
		},
	}
)

type Job struct {
	Name    string
	Command string
	Image   string
}

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		log.Fatalln("Error while running Dagger:", err)
	}
	defer client.Close()

	src := getSourceDirectory(client)

	for stageName, jobs := range pipeline {
		for _, job := range jobs {
			ci := client.Container().Pipeline(fmt.Sprintf("%s/%s", stageName, job.Name)).
				From(job.Image).
				WithDirectory("/src", src).
				WithWorkdir("/src").
				WithExec(strings.Split(job.Command, " "))

			_, err := ci.Stderr(ctx)
			if err != nil {
				log.Printf("Job %s/%s failed with error: %v\n", stageName, job.Name, err)
				continue
			}
			log.Printf("Job %s/%s succeeded\n", stageName, job.Name)
		}
	}

	// Pipeline: Format, Test, Build
	// Format:
	//	- golangci-lint
	//	- go fmt
	//  - go vet ./...
	// Test with MySQL
	//	- Race conditions
	//	- Unit tests
	//	- Integration tests
}

func getSourceDirectory(client *dagger.Client) *dagger.Directory {
	return client.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"build/ci", "deployments", ".gitignore", "buf.*"},
	})
}
