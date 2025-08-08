// 代码生成时间: 2025-08-08 08:43:33
package main

import (
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego"
    "net/http"
)

// ApiResponseFormatter defines a structure for formatting API responses
type ApiResponseFormatter struct {
    Data    interface{} `json:"data"`    // Data to be returned in the response
    Message string    `json:"message"` // Message to be returned in the response
    Success bool      `json:"success"`  // Indicates the success of the operation
}

// responseJSON formats the response to be returned to the client
func responseJSON(w http.ResponseWriter, statusCode int, resp ApiResponseFormatter) {
    respJSON, err := json.Marshal(resp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    fmt.Fprintf(w, "%s", respJSON)
}

// handler is the handler function for the API endpoint
func handler(w http.ResponseWriter, r *http.Request) {
    // Example usage of responseJSON
    resp := ApiResponseFormatter{
        Data:    "Sample Data",
        Message: "Success",
        Success: true,
    }
    responseJSON(w, http.StatusOK, resp)
}

func main() {
    beego.Router("/api/response", &handler{})
    beego.Run()
}