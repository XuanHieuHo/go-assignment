package importers

import (
	"context"
)


type CSVImporter interface {
	ParseFromRow(row []string, headerCheck map[string]int) error
	Save(ctx context.Context) error
}
