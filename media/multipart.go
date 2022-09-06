package media

import (
	"io"
	"mime/multipart"
	"strings"
)

func GetFileReaderAndType(resource *multipart.FileHeader) (io.Reader, string, error) {
	reader, err := resource.Open()
	if err != nil {
		return nil, "", err
	}

	fileType := resource.Header.Get("Content-Type")
	arr := strings.Split(fileType, "/")
	if len(arr) < 2 {
		return nil, "", err
	}

	if strings.Contains(arr[1], "svg") {
		arr[1] = "svg"
	}

	return reader, arr[1], nil
}
