// 代码生成时间: 2025-08-24 22:06:51
package main

import (
    "fmt"
    "strings"
    "regexp"
    "beego/validation"
)

// Define a struct with fields that correspond to the form fields we want to validate
type MyForm struct {
    Email    string `form:"email"`
    Age      string `form:"age"`
    Password string `form:"password"`
}

// Define a validator function for the form
func (mf *MyForm) Valid(v *validation.Validation) {
    // Validate Email
    if ok, _ := regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`, mf.Email); !ok {
        v.SetError("Email", "Email format is not correct")
    }

    // Validate Age, should be a number and between 18 and 100
    if _, err := strconv.Atoi(mf.Age); err != nil || mf.Age == "" || len(mf.Age) == 1 {
        v.SetError("Age", "Age must be a number and greater than 0")
    } else if age, _ := strconv.Atoi(mf.Age); age < 18 || age > 100 {
        v.SetError("Age", "Age must be between 18 and 100")
    }

    // Validate Password, should be at least 6 characters long
    if len(mf.Password) < 6 {
        v.SetError("Password", "Password must be at least 6 characters long")
    }
}

func main() {
    // Create a new instance of MyForm
    form := MyForm{Email: "test@example.com", Age: "25", Password: "password123"}

    // Create a validation object
    valid := validation.Validation{}

    // Check if the form is valid
    if b, _ := valid.Valid(form); !b {
        // If not valid, print the error messages
        for _, err := range valid.Errors {
            fmt.Printf("%s: %s
", err.Field, err.Message)
        }
    } else {
        // If valid, print a success message
        fmt.Println("The form is valid!")
    }
}
