package sign

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

// RsaSign 非对称加密
func RsaSign(secretKey, body string) []byte {
	ret, err := PublicEncrypt(body, secretKey)
	if err != nil {
		panic(err)
	}
	return []byte(ret)
}

func RsaDecryptSign(decryptStr, path string) []byte {
	ret, err := PrivateDecrypt(decryptStr, path)
	if err != nil {
		panic(err)
	}
	return []byte(ret)
}

// PublicEncrypt 公钥加密
func PublicEncrypt(encryptStr string, path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	// 读取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, _ = file.Read(buf)

	// pem 解码
	block, _ := pem.Decode(buf)

	// x509 解码
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	//publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 类型断言
	//publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	//if !ok {
	//	return "", errors.New("rsa公钥解析失败")
	//}

	//对明文进行加密
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(encryptStr))
	if err != nil {
		return "", err
	}

	//返回密文
	return base64.URLEncoding.EncodeToString(encryptedStr), nil
}

// PrivateDecrypt 私钥解密
func PrivateDecrypt(decryptStr string, path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	// 获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, _ = file.Read(buf)

	// pem 解码
	block, _ := pem.Decode(buf)

	// X509 解码
	//privateKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//privateKey, ok := privateKeyInterface.(*rsa.PrivateKey)
	//if !ok {
	//	return "", errors.New("rsa私钥解析失败")
	//}
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}
	//对密文进行解密
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decryptBytes)
	if err != nil {
		return "", err
	}
	//返回明文
	return string(decrypted), nil
}
