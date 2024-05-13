package controllers

import (
	"SoftTech/blogs/models"
	"SoftTech/blogs/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetListBlog(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": 200, "message": "This is GetListBlog() controllers!", "data": models.GetBlogList()})
}

func UpdateBlog(c *fiber.Ctx) error {
	params := new(types.Blog)
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Data request invalid", "data": err})
	}
	err := models.UpdateBlog(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Create category failed!", "error": err})
	}

	return c.JSON(fiber.Map{"status": 200, "message": "This is UpdateBlog() controllers!"})
}

func GetDetailBlog(c *fiber.Ctx) error {
	type Id struct {
		Id string `json:"id"`
	}
	params := new(Id)
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Data request invalid", "data": err})
	}

	id, err := strconv.ParseUint(params.Id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Id invalid"})
	}
	result, err := models.GetBlogDetails(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 401, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"status": 200, "message": "This is GetDetailBlog() controllers!", "data": result})
}

func DeleteBlog(c *fiber.Ctx) error {
	type BlogDelete struct {
		Id string `json:"id" validate:"required,min=19"`
	}
	params := new(BlogDelete)
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Data request invalid", "data": err})
	}

	id, err := strconv.ParseUint(params.Id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Id invalid"})
	}

	err = models.DeleteBlog(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "error": err})
	}
	return c.JSON(fiber.Map{"status": 200, "message": "This is DeleteBlog() controllers!"})
}

func CreateBlog(c *fiber.Ctx) error {
	params := new(types.Blog)
	if err := c.BodyParser(params); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Data request invalid"})
	}

	if err := models.CreateBlog(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 404, "message": "Create category failed!", "error": err})
	}
	return c.JSON(fiber.Map{"status": 200, "message": "This is CreateBlog() controllers!"})
}
