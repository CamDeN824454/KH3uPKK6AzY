// 代码生成时间: 2025-09-06 07:36:05
package main

import (
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    ID         int    `orm:"auto"`
    Name       string
    Quantity   int
    Price      float64
}

// InventoryController handles requests related to inventory items
type InventoryController struct {
    beego.Controller
}

// Prepare function to set up the controller
func (c *InventoryController) Prepare() {
    // Check if the user is authenticated
    // For demonstration purposes, we assume authentication is handled elsewhere
}

// GetItems returns all inventory items
func (c *InventoryController) GetItems() {
    var items []InventoryItem
    if _, err := orm.NewOrm().QueryTable("inventory_item").All(&items); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Failed to retrieve items"}
        c.Ctx.Output.Status = 500
        c.ServeJSON()
        return
    }
    c.Data["json"] = items
    c.ServeJSON()
}

// GetItem returns a specific inventory item by ID
func (c *InventoryController) GetItem() {
    id := c.Ctx.Input.Param(":id")
    if id == "" {
        c.Data["json"] = map[string]interface{}{"error": "Invalid ID"}
        c.ServeJSON()
        return
    }
    item := InventoryItem{ID: atoi(id)}
    if err := orm.NewOrm().Read(&item); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Item not found"}
        c.Ctx.Output.Status = 404
        c.ServeJSON()
        return
    }
    c.Data["json"] = item
    c.ServeJSON()
}

// AddItem adds a new inventory item
func (c *InventoryController) AddItem() {
    var item InventoryItem
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &item); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Invalid data"}
        c.ServeJSON()
        return
    }
    if _, err := orm.NewOrm().Insert(&item); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Failed to add item"}
        c.Ctx.Output.Status = 500
        c.ServeJSON()
        return
    }
    c.Data["json"] = map[string]interface{}{"success": "Item added"}
    c.ServeJSON()
}

// UpdateItem updates an existing inventory item
func (c *InventoryController) UpdateItem() {
    id := c.Ctx.Input.Param(":id")
    if id == "" {
        c.Data["json"] = map[string]interface{}{"error": "Invalid ID"}
        c.ServeJSON()
        return
    }
    var item InventoryItem
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &item); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Invalid data"}
        c.ServeJSON()
        return
    }
    item.ID = atoi(id)
    if _, _, err := orm.NewOrm().Update(&item); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Failed to update item"}
        c.Ctx.Output.Status = 500
        c.ServeJSON()
        return
    }
    c.Data["json"] = map[string]interface{}{"success": "Item updated"}
    c.ServeJSON()
}

// DeleteItem deletes an inventory item
func (c *InventoryController) DeleteItem() {
    id := c.Ctx.Input.Param(":id")
    if id == "" {
        c.Data["json"] = map[string]interface{}{"error": "Invalid ID"}
        c.ServeJSON()
        return
    }
    item := InventoryItem{ID: atoi(id)}
    if _, err := orm.NewOrm().Delete(&item); err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Failed to delete item"}
        c.Ctx.Output.Status = 500
        c.ServeJSON()
        return
    }
    c.Data["json"] = map[string]interface{}{"success": "Item deleted"}
    c.ServeJSON()
}

// atoi converts a string to an integer
func atoi(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        return 0
    }
    return i
}

func main() {
    beego.Router("/inventory/items", &InventoryController{}, "get:GetItems;post:AddItem")
    beego.Router("/inventory/item/:id", &InventoryController{}, "get:GetItem;put:UpdateItem;delete:DeleteItem")
    beego.Run()
}