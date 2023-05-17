package main

import (
	"context"
	_ "embed"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/api/option"

	"google.golang.org/api/sheets/v4"
)

//go:embed credential.json
var credential []byte

var (
	spreadSheetID string
	sheetName     = "test"
)

func main() {
	flag.StringVar(&spreadSheetID, "s", "", "sheetID")
	flag.Parse()

	credential := option.WithCredentialsJSON(credential)
	srv, err := sheets.NewService(context.Background(), credential)
	if err != nil {
		panic(err)
	}
	if err := execute(srv); err != nil {
		panic(err)
	}
}

func execute(srv *sheets.Service) error {
	resp, err := srv.Spreadsheets.Values.Get(spreadSheetID, sheetName+"!A1:C100").Do()
	if err != nil {
		return fmt.Errorf("failed to get values: %w", err)
	}
	filePath := filepath.Join("dst", "test.csv")
	destDir := filepath.Dir(filePath)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	err = csv.NewWriter(file).WriteAll(convertToSliceOfString(resp.Values))
	if err != nil {
		return fmt.Errorf("failed to write csv: %w", err)
	}
	return nil
}

// convertToSliceOfString converts [][]interface{} to [][]string
func convertToSliceOfString(input [][]interface{}) [][]string {
	result := make([][]string, len(input))
	for i, row := range input {
		result[i] = make([]string, len(row))
		for j, val := range row {
			result[i][j] = val.(string)
		}
	}
	return result
}
