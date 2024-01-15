package usercontrollers

import (
	"fmt"
	"log"
	"react_go_catalog_web/database"
	"react_go_catalog_web/helper"
	"react_go_catalog_web/models"
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("can't parse to body")
	}

	// check if pass < 6
	if len(data["password"].(string)) <= 6{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Password must be greater or equals to 6",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	// check if email already in database
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email Already Exist",
		})
	}

	user := models.User{
		Firstname: data["firstname"].(string),
		Lastname:  data["lastname"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
		Phone:     data["phone"].(string),
	}

	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":user,
		"message":"Account Created",
	})
}


func Login(c *fiber.Ctx)error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("can't parse to body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Is Doesn't Exist",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password Incorrect",
		})
	}

	token, err := helper.GenerateJWT(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login Success",
		"user": user,
	})
}
