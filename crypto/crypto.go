package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	//需要去加密的字符串
	plaintext := []byte("My name is Astaxie")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	// 使用 GCM 模式替代 CFB 模式
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Printf("Error: NewGCM = %s", err)
		os.Exit(-1)
	}

	// 生成随机 nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		fmt.Printf("Error: rand.Read = %s", err)
		os.Exit(-1)
	}

	//加密字符串
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字符串
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Printf("Error: ciphertext too short")
		os.Exit(-1)
	}

	nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintextCopy, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Printf("Error: gcm.Open = %s", err)
		os.Exit(-1)
	}
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}
