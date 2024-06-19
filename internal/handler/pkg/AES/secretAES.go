package AES

import (
	"crypto/aes"
	"encoding/hex"
)

// AES 简单对称加密
// 主要提供用户登录账号密码加密传输.请保证前后端使用相同秘钥key

const (
	// cipher key
	key = "foodis32bitlongpassphraseimusing"

	// 补位占位符
	placeholder = "0"
)

// EncryptAES
//
// @Description: 加密
// @param plaintext 需要加密内容
// @return string 经过加密后的字符串
// @return int 对原字符串的补位长度
func EncryptAES(plaintext string) (string, int, error) {
	var (
		length int
	)

	// AES-CBC 进行加密的数据的长度必须是 16 的倍数，对长度不符的进行+ “0” 补位
	if len(plaintext)%16 != 0 {
		length = 16 - (len(plaintext) % 16)
		for i := 0; i < length; i++ {
			plaintext = plaintext + placeholder
		}
	}
	if c, err := aes.NewCipher([]byte(key)); err != nil {
		return "", 0, err
	} else {
		out := make([]byte, len(plaintext))
		c.Encrypt(out, []byte(plaintext))

		return string(out), length, nil
	}
}

// DecryptAES
//
//	@Description: 解密
//	@param ct 需要解密内容
//	@param placeholderLen 补位长度
//	@return string
func DecryptAES(ct string, placeholderLen int) (string, error) {
	ciphertext, _ := hex.DecodeString(ct)
	if c, err := aes.NewCipher([]byte(key)); err != nil {
		return "", err
	} else {
		pt := make([]byte, len(ciphertext))
		c.Decrypt(pt, ciphertext)

		return string(pt[:len(pt)-placeholderLen]), nil
	}
}
