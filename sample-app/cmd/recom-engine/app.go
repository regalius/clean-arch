package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	config "github.com/regalius/clean-arch/sample-app/common/config"
	"google.golang.org/grpc"

	gPAFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity/file"
	pFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product/file"
	uPAFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity/file"
	uFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user/file"

	pRecomGenderUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom/gender-affinity"
	pRecomMultiUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom/multi-affinity"
	pRecomUserUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom/user-affinity"

	rEngineGRPCDelivery "github.com/regalius/clean-arch/sample-app/internal/delivery/recom-engine/grpc"
	rEngineHTTPDelivery "github.com/regalius/clean-arch/sample-app/internal/delivery/recom-engine/http"
)

func main() {
	openFiles := make(map[string]*os.File)
	var err error
	for context, path := range fileList {
		openFiles[context], err = os.Open(path)
		if err != nil {
			log.Fatal("[RecomEngine] Failed to open", path, err)
		}
	}

	ENV := os.Getenv("GO_ENV")
	if ENV == "" {
		ENV = "production"
		os.Setenv("GO_ENV", ENV)
	}

	err = config.NewConfigFromFile("recomengine", "toml", configPaths[ENV], config.NewConfigOptions{
		IsWatch: true,
	})
	if err != nil {
		log.Fatal("[RecomEngine] Error reading config from file", err)
	}

	pRepo := pFileRepo.NewFileProductRepository(openFiles["product"])
	uRepo := uFileRepo.NewFileUserRepository(openFiles["user"])
	uPARepo := uPAFileRepo.NewFileUserProductAffinityRepository(openFiles["userPA"])
	gPARepo := gPAFileRepo.NewFileGenderProductAffinityRepository(openFiles["genderPA"])

	pRGenderUcase := pRecomGenderUcase.NewProductRecomGenderUsecase(uRepo, pRepo, gPARepo)
	pRUserUcase := pRecomUserUcase.NewProductRecomUserUsecase(uRepo, pRepo, uPARepo)
	pRMultiUcase := pRecomMultiUcase.NewProductRecomMultiUsecase(uRepo, pRepo, uPARepo, gPARepo)

	router := httprouter.New()
	rEngineHTTPDelivery.NewHttpRecomEngineDelivery(router, pRGenderUcase, pRUserUcase, pRMultiUcase)

	grpcServer := grpc.NewServer()
	rEngineGRPCDelivery.NewGrpcRecomEngineDelivery(grpcServer, configName, pRGenderUcase, pRUserUcase, pRMultiUcase)

	grpcPort, err := config.Get(configName, "Delivery.GRPC.port")
	if err != nil {
		log.Fatal("[RecomEngine] Couldn't get grpc port from config", err)
	}

	httpPort, err := config.Get(configName, "Delivery.HTTP.port")
	if err != nil {
		log.Fatal("[RecomEngine] Couldn't get http port from config", err)
	}

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("[RecomEngine] Couldn't create net listener in port %s %v", grpcPort, err)
		return
	}

	go func() {
		log.Printf("[RecomEngine] GRPC Handler listen on :%s\n", grpcPort)
		log.Fatal(grpcServer.Serve(listener))
	}()

	log.Printf("[RecomEngine] HTTP Handler listen on :%s\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), router))

}
