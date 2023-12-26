package api

import (
	"api/config"
	"api/db"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var Client = config.Client

type PostResponse struct {
	Status  int            `json:"status"`
	Data    []db.PostModel `json:"data"`
	Count   int            `json:"count"`
	HasMore bool           `json:"hasMore"`
}

var DeletePost = func(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")

	_, err := Client.Post.FindUnique(db.Post.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	newData := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{}
	err := c.BodyParser(newData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//TODO: Research
	updateDoc, err := Client.Post.FindUnique(db.Post.ID.Equals(id)).Update(
		db.Post.Title.SetIfPresent(&newData.Title),
		db.Post.Body.SetIfPresent(&newData.Body),
	).Exec(ctx)
	return c.Status(200).JSON(updateDoc)
}

func GetPost(c *fiber.Ctx) error {
	ctx := context.Background()
	page, err := strconv.Atoi(c.Query("page"))
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "limit & page is required",
		})
	}
	count, err := Client.Post.FindMany().Exec(ctx)
	rows, err := Client.Post.FindMany().Take(limit).Skip((page - 1) * limit).Exec(ctx)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	res := PostResponse{
		Status:  200,
		Data:    rows,
		HasMore: len(count) > page*limit,
		Count:   len(rows),
	}
	return c.JSON(res)
}

func GetById(c *fiber.Ctx) error {
	id := c.Query("id")
	ctx := context.Background()
	row, err := Client.Post.FindUnique(
		db.Post.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return c.JSON(row)
}

func CreatePost(c *fiber.Ctx) error {
	p := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{}
	err := c.BodyParser(&p)
	if err != nil {
		return err
	}
	ctx := context.Background()
	row, err := Client.
		Post.CreateOne(
		db.Post.Title.Set(p.Title),
		db.Post.Body.Set(p.Body),
	).Exec(ctx)

	return c.Status(201).JSON(row)
}
