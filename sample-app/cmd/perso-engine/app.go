package main

import (
	"context"
	"log"
	"os"

	pFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product/file"
)

func main() {
	productFile, err := os.Open("./configs/var/data/product.json")
	if err != nil {
		log.Println("[PersoEngine] Failed to open product.json", err)
	}

	productFileRepository := pFileRepo.NewFileProductRepository(productFile)
	ctx := context.Background()
	product, err := productFileRepository.GetByIDRange(ctx, 3, 5)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(product)
}
