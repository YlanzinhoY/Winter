package services

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExtractZip(nameZip, dstDir string) {

	archive, err := zip.OpenReader(nameZip)

	if err != nil {
		log.Fatal("Erro ao extrair o zip", err)
		return
	}

	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dstDir, f.Name)

		fmt.Println("unzipping file ", filePath)

		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
		os.Remove(nameZip)

	}

}
