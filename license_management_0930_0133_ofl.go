// 代码生成时间: 2025-09-30 01:33:38
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "os"
    "strings"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

type License struct {
    Id      int    `orm:"auto"`
    License string `orm:"size(255)"`
    Product string `orm:"size(255)"`
    Status  string `orm:"size(50)"`
}

// LicenseService is responsible for managing license operations.
type LicenseService struct {
    // Additional fields or methods can be added here.
}

// NewLicenseService creates a new instance of the LicenseService.
func NewLicenseService() *LicenseService {
    return &LicenseService{}
}

// AddLicense adds a new license to the system.
func (svc *LicenseService) AddLicense(license string, product string) (*License, error) {
    var l License
    l.License = license
    l.Product = product
    l.Status = "active"
    _, err := orm.Insert(&l)
    if err != nil {
        return nil, err
    }
    return &l, nil
}

// GetAllLicenses retrieves all licenses from the system.
func (svc *LicenseService) GetAllLicenses() ([]License, error) {
    var licenses []License
    _, err := orm.QueryTable(License{}).All(&licenses)
    if err != nil {
        return nil, err
    }
    return licenses, nil
}

// UpdateLicense updates the status of a license.
func (svc *LicenseService) UpdateLicense(id int, status string) error {
    o := orm.NewOrm()
    _, err := o.QueryTable(License{}).Filter("Id", id).Update(orm.Params{ "Status": status })
    if err != nil {
        return err
    }
    return nil
}

func main() {
    // Initialize Beego.
    beego.SetLogger("console")
    beego.SetLevel(beego.LevelTrace)

    // Register models.
    orm.RegisterModel(new(License))

    // Register the LicenseService controller.
    beego.Router("/license/add", &LicenseController{})
    beego.Router("/license/get", &LicenseController{})
    beego.Router("/license/update", &LicenseController{})

    // Start the Beego server.
    beego.Run()
}

// LicenseController handles HTTP requests for license operations.
type LicenseController struct {
    beego.Controller
}

// Post adds a new license.
func (c *LicenseController) Post() {
    var req struct {
        License string `json:"license"`
        Product string `json:"product"`
    }
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request"}
        c.ServeJSON(false)
        return
    }
    svc := NewLicenseService()
    license, err := svc.AddLicense(req.License, req.Product)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to add license"}
        c.ServeJSON(false)
        return
    }
    c.Data["json"] = map[string]interface{}{
        "license": license.License,
        "product": license.Product,
        "status": license.Status,
    }
    c.ServeJSON()
}

// Get retrieves all licenses.
func (c *LicenseController) Get() {
    svc := NewLicenseService()
    licenses, err := svc.GetAllLicenses()
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to retrieve licenses"}
        c.ServeJSON(false)
        return
    }
    c.Data["json"] = licenses
    c.ServeJSON()
}

// Put updates a license's status.
func (c *LicenseController) Put() {
    var req struct {
        Id    int    `json:"id"`
        Status string `json:"status"`
    }
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request"}
        c.ServeJSON(false)
        return
    }
    svc := NewLicenseService()
    err := svc.UpdateLicense(req.Id, req.Status)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to update license"}
        c.ServeJSON(false)
        return
    }
    c.Data["json"] = map[string]string{"message": "License updated successfully"}
    c.ServeJSON()
}