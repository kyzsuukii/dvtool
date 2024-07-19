package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func Base64Decode(input string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(input)
	return string(bytes), err
}

func XOREncrypt(input, key []byte) []byte {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return output
}

func EncryptString(input string, key []byte) string {
	base64Encoded := Base64Encode(input)
	encrypted := XOREncrypt([]byte(base64Encoded), key)
	return Base64Encode(string(encrypted))
}

func DecryptString(input string, key []byte) (string, error) {
	base64Decoded, err := Base64Decode(input)
	if err != nil {
		return "", err
	}
	decrypted := XOREncrypt([]byte(base64Decoded), key)
	original, err := Base64Decode(string(decrypted))
	if err != nil {
		return "", err
	}
	return original, nil
}

func GenerateRandomKey(length int) ([]byte, error) {
	if length <= 0 {
		return nil, errors.New("invalid key length")
	}
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
