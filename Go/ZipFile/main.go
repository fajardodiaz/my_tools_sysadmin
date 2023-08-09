package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func zipFolder(sourceFolder, zipFilePath string) error {
	// create a new zip file
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source folder and add files to the zip archive
	err = filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Create a new entry for the file in the zip archive
		relPath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}
		zipEntry, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// Open the source file
		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		// Copy the file's content to the zip entry
		_, err = io.Copy(zipEntry, sourceFile)
		return err
	})

	return err

}

func main() {
	args := os.Args

	sourceFolder := args[1]
	zipFilePath := args[2]

	err := zipFolder(sourceFolder, zipFilePath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Folder zipped successfully!")
	}

}
