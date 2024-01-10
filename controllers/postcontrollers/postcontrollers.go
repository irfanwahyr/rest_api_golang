package postcontrollers

import (
	"errors"
	"fmt"
	"math"
	"react_go_catalog_web/database"
	"react_go_catalog_web/helper"
	"react_go_catalog_web/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var contentpost models.Content
	if err := c.BodyParser(&contentpost); err != nil {
		fmt.Println("can't parse to body")
	}

	if err := database.DB.Create(&contentpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Create Content",
	})
}

func ReadPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getcontent []models.Content
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getcontent)
	database.DB.Model(&models.Content{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getcontent,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var contentpost models.Content
	database.DB.Where("id=?", id).Preload("User").First(&contentpost)
	return c.JSON(fiber.Map{
		"data": contentpost,
	})
}

func UpdatePost(c *fiber.Ctx) error  {
	id, _ := strconv.Atoi(c.Params("id"))
	content := models.Content{
		Id:uint(id),
	}
	if err := c.BodyParser(&content); err != nil {
		fmt.Println("Unable To Parse body")
	}
	database.DB.Model(&content).Updates(content)
	return c.JSON(fiber.Map{
		"message": "Update Content Success",
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	content := models.Content{
		Id:uint(id),
	}
	deletecontent := database.DB.Delete(&content)
	if errors.Is(deletecontent.Error, gorm.ErrRecordNotFound){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success Delete Data",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := helper.ParseJWT(cookie)
	var content []models.Content

	// Hitung jumlah user ID yang sama
	var count int64
	database.DB.Model(&content).Where("userid=?", id).Count(&count)

	// Periksa apakah ada duplikat atau tidak
	var message string
	if count > 0 {
		message = "There is a duplicate user"
	} else {
		message = "There is no duplicate user"
	}

	return c.JSON(fiber.Map{
		"message": message,
	})
}
