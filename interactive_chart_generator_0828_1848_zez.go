// 代码生成时间: 2025-08-28 18:48:21
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/astaxie/beego"
)

// ChartData holds the data for a chart
type ChartData struct {
    Labels   []string `json:"labels"`
    Data     []int    `json:"data"`
    ChartName string  `json:"chart_name"`
}

// ChartService is a service to handle chart generation
type ChartService struct {
    data []ChartData
}

// NewChartService creates a new ChartService with empty data
func NewChartService() *ChartService {
    return &ChartService{data: make([]ChartData, 0)}
}

// AddChartData appends a new chart data to the service
func (cs *ChartService) AddChartData(chartData ChartData) error {
    cs.data = append(cs.data, chartData)
    return nil
}

// GetChartData returns the current chart data
func (cs *ChartService) GetChartData() ([]ChartData, error) {
    return cs.data, nil
}

// ChartController handles HTTP requests related to chart data
type ChartController struct {
    beego.Controller
    cs *ChartService
}

// Prepare overrides beego.Controller's Prepare method
func (cc *ChartController) Prepare() {
    cc.cs = NewChartService()
}

// AddChartData handles POST requests to add new chart data
func (cc *ChartController) AddChartData() {
    var chartData ChartData
    if err := json.Unmarshal(cc.Ctx.Input.RequestBody, &chartData); err != nil {
        cc.CustomAbort(400, fmt.Sprintf("Error unmarshalling chart data: %s", err.Error()))
        return
    }
    if err := cc.cs.AddChartData(chartData); err != nil {
        cc.CustomAbort(500, fmt.Sprintf("Error adding chart data: %s", err.Error()))
        return
    }
    cc.Data["json"] = map[string]string{"message": "Chart data added successfully"}
    cc.ServeJSON()
}

// GetChartData handles GET requests to retrieve chart data
func (cc *ChartController) GetChartData() {
    chartData, err := cc.cs.GetChartData()
    if err != nil {
        cc.CustomAbort(500, fmt.Sprintf("Error retrieving chart data: %s", err.Error()))
        return
    }
    cc.Data["json"] = chartData
    cc.ServeJSON()
}

func main() {
    // Set up the Beego router
    beego.Router("/chart/add", &ChartController{}, "post:AddChartData")
    beego.Router("/chart/get", &ChartController{}, "get:GetChartData")

    // Start the Beego server
    beego.Run()
}
