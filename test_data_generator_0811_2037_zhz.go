// 代码生成时间: 2025-08-11 20:37:39
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// TestData represents the structure of the test data.
type TestData struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Age       int    `json:"age"`
    CreatedOn string `json:"created_on"`
}

// GenerateTestData creates and returns an array of TestData.
func GenerateTestData(count int) ([]TestData, error) {
    if count <= 0 {
        return nil, errors.New("count must be greater than 0")
    }

    testData := make([]TestData, count)
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < count; i++ {
        testData[i] = TestData{
            ID:        i + 1,
            Name:      fmt.Sprintf("User%d", i+1),
            Email:     fmt.Sprintf("user%d@example.com", i+1),
            Age:       rand.Intn(100) + 1, // Random age between 1 and 100
            CreatedOn: time.Now().Format(time.RFC3339),
        }
    }

    return testData, nil
}

// main function to demonstrate the test data generator.
func main() {
    count := 10 // Number of test data entries to generate

    testData, err := GenerateTestData(count)
    if err != nil {
        fmt.Println("Error generating test data: ", err)
        return
    }

    // Convert the test data to JSON for output.
    jsonData, err := json.MarshalIndent(testData, "", "    ")
    if err != nil {
        fmt.Println("Error marshalling test data to JSON: ", err)
        return
    }

    fmt.Println(string(jsonData))
}
