package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product/internal/config"
	"product/internal/router"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	cfg := config.Load()
     
    ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
    if err!= nil{
		log.Fatalf("Failed to create mongodb client: %v",err)
	}
	err = client.Ping(ctx, nil)
	if err != nil{
		log.Fatalf("Failed to connect to mongodb %v", err)
	}
	defer client.Disconnect(ctx)

    r := router.SetupRoutes(cfg, client)

    fmt.Println("Startingserver on port", cfg.Port)
    err = http.ListenAndServe(":"+cfg.Port, r)
    if err!= nil{
     	log.Fatalf("server failed %v", err)
     }
}