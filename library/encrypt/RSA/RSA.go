package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func EncyptogRSA(src []byte, key []byte) ([]byte, error) {

	// 将得到的字符串解码
	block, _ := pem.Decode(key)

	// 使用X509将解码之后的数据 解析出来
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

	//4.使用公钥加密数据
	res, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	return res, err
}

// 对数据进行解密操作
func DecrptogRSA(src []byte, key []byte) (res []byte, err error) {

	block, _ := pem.Decode(key)                               //解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes) //还原数据
	res, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	return
}
