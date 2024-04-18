package service

import (
	_ "embed"
	"errors"
	"io"
	"yssim-go/config"
	"yssim-go/library/encrypt/AES"
	"yssim-go/library/encrypt/RSA"
	"yssim-go/library/fileOperation"

	"log"
)

func ModelEncrypt(data []byte) []byte {
	random := AES.GetRandom()
	// 对data数据进行AES加密
	AesEnc, err := AES.EncryptByAes(data, random)
	if err != nil {
		log.Println(err)
	}
	// 对随机字符进行RSA加密
	randomByteRsa, _ := RSA.EncyptogRSA(random, config.PublicKey)
	// 将加密模型与加密随机字符进行组合
	res := append(randomByteRsa, AesEnc...)
	return res
}

func ModelDecrypt(data []byte) ([]byte, error) {
	// 对随机字符进行RSA解密
	randomStr, err := RSA.DecrptogRSA(data[:512], config.PrivateKey)

	// 对加密模型进行AES解密
	aesDecrypt, err := AES.AesDecrypt(data[512:], randomStr)
	if err != nil {
		return nil, errors.New("模型解密失败")
	}
	return aesDecrypt, nil
}

func FileDecrypt(file io.Reader, filePath string) bool {
	fileData, _ := io.ReadAll(file)
	model, err := ModelDecrypt(fileData)
	if err != nil {
		return false
	}
	fileOperation.WriteFile(filePath, string(model))
	return true
}
