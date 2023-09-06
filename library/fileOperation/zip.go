package fileOperation

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver/v3"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func UnZip(filePath string, toPath string) error {
	err := errors.New("")
	if strings.HasSuffix(filePath, ".zip") {
		err = unZip(filePath, toPath)
	} else {
		err = archiver.Unarchive(filePath, toPath)
	}
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func unZip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	var decodeName string
	for _, f := range zipReader.File {
		if f.Flags == 0 || f.Flags == 8 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GBK.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = f.Name
		}

		fpath := filepath.Join(destDir, decodeName)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Zip(filePath string, toPath string) error {
	filePathList := []string{}
	file, _ := os.ReadDir(filePath)
	for _, name := range file {
		filePathList = append(filePathList, filePath+"/"+name.Name())
	}
	err := archiver.Archive(filePathList, toPath)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
