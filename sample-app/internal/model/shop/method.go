package shop

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

// GetPreviousNamesAsString Get shop's previous names in comma delimited string
func (s Shop) GetPreviousNamesAsString() (result string) {
	if len(s.PreviousNames) > 0 {
		result = strings.Join(s.PreviousNames, ",")
	}
	return
}

// GetURL Get URL of the shops given baseURL to be appended
func (s Shop) GetURL(baseURL string) (result string) {
	fullURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println("[Model/Shop] Error parsing baseURL: ", err.Error())
		return
	}
	fullURL.Path += fmt.Sprintf("/%v/%s", s.ID, strings.Replace(s.Name, " ", "-", -1))
	result = fullURL.String()
	return
}

// GetShopTypeString Get Shop type in human readable string
func (s Shop) GetShopTypeString() (result string) {
	return shopTypesEnum[s.Type]
}

// GetReputationString Get Reputation in human readable string
func (s Shop) GetReputationString() (result string) {
	for _, enum := range shopReputationEnum {
		if s.ReputationScore >= enum.GreaterThan &&
			(enum.LessThan == 0 || s.ReputationScore <= enum.LessThan) {
			result = enum.Name
			break
		}
	}
	return
}
