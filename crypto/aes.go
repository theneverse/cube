package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

var aesKey = "abcdef0123456789"
var aesIV = "0123456789abcdef"

func AESEncrypt(data []byte) ([]byte, error) {
	aesBlockEncrypter, err := aes.NewCipher([]byte(aesKey))
	content := PKCS5Padding(data, aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, []byte(aesIV))
	aesEncrypter.CryptBlocks(encrypted, content)
	return encrypted, nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
