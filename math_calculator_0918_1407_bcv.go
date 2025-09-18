// 代码生成时间: 2025-09-18 14:07:34
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math"

    "github.com/astaxie/beego"
)

// MathOperations defines the structure for mathematical operations
type MathOperations struct {
    Operand1 float64 `json:"operand1"`
    operand2 float64 `json:"operand2"`
    Operator  string `json:"operator"`
}

// Result defines the structure for the result of a mathematical operation
type Result struct {
    Result float64 `json:"result"`
    Error   string `json:"error"`
}

// Calculate performs the mathematical operation based on the operator provided
func (m *MathOperations) Calculate() *Result {
    var result float64
    var err error
    switch m.Operator {
    case "+":
        result = m.Operand1 + m.operand2
    case "-":
        result = m.Operand1 - m.operand2
    case "*":
        result = m.Operand1 * m.operand2
    case "/":
        if m.operand2 == 0 {
            return &Result{Error: "Cannot divide by zero"}
        }
        result = m.Operand1 / m.operand2
    default:
        return &Result{Error: "Invalid operator"}
    }
    return &Result{Result: result}
}

// PrepareJsonResponse prepares the JSON response for the math operation result
func PrepareJsonResponse(result *Result) ([]byte, error) {
    data, err := json.Marshal(result)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func main() {
    beego.Router("/math", &MathOperations{}, "post:Calculate")
    beego.Run()
}

// The Beego controller method that will handle the POST request for the math operations
func (m *MathOperations) Calculate() {
    resp, err := PrepareJsonResponse(m.Calculate())
    if err != nil {
        // Handle JSON marshaling error
        fmt.Println("Error marshaling JSON: ", err)
        beego.Error(err)
        return
    }

    // Set the content type and write the response to the client
    beego.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
    beego.Ctx.ResponseWriter.Write(resp)
}