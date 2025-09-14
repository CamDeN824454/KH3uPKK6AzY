// 代码生成时间: 2025-09-14 16:48:02
package main

import (
    "beego框架/beeweb"
    "encoding/json"
    "fmt"
    "os"
)

// CartItem 代表购物车中的单个商品项
type CartItem struct {
    ProductID int    "json:product_id"
    Quantity  int    "json:quantity"
    Price     float64 "json:price"
}

// ShoppingCart 代表用户的购物车，包含多个商品项
type ShoppingCart struct {
    Items map[int]CartItem `json:items`
}

// AddItem 添加商品到购物车
func (c *ShoppingCart) AddItem(productID int, quantity int, price float64) error {
    if _, exists := c.Items[productID]; exists {
        c.Items[productID].Quantity += quantity
    } else {
        c.Items[productID] = CartItem{
            ProductID: productID,
            Quantity:  quantity,
            Price:     price,
        }
    }
    return nil
}

// RemoveItem 从购物车删除商品
func (c *ShoppingCart) RemoveItem(productID int) error {
    if _, exists := c.Items[productID]; !exists {
        return fmt.Errorf("product with ID %d does not exist in the cart", productID)
    }
    delete(c.Items, productID)
    return nil
}

// GetItem 获取购物车中的商品项
func (c *ShoppingCart) GetItem(productID int) (*CartItem, error) {
    item, exists := c.Items[productID]
    if !exists {
        return nil, fmt.Errorf("product with ID %d does not exist in the cart", productID)
    }
    return &item, nil
}

// CalculateTotal 计算购物车总价
func (c *ShoppingCart) CalculateTotal() float64 {
    total := 0.0
    for _, item := range c.Items {
        total += float64(item.Quantity) * item.Price
    }
    return total
}

// SetWebRoutes 设置路由
func SetWebRoutes() {
    web.Router("/add_item", &controllers.ShoppingCartController{}, "post:AddItem")
    web.Router("/remove_item", &controllers.ShoppingCartController{}, "post:RemoveItem")
    web.Router("/get_item", &controllers.ShoppingCartController{}, "get:GetItem")
    web.Router("/calculate_total", &controllers.ShoppingCartController{}, "get:CalculateTotal")
}

func main() {
    // 初始化购物车
    cart := ShoppingCart{Items: make(map[int]CartItem)}

    // 设置路由
    SetWebRoutes()

    // 启动Web服务
    web.Run(":8080")
}

// ShoppingCartController 控制器，处理购物车相关的请求
type ShoppingCartController struct {
    web.Controller
}

// AddItem 添加商品到购物车
func (c *ShoppingCartController) AddItem() {
    var cartItem CartItem
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &cartItem); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request body"}
        c.ServeJSON(true)
        return
    }
    if err := cart.AddItem(cartItem.ProductID, cartItem.Quantity, cartItem.Price); err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON(true)
        return
    }
    c.Data["json"] = map[string]string{"message": "Item added to cart successfully"}
    c.ServeJSON(true)
}

// RemoveItem 从购物车删除商品
func (c *ShoppingCartController) RemoveItem() {
    var productID int
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &productID); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request body"}
        c.ServeJSON(true)
        return
    }
    if err := cart.RemoveItem(productID); err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON(true)
        return
    }
    c.Data["json"] = map<string]string{"message": "Item removed from cart successfully"}
    c.ServeJSON(true)
}

// GetItem 获取购物车中的商品项
func (c *ShoppingCartController) GetItem() {
    var productID int
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &productID); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request body"}
        c.ServeJSON(true)
        return
    }
    item, err := cart.GetItem(productID)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON(true)
        return
    }
    c.Data["json"] = map[string]CartItem{ "item": *item }
    c.ServeJSON(true)
}

// CalculateTotal 计算购物车总价
func (c *ShoppingCartController) CalculateTotal() {
    total := cart.CalculateTotal()
    c.Data["json"] = map[string]float64{ "total": total }
    c.ServeJSON(true)
}