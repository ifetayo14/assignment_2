package controller

import (
	"assignment_2/config"
	"assignment_2/structs"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItems(c *gin.Context) {
	var (
		items  []structs.Items
		orders structs.CreateOrders
	)

	db := config.GetDB()

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insertOrder := structs.Orders{
		CustomerName: orders.CustomerName,
		OrderedAt:    orders.OrderedAt,
	}

	db.Create(&insertOrder)
	orderID := insertOrder.ID

	for _, v := range orders.Item {
		item := structs.Items{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderId:     insertOrder.ID,
		}

		items = append(items, item)
	}

	result := db.Create(&items)
	log.Println(orderID, result.RowsAffected)

	responseData := structs.Orders{
		ID:           items[0].OrderId,
		CustomerName: orders.CustomerName,
		OrderedAt:    time.Now(),
		Item:         items,
	}

	c.JSON(http.StatusOK, responseData)

}

func GetItems(c *gin.Context) {
	db := config.GetDB()

	orders := []structs.Orders{}
	err := db.Preload("Item").Find(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func DeleteItems(c *gin.Context) {
	db := config.GetDB()

	strId := c.Param("orderID")
	id, _ := strconv.Atoi(strId)

	orders := structs.Orders{}
	items := structs.Items{}

	err := db.Where("order_id = ?", id).Delete(&items).Error
	rowsAff := db.Where("ID = ?", id).Delete(&orders).RowsAffected

	if rowsAff == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID not found",
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Deleted",
	})
}

func UpdateItems(c *gin.Context) {
	db := config.GetDB()

	strId := c.Param("orderID")
	id, _ := strconv.Atoi(strId)

	var orders = structs.Orders{}

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateOrder := structs.Orders{
		ID:           uint(id),
		CustomerName: orders.CustomerName,
		OrderedAt:    orders.OrderedAt,
		Item:         orders.Item,
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateOrder)

	c.JSON(http.StatusOK, updateOrder)
}
