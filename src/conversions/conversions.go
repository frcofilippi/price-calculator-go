package conversions

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	result := make([]float64, len(strings))
	for index, stringValue := range strings {
		price, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			errorMessage := fmt.Sprintf("error converting string to float. %v", stringValue)
			return nil, errors.New(errorMessage)
		}

		result[index] = price
	}
	return result, nil
}
