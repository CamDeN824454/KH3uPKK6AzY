// 代码生成时间: 2025-10-12 02:35:21
package models

import (
    "errors"
    "github.com/astaxie/beego/orm"
)

// Product is the data model for a product entity
type Product struct {
    ID       int    `orm:"auto"` // Auto increment ID
    Name     string `orm:"size(100)"` // Product name
    Price    float64 `orm:"index"` // Product price
    Stock    int    `orm:"index"` // Product stock quantity
    Category string `orm:"size(100)"` // Product category
    // You can add more fields as needed
}

// TableName returns the table name of the Product model
func (u *Product) TableName() string {
    return "products"
}

// AddProduct adds a new product to the database
func AddProduct(p *Product) (int64, error) {
    o := orm.NewOrm()
    // Begin transaction
    if _, err := o.Begin(); err != nil {
        return 0, err
    }
    defer o.Rollback()

    // Insert product into the database
    if _, err := o.Insert(p); err != nil {
        return 0, err
    }

    // Commit transaction
    if err := o.Commit(); err != nil {
        return 0, err
    }

    return p.ID, nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(id int) (*Product, error) {
    o := orm.NewOrm()
    p := &Product{ID: id}
    if err := o.Read(p); err != nil {
        return nil, err
    }
    return p, nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(p *Product) error {
    o := orm.NewOrm()
    if _, err := o.Update(p); err != nil {
        return err
    }
    return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
    o := orm.NewOrm()
    p := &Product{ID: id}
    if _, err := o.Delete(p); err != nil {
        return err
    }
    return nil
}
