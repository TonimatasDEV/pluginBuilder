package utils

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFile(url, filepath string) error {
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

func UnzipFile(zipPath, dest string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}

	firstFile := true
	var mainDir string
	for _, file := range r.File {
		if firstFile {
			mainDir = file.Name
			firstFile = false
			continue
		}

		fpath := filepath.Join(dest, strings.ReplaceAll(file.Name, mainDir, ""))
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {

			Fatal("illegal file path: " + fpath)
		}

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, os.ModePerm)

			if err != nil {
				Fatal(err.Error())
			}

			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			Fatal(err.Error())
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())

		if err != nil {
			Fatal(err.Error())
		}

		rc, err := file.Open()

		if err != nil {
			Fatal(err.Error())
		}

		_, err = io.Copy(outFile, rc)

		if err != nil {
			Fatal(err.Error())
		}

		CloseFile(outFile)
		CloseReadCloser(rc)
	}

	CloseZipReadCloser(r)

	return deleteFile(zipPath)
}

func deleteFile(path string) error {
	return os.Remove(path)
}
