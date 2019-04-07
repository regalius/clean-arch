package file

import (
	"io"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
)

// fileToProducts convert given file to array of product struct
func fileToUserPAffinity(file io.Reader) (userPAffinity []uPAModel.UserProductAffinity, err error) {
	var arrByte []byte

	arrByte, err = ioutil.ReadAll(file)
	if err != nil {
		log.Println("[Repo/UPAffinity/File] Failed to read file to array byte", err)
		return
	}
	err = jsoniter.ConfigFastest.Unmarshal(arrByte, &userPAffinity)
	if err != nil {
		log.Println("[Repo/UPAffinity/File] Failed to read file to array byte", err)
		return
	}
	return
}
