package service

import (
	"io"
	"mime/multipart"
	"yssim-go/library/fileOperation"
)

func SaveBackground(path string, fileHeader *multipart.FileHeader) bool {
	file, _ := fileHeader.Open()
	data, _ := io.ReadAll(file)
	ok := fileOperation.WriteFileByte(path, data)
	return ok
}
