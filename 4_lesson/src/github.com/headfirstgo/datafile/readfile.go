// Пакет datafile предназначен для чтения данных из файлов.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки файла.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	//numbers := make([]float64, 10)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	number := 0.0
	for scanner.Scan() {
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		i++
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
