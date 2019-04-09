package product

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

// GetFullImagePath get full image path from image object
func (i Image) GetFullImagePath() string {
	return fmt.Sprintf("%s/%s", i.Path, i.Name)
}

// GetURL Get URL of the product given baseURL to be appended
func (p Product) GetURL(baseURL string) (result string) {
	fullURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println("[Model/Product] Error parsing baseURL: ", err.Error())
		return
	}
	fullURL.Path += fmt.Sprintf("/%v/%v/%s", p.ShopID, p.ID, strings.Replace(p.Name, " ", "-", -1))
	result = fullURL.String()
	return
}
