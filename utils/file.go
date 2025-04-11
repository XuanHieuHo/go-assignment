package utils

import (
	"encoding/csv"
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func ReadCSVFile(reader *csv.Reader) ([]string, error) {
	data, err := reader.Read()
	return data, err
}

func WriteCSVFile(filePath string, data [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
