// 代码生成时间: 2025-09-22 09:35:47
package main

import (
    "beego/logs"
    "encoding/json"
    "html/template"
    "net/http"
    "strings"
)

// ChartData defines the structure for chart data
type ChartData struct {
    Labels []string      "json:\"labels\""
    Values []float64    "json:\"values\""
}

// Response defines the structure for API response
type Response struct {
    Status  string        "json:\"status\""
    Data    ChartData      "json:\"data\""
}

// GenerateChartHandler handles chart generation request
func GenerateChartHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }

    var requestData ChartData
    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Validate requestData
    if len(requestData.Labels) != len(requestData.Values) {
        http.Error(w, "Labels and values must be of the same length", http.StatusBadRequest)
        return
    }

    // Generate chart HTML
    chartHTML := generateChartHTML(requestData)

    // Write chart HTML to response
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(chartHTML))
}

// generateChartHTML generates chart HTML using chart.js
func generateChartHTML(data ChartData) string {
    var html strings.Builder
    html.WriteString("<html><head><title>Interactive Chart</title></head><body>")
    html.WriteString("<canvas id=\"chartCanvas\"></canvas>")
    html.WriteString("<script src=\"https://cdn.jsdelivr.net/npm/chart.js"></script>")
    html.WriteString("<script>")
    html.WriteString("var ctx = document.getElementById('chartCanvas').getContext('2d');")
    html.WriteString("var myChart = new Chart(ctx, {")
    html.WriteString("type: 'line',")
    html.WriteString(