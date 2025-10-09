// 代码生成时间: 2025-10-09 16:32:20
package main

import (
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    \_ "github.com/go-sql-driver/mysql"
)

// Certificate represents the structure of a certificate
type Certificate struct {
    Id       int    `orm:"auto"`
    Name     string
    Issuer   string
    ValidTill string
}

// CertificateModel is the model interface for database operations
type CertificateModel interface {
    AddCertificate(certificate *Certificate) error
    GetCertificate(id int) (*Certificate, error)
    UpdateCertificate(id int, certificate *Certificate) error
    DeleteCertificate(id int) error
}

// certificateModel implements CertificateModel interface
type certificateModel struct{}

// AddCertificate adds a new certificate to the database
func (m *certificateModel) AddCertificate(certificate *Certificate) error {
    o := orm.NewOrm()
    _, err := o.Insert(certificate)
    return err
}

// GetCertificate retrieves a certificate by its ID
func (m *certificateModel) GetCertificate(id int) (*Certificate, error) {
    o := orm.NewOrm()
    certificate := &Certificate{Id: id}
    err := o.Read(certificate)
    if err != nil {
        return nil, err
    }
    return certificate, nil
}

// UpdateCertificate updates an existing certificate
func (m *certificateModel) UpdateCertificate(id int, certificate *Certificate) error {
    o := orm.NewOrm()
    _, err := o.Update(certificate)
    return err
}

// DeleteCertificate deletes a certificate by its ID
func (m *certificateModel) DeleteCertificate(id int) error {
    o := orm.NewOrm()
    _, err := o.Delete(&Certificate{Id: id})
    return err
}

// certificateController handles HTTP requests for certificate management
type certificateController struct {
    beego.Controller
}

// URLMapping sets the routing for certificateController
func init() {
    beego.Router("/certificate/add", &certificateController{}, "post:Add")
    beego.Router("/certificate/:id", &certificateController{}, "get:Get")
    beego.Router("/certificate/update/:id", &certificateController{}, "post:Update")
    beego.Router("/certificate/delete/:id", &certificateController{}, "get:Delete")
}

// Add adds a new certificate
func (c *certificateController) Add() {
    var certificate Certificate
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &certificate); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid JSON data"}
        c.ServeJSON()
        return
    }
    model := &certificateModel{}
    if err := model.AddCertificate(&certificate); err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to add certificate"}
    } else {
        c.Data["json"] = map[string]string{"success": "Certificate added successfully"}
    }
    c.ServeJSON()
}

// Get retrieves a certificate by ID
func (c *certificateController) Get() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    model := &certificateModel{}
    certificate, err := model.GetCertificate(id)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Certificate not found"}
    } else {
        c.Data["json"] = certificate
    }
    c.ServeJSON()
}

// Update updates an existing certificate
func (c *certificateController) Update() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    var certificate Certificate
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &certificate); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid JSON data"}
        c.ServeJSON()
        return
    }
    model := &certificateModel{}
    if err := model.UpdateCertificate(id, &certificate); err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to update certificate"}
    } else {
        c.Data["json"] = map[string]string{"success": "Certificate updated successfully"}
    }
    c.ServeJSON()
}

// Delete deletes a certificate by ID
func (c *certificateController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    model := &certificateModel{}
    if err := model.DeleteCertificate(id); err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to delete certificate"}
    } else {
        c.Data["json"] = map[string]string{"success": "Certificate deleted successfully"}
    }
    c.ServeJSON()
}

func main() {
    beego.Run()
}