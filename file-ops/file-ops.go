package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 5000, errors.New("Failed To Finding File")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 5000, errors.New("Failed to parse store balanced value")
	}
	return value, nil
}

func WriteFloatIntoFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	//	fmt.Sprint(Balance)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
