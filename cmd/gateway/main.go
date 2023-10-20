package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	endpoint := os.Getenv("HOST")
	port := os.Getenv("PORT")
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := tasksv1.RegisterTasksServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalln("Failed to register tasks service:", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Handle("/api", mux)
	log.Println("Listening on port:", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalln("Failed to listen and serve:", err)
	}
}
