package gzipx

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Compress 压缩文本
func Compress(text string) (compressText string, err error) {
	if text == "" {
		return
	}

	// 中文转码
	var chineseBuf bytes.Buffer
	chineseWrite := transform.NewWriter(&chineseBuf, simplifiedchinese.GB18030.NewEncoder())
	_, err = chineseWrite.Write([]byte(text))
	if err != nil {
		return
	}
	byteText := chineseBuf.Bytes()

	// 压缩
	var gzipBuf bytes.Buffer
	zw := gzip.NewWriter(&gzipBuf)
	defer zw.Close()
	if _, err = zw.Write(byteText); err != nil {
		return
	}
	if err = zw.Flush(); err != nil {
		return
	}
	gzipBytes := gzipBuf.Bytes()

	compressText = base64.StdEncoding.EncodeToString(gzipBytes)
	return
}

// Decompress 解压缩
func Decompress(text string) (string, error) {
	if text == "" {
		return "", nil
	}
	// base64解码
	z, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	reader, err := gzip.NewReader(bytes.NewReader(z))
	if err != nil {
		return "", err
	}
	defer func(reader *gzip.Reader) {
		err := reader.Close()
		if err != nil {
			if err != io.ErrUnexpectedEOF {
				return
			}
		}
	}(reader)

	te, _ := ioutil.ReadAll(reader)
	// 中文转码
	decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(te)
	return string(decodeBytes), nil
}
