package cf

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Zipper interface {
	Zip(dirToZip string) (zip *bytes.Buffer, err error)
}

type ApplicationZipper struct{}

func (zipper ApplicationZipper) Zip(dirToZip string) (zipBuffer *bytes.Buffer, err error) {
	zipBuffer = new(bytes.Buffer)
	writer := zip.NewWriter(zipBuffer)

	addFileToZip := func(path string, f os.FileInfo, inErr error) (err error) {
		err = inErr
		if err != nil {
			return
		}

		if f.IsDir() {
			return
		}

		fileName := strings.TrimPrefix(path, dirToZip+"/")
		zipFile, err := writer.Create(fileName)
		if err != nil {
			return
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return
		}

		_, err = zipFile.Write(content)
		if err != nil {
			return
		}

		return
	}

	err = filepath.Walk(dirToZip, addFileToZip)

	if err != nil {
		return
	}

	err = writer.Close()
	return
}