package milonga_render

import (
	"milonga/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SendHTML(c *fiber.Ctx, s string) error {
	c.Set("Content-Type", "text/html;charset=UTF-8")
	_, err := c.WriteString(s)
	return err
}

func SendHTMLFromFile(c *fiber.Ctx, filepath string) error {

	// read HTML file
	content := utils.OpenFile(filepath)

	c.Set("Content-Type", "text/html;charset=UTF-8")
	_, err := c.Write(content)
	return err
}
