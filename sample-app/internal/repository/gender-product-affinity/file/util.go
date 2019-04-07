package file

import (
	"io"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
)

// fileToProducts convert given file to array of product struct
func fileToGenderPAffinity(file io.Reader) (genderPAffinity []gPAModel.GenderProductAffinity, err error) {
	var arrByte []byte

	arrByte, err = ioutil.ReadAll(file)
	if err != nil {
		log.Println("[Repo/GPAffinity/File] Failed to read file to array byte", err)
		return
	}
	err = jsoniter.ConfigFastest.Unmarshal(arrByte, &genderPAffinity)
	if err != nil {
		log.Println("[Repo/GPAffinity/File] Failed to read file to array byte", err)
		return
	}
	return
}
