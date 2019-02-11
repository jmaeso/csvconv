package main

import (
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		log.Fatal("Need xlsx file as first arg")
	}

	baseInputPath, baseInputName := getFileInfo(args[1])

	xlFile, err := xlsx.OpenFile(args[1])
	if err != nil {
		log.Fatal("ERROR: Could not open xlsx: ", err)
	}

	for i, sheet := range xlFile.Sheets {
		f, err := os.Create(baseInputPath + baseInputName + "-" + strconv.Itoa(i) + ".csv")
		if err != nil {
			log.Fatal("ERROR: Could not create csv: ", err)
		}

		for _, row := range sheet.Rows {
			cells := make([]string, 0)

			for _, cell := range row.Cells {
				cellText := cell.String()
				cellText = strings.Replace(cellText, ",", ".", -1)

				cells = append(cells, cellText)
			}

			rowText := strings.Join(cells, ",")

			_, err := f.WriteString(rowText + "\n")
			if err != nil {
				log.Fatal("ERROR: Could not write to csv: ", err)
			}
		}

		if err := f.Close(); err != nil {
			log.Fatal("ERROR: Could not close csv: ", err)
		}
	}
}

func getFileInfo(filePath string) (string, string) {
	basePath, fileName := path.Split(filePath)

	dotPos := strings.IndexRune(fileName, '.')
	if dotPos < 0 {
		dotPos = len(fileName)
	}
	baseName := string([]byte(fileName)[:dotPos])

	return basePath, baseName
}
