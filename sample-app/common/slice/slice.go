package slice

import (
	"errors"
	"strconv"
	"strings"
)

// SliceAtoi Convert given slice of string to slice of integer
func SliceAtoi(stringSlice []string) (intSlice []int, err error) {
	if stringSlice == nil {
		err = errors.New("[SliceAtoi] stringSlice is nil")
		return
	}
	intSlice = make([]int, 0, len(stringSlice))
	for _, str := range stringSlice {
		integer, errSlice := strconv.Atoi(str)
		if err != nil {
			return nil, errSlice
		}
		intSlice = append(intSlice, integer)
	}
	return intSlice, nil
}

// SliceItoa Convert given slice of integer to slice of string
func SliceItoa(intSlice []int) (stringSlice []string, err error) {
	if intSlice == nil {
		err = errors.New("[SliceItoa] intSlice is nil")
		return
	}
	stringSlice = make([]string, 0, len(intSlice))
	for _, integer := range intSlice {
		str := strconv.Itoa(integer)
		stringSlice = append(stringSlice, str)
	}
	return
}

// ConvertIntSliceToCommaDelimitedString Convert Slice of ints to comma delimited string
func ConvertIntSliceToCommaDelimitedString(intSlice []int) (str string, err error) {
	if intSlice == nil {
		err = errors.New("[ConvertIntSliceToCommaDelimitedString] intSlice is nil")
		return
	}
	var stringSlice []string
	stringSlice, err = SliceItoa(intSlice)
	if err != nil {
		return
	}
	str = strings.Join(stringSlice, ",")
	return
}

// ConvertStringSliceToCommaDelimitedString Convert Slice of strings to comma delimited string
func ConvertStringSliceToCommaDelimitedString(stringSlice []string) (str string, err error) {
	if stringSlice == nil {
		err = errors.New("[ConvertStringSliceToCommaDelimitedString] stringSlice is nil")
		return
	}
	str = strings.Join(stringSlice, ",")
	return
}
