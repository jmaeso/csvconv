# csvconv

## Usage instructions

### Windows

Drag a `xlsx` file and drop it to the `exe` file.

### MacOS - Linux

`csvconv ./path/file.xslx`

## Details

- The resultant `csv` file/s will be created in the same directory than the `xslx` input file.

- Any coma (`,`) in a cell will be replaced with a dot (`.`).

- If there are multiple sheets in the excel file, multiple `csv` files will be generated with a trailing incremental number.

## Compile for windows

- `GOOS=windows GOARCH=386 go build -o csvconv.exe main.go`.
