package file

import (
	"io"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
)

// fileToProducts convert given file to array of product struct
func fileToProducts(file io.Reader) (products []pModel.Product, err error) {
	var arrByte []byte

	arrByte, err = ioutil.ReadAll(file)
	if err != nil {
		log.Println("[Repo/Product/File] Failed to read file to array byte", err)
		return
	}
	err = jsoniter.ConfigFastest.Unmarshal(arrByte, &products)
	if err != nil {
		log.Println("[Repo/Product/File] Failed to read file to array byte", err)
		return
	}
	return
}
