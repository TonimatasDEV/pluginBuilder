package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func CloseFile(outFile *os.File) {
	err := outFile.Close()
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}

func CloseZipReadCloser(r *zip.ReadCloser) {
	err := r.Close()
	if err != nil {
		log.Fatal("Error:" + err.Error())
	}
}

func CloseReadCloser(r io.ReadCloser) {
	err := r.Close()
	if err != nil {
		log.Fatal("Error:" + err.Error())
	}
}
