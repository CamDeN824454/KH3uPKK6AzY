// 代码生成时间: 2025-08-27 08:46:24
// interactive_chart_generator.go
package main

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    \_ "github.com/wcharczuk/go-chart" // Importing for side effects.
)

// Define models for storing chart data.
// Chart struct represents the data for a single chart.
type Chart struct {
    Id    int    `orm:"auto"`
    Title string
    Data  string `orm:"size(1024)"` // Store the serialized chart data.
}

// Register models.
func init() {
    orm.RegisterModel(new(Chart))
}

// ChartService provides functionality for chart operations.
type ChartService struct {
}

// NewChartService creates a new instance of ChartService.
func NewChartService() *ChartService {
    return &ChartService{}
}

// CreateChart saves a new chart to the database.
func (s *ChartService) CreateChart(title string, data *chart.Chart) error {
    // Serialize the chart data before storing.
    serializedData, err := data.MarshalSVG()
    if err != nil {
        return err
    }
    chart := Chart{Title: title, Data: serializedData}
    _, err = orm.Insert(&chart)
    return err
}

// GetChart retrieves a chart from the database by its ID.
func (s *ChartService) GetChart(chartId int) (*Chart, error) {
    chart := Chart{Id: chartId}
    err := orm.Read(&chart)
    if err != nil {
        return nil, err
    }
    return &chart, nil
}

func main() {
    beego.RunMode = "dev"
    beego.Router("/chart", &ChartController{})
    beego.Run()
}

// ChartController handles HTTP requests related to charts.
type ChartController struct {
    beego.Controller
}

// Post handles the creation of a new chart.
func (c *ChartController) Post() {
    title := c.GetString("title")
    // Assume data is provided in the request body and can be parsed into a chart.Chart object.
    // For simplicity, this example does not handle the actual parsing of chart data.
    // In a real-world scenario, you would parse the request body to create the chart.
    
    service := NewChartService()
    err := service.CreateChart(title, nil) // Replace nil with actual chart data.
    if err != nil {
        c.CustomAbort(500, "Failed to create chart")
        return
    }
    c.Data["json"] = map[string]string{
        "message": "Chart created successfully",
    }
    c.ServeJSON()
}

// Get handles the retrieval of a chart by its ID.
func (c *ChartController) Get() {
    chartIdStr := c.GetString(":id")
    chartId, _ := strconv.Atoi(chartIdStr)
    service := NewChartService()
    chart, err := service.GetChart(chartId)
    if err != nil {
        c.CustomAbort(404, "Chart not found")
        return
    }
    c.Data["json"] = map[string]interface{}{
        "id":    chart.Id,
        "title": chart.Title,
        "data":  chart.Data,
    }
    c.ServeJSON()
}