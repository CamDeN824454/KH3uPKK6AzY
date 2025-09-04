// 代码生成时间: 2025-09-04 16:58:57
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "fmt"
    "os"
)

// PasswordTool 结构体，用于存储密钥
type PasswordTool struct {
    key []byte
}

// NewPasswordTool 创建一个新的PasswordTool实例
func NewPasswordTool(key string) *PasswordTool {
    return &PasswordTool{
        key: []byte(key),
    }
}

// Encrypt 加密密码
func (pt *PasswordTool) Encrypt(plainText string) (string, error) {
    if len(pt.key) != 32 {
        return "", errors.New("密钥必须是32字节")
    }
    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }
    plainTextData := pad([]byte(plainText), aes.BlockSize)
    cipherText := make([]byte, aes.BlockSize+len(plainTextData))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText[aes.BlockSize:], plainTextData)
    return hex.EncodeToString(cipherText), nil
}

// Decrypt 解密密码
func (pt *PasswordTool) Decrypt(cipherText string) (string, error) {
    if len(pt.key) != 32 {
        return "", errors.New("密钥必须是32字节")
    }
    cipherTextData, err := hex.DecodeString(cipherText)
    if err != nil {
        return "", err
    }
    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }
    if len(cipherTextData) < aes.BlockSize {
        return "", errors.New("密文长度不足")
    }
    iv := cipherTextData[:aes.BlockSize]
    cipherTextData = cipherTextData[aes.BlockSize:]
    mode := cipher.NewCBCDecrypter(block, iv)
    if mode == nil {
        return "", errors.New("解密失败")
    }
    mode.CryptBlocks(cipherTextData, cipherTextData)
    cipherTextData = unPad(cipherTextData, aes.BlockSize)
    return string(cipherTextData), nil
}

// pad 填充数据以满足块大小要求
func pad(buf []byte, blockSize int) []byte {
    padding := blockSize - len(buf)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(buf, padtext...)
}

// unPad 移除填充数据
func unPad(buf []byte, blockSize int) []byte {
    length := len(buf)
    unpadding := int(buf[length-1])
    return buf[:(length - unpadding)]
}

func main() {
    var key string
    fmt.Print("请输入32字节的密钥: ")
    fmt.Scanln(&key)
    pt := NewPasswordTool(key)
    fmt.Print("请输入要加密的密码: ")
    plainText := ""
    fmt.Scanln(&plainText)
    encrypted, err := pt.Encrypt(plainText)
    if err != nil {
        fmt.Println("加密失败: ", err)
        return
    }
    fmt.Printf("加密后的密码: %s
", encrypted)
    fmt.Print("请输入要解密的密码: ")
    cipherText := ""
    fmt.Scanln(&cipherText)
    decrypted, err := pt.Decrypt(cipherText)
    if err != nil {
        fmt.Println("解密失败: ", err)
        return
    }
    fmt.Printf("解密后的密码: %s
", decrypted)
}