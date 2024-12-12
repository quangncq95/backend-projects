#!/bin/bash

OUTPUT_DIR="../bin"
OUTPUT_FILE="main.exe"

# Create the output directory if it doesn't exist
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR/web
cp -r ../web $OUTPUT_DIR

# Build the executable
go build -o $OUTPUT_DIR/cmd/$OUTPUT_FILE ../cmd/main.go
echo "Executable created at $OUTPUT_DIR/cmd/$OUTPUT_FILE"
cd $OUTPUT_DIR/cmd
./$OUTPUT_FILE
echo "Run executable file"
