// 代码生成时间: 2025-10-02 01:55:28
package main

import (
    "fmt"
    "math/rand"
    "time"
    "beego/logs"
    "beego/orm"
)

// TestDataGenerator struct to hold the testing data
type TestDataGenerator struct {
    // ORM struct fields go here
}

// NewTestDataGenerator creates a new instance of TestDataGenerator
func NewTestDataGenerator() *TestDataGenerator {
    return &TestDataGenerator{}
}

// GenerateData generates random test data
func (t *TestDataGenerator) GenerateData() ([]map[string]interface{}, error) {
    var testData []map[string]interface{}

    // Set random seed for reproducibility
    rand.Seed(time.Now().UnixNano())

    // Define the fields we want to generate data for
    fields := []string{"id", "name", "email", "age"}

    for i := 0; i < 10; i++ { // Generate 10 records for testing
        record := make(map[string]interface{})
        record["id"] = i + 1 // Simple incremental ID
        record["name"] = generateRandomName() // Generate a random name
        record["email"] = generateRandomEmail() // Generate a random email
        record["age"] = rand.Intn(100) // Random age between 0 and 100

        testData = append(testData, record)
    }

    return testData, nil
}

// generateRandomName generates a random name
func generateRandomName() string {
    names := []string{"John", "Jane", "Alice", "Bob"}
    return names[rand.Intn(len(names))]
}

// generateRandomEmail generates a random email
func generateRandomEmail() string {
    domains := []string{"@example.com", "@sample.org", "@test.net"}
    names := []string{"user", "admin", "support", "info"}
    return names[rand.Intn(len(names))] + domains[rand.Intn(len(domains))]
}

// main function to run the test data generator
func main() {
    err := orm.RunSyncdb("default", false, true)
    if err != nil {
        logs.Error("Failed to sync database: %v", err)
        return
    }

    testDataGenerator := NewTestDataGenerator()
    data, err := testDataGenerator.GenerateData()
    if err != nil {
        logs.Error("Failed to generate test data: %v", err)
        return
    }

    fmt.Println("Generated Test Data:")
    for _, record := range data {
        fmt.Println(record)
    }
}
