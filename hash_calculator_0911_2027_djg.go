// 代码生成时间: 2025-09-11 20:27:10
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "strings"
)

// CalculateSHA256 hashes a string and returns the SHA256 hexadecimal string
func CalculateSHA256(input string) (string, error) {
    // Create a new hash interface to calculate SHA256
    hash := sha256.New()
    // Write the input string to the hash
    if _, err := hash.Write([]byte(input)); err != nil {
        return "", err // Return an empty string and the error encountered
    }
    // Calculate the final hash and convert it to hexadecimal
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// main function that serves as the entry point of the program
func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: hash_calculator <input_string>")
        os.Exit(1)
    }
    fmt.Println("Calculating SHA256 hash...")
    // Extract the input string from command line arguments
    input := os.Args[1]
    // Trim any leading or trailing whitespace from the input
    input = strings.TrimSpace(input)

    hash, err := CalculateSHA256(input)
    if err != nil {
        // If an error occurs, print the error and exit
        fmt.Printf("Error calculating hash: %v
", err)
        os.Exit(1)
    }

    // Print the resulting hash
    fmt.Printf("The SHA256 hash of '%s' is: %s
", input, hash)
}
