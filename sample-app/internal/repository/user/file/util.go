package file

import (
	"io"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
	uModel "github.com/regalius/clean-arch/sample-app/internal/model/user"
)

// fileToUsers convert given file to array of user struct
func fileToUsers(file io.Reader) (users []uModel.User, err error) {
	var arrByte []byte

	arrByte, err = ioutil.ReadAll(file)
	if err != nil {
		log.Println("[Repo/User/File] Failed to read file to array byte", err)
		return
	}
	err = jsoniter.ConfigFastest.Unmarshal(arrByte, &users)
	if err != nil {
		log.Println("[Repo/User/File] Failed to read file to array byte", err)
		return
	}
	return
}
