package utils

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url, filepath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer CloseReadCloser(response.Body)

	outFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer CloseFile(outFile)

	_, err = io.Copy(outFile, response.Body)
	return err
}

func unzipFile(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer CloseZipReadCloser(r)

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)

		rc, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}

		CloseReadCloser(rc)
		CloseFile(outFile)
	}

	return nil
}
