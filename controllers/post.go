package controllers

import (
  "fmt"
  "math/rand"
  "github.com/gofiber/fiber/v2"
)

type Post struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Likes int32 `json:"likes"`
  Comments int32 `json:"comments"`
  ShareLink string `json:"shareLink"`
}

func GetPost(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

  dataList := []Post{ }
  var data Post
  const n int = 4
  for i:=0; i < n ; i++ {
    data = Post{
      Title: fmt.Sprintf("JWT auth with golang - %v", i+1),
      Description: "We are working on a new project template with golang with jwt authentication included",
      Likes: rand.Int31n(1024),
      Comments: rand.Int31n(1024),
      ShareLink: "https://github.com/smc181002/smc181002",
    }
    dataList = append(dataList, data)
  }

  return c.JSON(dataList)
}

func GetPrivatePost(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

  dataList := []Post{ }
  var data Post
  const n int = 4
  for i:=0; i < n ; i++ {
    data = Post{
      Title: fmt.Sprintf("Secret content with JWT and fiber - %v", i+1),
      Description: "This is secret content being sent by Agent Anya.",
      Likes: rand.Int31n(1024) + 256,
      Comments: rand.Int31n(1024) + 256,
      ShareLink: "https://github.com/ceoldevs/nvim",
    }
    dataList = append(dataList, data)
  }

  return c.JSON(dataList)
}
