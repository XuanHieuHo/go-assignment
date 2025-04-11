package importers

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/XuanHieuHo/go-assignment/utils"
)

func ImportCSV[M CSVImporter](ctx context.Context, file io.Reader, model M) error {
	reader := csv.NewReader(file)
	headers, err := utils.ReadCSVFile(reader)
	if err != nil {
		return err
	}
	headerCheck := make(map[string]int)
	for i, header := range headers {
		headerCheck[strings.ToLower(strings.TrimSpace(header))] = i
	}

	var importErr error
	for {
		record, err := utils.ReadCSVFile(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if err := model.ParseFromRow(record, headerCheck); err != nil {
			importErr = errors.Join(importErr, err)
			continue
		}

		if err := model.Save(ctx); err != nil {
			importErr = errors.Join(importErr, err)
			continue
		}
	}

	if importErr != nil {
		log.Printf("IMPORT ERROR: %+v", importErr)
	}
	return nil
}
