// 代码生成时间: 2025-09-12 18:03:52
package main

import (
    "os"
    "fmt"
    "log"
    "math/rand"
    "time"
    "strconv"
    "beego"
)

// 定义一个用于测试数据生成的结构体
type TestData struct {
    ID        int
    Name      string
    Age       int
    Gender    string
    Email     string
    CreatedAt int64
}

// 初始化随机种子
func init() {
    rand.Seed(time.Now().UnixNano())
}

// GenerateTestData 生成测试数据
func GenerateTestData(count int) ([]TestData, error) {
    var data []TestData
    for i := 0; i < count; i++ {
        // 随机生成姓名
        name := randomName()
        // 随机生成年龄
        age := rand.Intn(100)
        // 随机生成性别
        gender := randomGender()
        // 随机生成邮箱
        email := fmt.Sprintf("%s@example.com", name)
        // 记录创建时间
        createdAt := time.Now().Unix()

        data = append(data, TestData{
            ID:        i + 1,
            Name:      name,
            Age:       age,
            Gender:    gender,
            Email:     email,
            CreatedAt: createdAt,
        })
    }
    return data, nil
}

// randomName 生成随机姓名
func randomName() string {
    var names = []string{"Alice", "Bob", "Charlie", "David", "Eve"}
    index := rand.Intn(len(names))
    return names[index]
}

// randomGender 生成随机性别
func randomGender() string {
    var genders = []string{"Male", "Female"}
    index := rand.Intn(len(genders))
    return genders[index]
}

// SaveTestData 将测试数据保存到文件
func SaveTestData(data []TestData) error {
    // 创建文件
    file, err := os.Create("testdata.csv")
    if err != nil {
        return err
    }
    defer file.Close()

    // 写入CSV头部
    _, err = file.WriteString("ID,Name,Age,Gender,Email,CreatedAt\
")
    if err != nil {
        return err
    }

    // 写入测试数据
    for _, item := range data {
        line := fmt.Sprintf("%d,%s,%d,%s,%s,%d\
", item.ID, item.Name, item.Age, item.Gender, item.Email, item.CreatedAt)
        _, err = file.WriteString(line)
        if err != nil {
            return err
