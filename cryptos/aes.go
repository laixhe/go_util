package cryptos

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// 对称加密算法
// AES 加密数据块分组长度必须为 128bit(16bytes)
// 密钥长度可以是 128bit(16bytes)、192bit(24bytes)、256bit(32bytes) 中的任意一个
// 秘钥长度需要时 AES-128(16bytes) 或者 AES-256(32bytes)

// AesEncrypt 使用 AES 加密文本,加密的文本不能为空
func AesEncrypt(data []byte, key []byte) ([]byte, error) {

	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	data = PKCS7Padding(data, blockSize)
	// 加密模式 CBC
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	// 加密
	cryted := make([]byte, len(data))
	blockMode.CryptBlocks(cryted, data)

	return cryted, nil

}

// AesDecrypt 使用 AES 解密文本
func AesDecrypt(data []byte, key []byte) ([]byte, error) {

	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式 CBC
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	// 解密
	orig := make([]byte, len(data))
	blockMode.CryptBlocks(orig, data)

	// 去补全码
	return PKCS7UnPadding(orig), nil
}

// PKCS7Padding 使用 AES 加密文本的时候文本必须定长
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7UnPadding 使用 AES 解密文本,解密收删除 padding 的文本
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
