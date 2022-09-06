package helper

import (
	"crypto/aes"
	"crypto/cipher"
	rand2 "crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/theneverse/go-hammer/encoding/base24"
	"io"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	// ENCRYPT
	c, err := rc4.NewCipher([]byte("dsadsad"))
	if err != nil {
		log.Fatalln(err)
	}
	src := []byte("asdsad")
	fmt.Println("Plaintext: ", src)

	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	fmt.Println("Ciphertext: ", dst)

	// DECRYPT
	c2, err := rc4.NewCipher([]byte("dsadsad"))
	if err != nil {
		log.Fatalln(err)
	}
	src2 := make([]byte, len(dst))
	c2.XORKeyStream(src2, dst)
	fmt.Println("Plaintext': ", src2)
}

func TestGenerateSMSCode(t *testing.T) {
	rand.Seed(time.Now().Unix())
	assetId := uint16(3005)
	orderNo := uint16(16)
	randUint32 := rand.Uint32()
	// 15: 9CYFGT5S2RPKHK
	// 16: 9CYFGT9AW8CHH9

	assetIdBytes := Uint16ToBytes(assetId)
	orderNoBytes := Uint16ToBytes(orderNo)
	randBytes := Uint32ToBytes(randUint32)

	var result []byte
	result = append(result, assetIdBytes[:]...)
	result = append(result, orderNoBytes[:]...)
	result = append(result, randBytes[:]...)
	fmt.Println("Plaintext: ", result)

	code, err := encrypt([]byte("neverse001122333"), hex.EncodeToString(result))
	if err != nil {
		panic(err)
	}

	fmt.Println("Ciphertext: ", code)

	src, err := decrypt([]byte("neverse001122333"), code)
	if err != nil {
		panic(err)
	}

	fmt.Println("Plaintext': ", src)
}

/*
 *	FUNCTION		: encrypt
 *	DESCRIPTION		:
 *		This function takes a string and a cipher key and uses AES to encrypt the message
 *
 *	PARAMETERS		:
 *		byte[] key	: Byte array containing the cipher key
 *		string message	: String containing the message to encrypt
 *
 *	RETURNS			:
 *		string encoded	: String containing the encoded user input
 *		error err	: Error message
 */
func encrypt(key []byte, message string) (encoded string, err error) {
	//Create byte array from the input string
	plainText := []byte(message)

	//Create a new AES cipher using the key
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand2.Reader, iv); err != nil {
		return
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//Return string encoded in base64
	return base24.StdEncoding.EncodeToString(cipherText)
}

/*
 *	FUNCTION		: decrypt
 *	DESCRIPTION		:
 *		This function takes a string and a key and uses AES to decrypt the string into plain text
 *
 *	PARAMETERS		:
 *		byte[] key	: Byte array containing the cipher key
 *		string secure	: String containing an encrypted message
 *
 *	RETURNS			:
 *		string decoded	: String containing the decrypted equivalent of secure
 *		error err	: Error message
 */
func decrypt(key []byte, secure string) (decoded string, err error) {
	//Remove base64 encoding:
	cipherText, err := base24.StdEncoding.DecodeString(secure)

	//IF DecodeString failed, exit:
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
