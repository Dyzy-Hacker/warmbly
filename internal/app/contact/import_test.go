package contact

import (
	"bytes"
	"strings"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestParseSpreadsheetCSV(t *testing.T) {
	t.Parallel()

	input := "Email,First Name,Company\ntest@example.com,John,Test Company\nhello@example.com,Sarah,Demo Company\n"
	rows, format, xerr := parseSpreadsheet(strings.NewReader(input), "contacts.csv")
	if xerr != nil {
		t.Fatalf("parse CSV: %v", xerr)
	}
	if format != "csv" || len(rows) != 3 || rows[1][0] != "test@example.com" {
		t.Fatalf("unexpected CSV result: format=%q rows=%v", format, rows)
	}
}

func TestParseSpreadsheetXLSX(t *testing.T) {
	t.Parallel()

	f := excelize.NewFile()
	defer f.Close()
	sheet := f.GetSheetName(f.GetActiveSheetIndex())
	values := [][]any{
		{"Email", "First Name", "Company"},
		{"test@example.com", "John", "Test Company"},
		{"hello@example.com", "Sarah", "Demo Company"},
	}
	for i, row := range values {
		cell, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			t.Fatal(err)
		}
		if err := f.SetSheetRow(sheet, cell, &row); err != nil {
			t.Fatal(err)
		}
	}
	var workbook bytes.Buffer
	if err := f.Write(&workbook); err != nil {
		t.Fatalf("write XLSX: %v", err)
	}

	rows, format, xerr := parseSpreadsheet(bytes.NewReader(workbook.Bytes()), "contacts.xlsx")
	if xerr != nil {
		t.Fatalf("parse XLSX: %v", xerr)
	}
	if format != "xlsx" || len(rows) != 3 || rows[2][0] != "hello@example.com" {
		t.Fatalf("unexpected XLSX result: format=%q rows=%v", format, rows)
	}
}
