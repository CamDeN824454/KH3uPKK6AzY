// 代码生成时间: 2025-08-28 00:41:26
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "log"
    "strings"
)

// EncryptPassword encrypts the password using AES-256-CBC
func EncryptPassword(password string) (string, error) {
    key := []byte("your_encryption_key_here") // Replace with your encryption key
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    passwordBytes := []byte(password)
    blockSize := block.BlockSize()
    padding := blockSize - len(passwordBytes)%blockSize
    paddedPassword := append(passwordBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

    iv := make([]byte, aes.BlockSize)
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    cipherText := make([]byte, len(paddedPassword))
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText, paddedPassword)

    return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptPassword decrypts the password using AES-256-CBC
func DecryptPassword(encryptedPassword string) (string, error) {
    key := []byte("your_encryption_key_here") // Replace with your encryption key
    encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedPassword)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    if len(encryptedBytes)%aes.BlockSize != 0 {
        return "", fmt.Errorf("ciphertext is not a multiple of the block size")
    }

    iv := encryptedBytes[:aes.BlockSize]
    cipherText := encryptedBytes[aes.BlockSize:]

    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(cipherText, cipherText)

    // Remove padding
    padding := int(cipherText[len(cipherText)-1])
    decryptedPassword := cipherText[:len(cipherText)-padding]

    return string(decryptedPassword), nil
}

func main() {
    password := "mysecretpassword"
    encrypted, err := EncryptPassword(password)
    if err != nil {
        log.Fatalf("Error encrypting password: %v", err)
    }
    fmt.Printf("Encrypted password: %s
", encrypted)

    decrypted, err := DecryptPassword(encrypted)
    if err != nil {
        log.Fatalf("Error decrypting password: %v", err)
    }
    fmt.Printf("Decrypted password: %s
", decrypted)
}
