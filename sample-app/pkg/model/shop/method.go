package shop

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

// getFullImagePath : get full image path from image object
func (i Image) getFullImagePath() string {
	return fmt.Sprintf("%s/%s", i.Path, i.Name)
}

// getgetPreviousNamesAsString: Get shop's previous names in comma delimited string
func (s Shop) getPreviousNamesAsString() (result string) {
	if len(s.PreviousNames) > 0 {
		result = strings.Join(s.PreviousNames, ",")
	}
	return
}

// getURL: Get URL of the shops given baseURL to be appended
func (s Shop) getURL(baseURL string) (result string) {
	fullURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println("[Model/Shop] Error parsing baseURL: ", err.Error())
		return
	}
	fullURL.Path += fmt.Sprintf("/%v/%s", s.ID, strings.Replace(s.Name, " ", "-", -1))
	result = fullURL.String()
	return
}

// getShopTypeString: Get Shop type in human readable string
func (s Shop) getShopTypeString() (result string) {
	return shopTypesEnum[s.Type]
}

// getReputationBadgeString: Get Reputation in human readable string
func (s Shop) getReputationString() (result string) {
	for _, enum := range shopReputationEnum {
		if s.ReputationScore >= enum.Gt &&
			(enum.Lt == 0 || s.ReputationScore <= enum.Lt) {
			result = enum.Name
			break
		}
	}
	return
}
