package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model  `json:"-"`
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReadTimes   uint   `json:"readTimes"`
}

var (
	db *gorm.DB
)

func main() {
	initDB()
	startServer()
}

func initDB() {
	var err error
	db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(Blog))
}

func newBlog(b Blog) error {
	b.ID = getSnakeCase(b.Name)
	b.ReadTimes = 0

	fb, err := getBlog(getSnakeCase(b.Name))
	if fb.Name == b.Name && err == nil {
		b.ID = fb.ID + "-" + uuid.New().String()[:8]
	}

	return db.Model(new(Blog)).Create(&b).Error
}

func updateBlog(newBlog Blog) error {
	_, err := getBlog(newBlog.ID)
	if err != nil {
		return err
	}

	newBlog.UpdatedAt = time.Now()

	return db.
		Model(new(Blog)).
		Where("id = ?", newBlog.ID).
		Updates(&newBlog).
		Error
}

func deleteBlog(id string) error {
	blog, err := getBlog(id)
	if err != nil {
		return err
	}

	return db.Model(new(Blog)).Delete(&blog, "id = ?", id).Error
}

func getBlog(id string) (Blog, error) {
	var blog Blog
	err := db.
		Model(new(Blog)).
		First(&blog, "id = ?", id).
		Error
	if err != nil {
		return Blog{}, err
	}

	blog.ReadTimes++
	err = db.
		Model(new(Blog)).
		Where("id = ?", blog.ID).
		Updates(&blog).
		Error
	if err != nil {
		return Blog{}, err
	}

	return blog, nil
}

func getBlogs() ([]Blog, error) {
	blogs := make([]Blog, 0)

	err := db.
		Model(new(Blog)).
		Find(&blogs).
		Error

	if err != nil {
		return nil, err
	}

	sort.Slice(blogs, func(i, j int) bool {
		return blogs[i].CreatedAt.After(blogs[j].CreatedAt)
	})

	return blogs, nil
}

func getSnakeCase(s string) string {
	sb := new(strings.Builder)

	for i := 0; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i])) {
			continue
		}
		if (unicode.IsUpper(rune(s[i]))) && i > 0 {
			sb.WriteRune('-')
		}
		if s[i] != ' ' && s[i] != '_' {
			sb.WriteRune(unicode.ToLower(rune(s[i])))
		}
	}

	return sb.String()
}

func startServer() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Static("/static", "./static")

	app.Use(func(c *fiber.Ctx) error {
		return cors.New(cors.Config{
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
			AllowOrigins: "*",
			AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		})(c)
	})

	blog := app.Group("/blog")
	blog.Get("/", func(c *fiber.Ctx) error {
		blogs, err := getBlogs()
		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(blogs)
	})

	blog.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		blog, err := getBlog(id)
		if err != nil || len(blog.ID) == 0 {
			return c.SendStatus(404)
		}

		return c.JSON(blog)
	})

	authBlog := app.Group("/blog")
	authBlog.Use(func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token != os.Getenv("AUTH") {
			return c.SendStatus(401)
		}
		return c.Next()
	})

	authBlog.Post("/", func(c *fiber.Ctx) error {
		var blog Blog

		err := c.BodyParser(&blog)
		if err != nil {
			return c.SendStatus(400)
		}

		return newBlog(blog)
	})

	authBlog.Put("/", func(c *fiber.Ctx) error {
		var blog Blog

		err := c.BodyParser(&blog)
		if err != nil {
			return c.SendStatus(400)
		}

		return updateBlog(blog)
	})

	authBlog.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return deleteBlog(id)
	})

	// project := app.Group("/project")

	app.Get("/save", func(c *fiber.Ctx) error {
		blogs, err := getBlogs()
		if err != nil {
			return err
		}

		f, _ := os.Create("blogs.csv")

		f.WriteString("name|description|read_times\n")

		for _, b := range blogs {
			fmt.Fprintf(f, "%s|%s|%d\n", b.Name, b.Description, b.ReadTimes)
		}

		err = f.Close()
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendStatus(200)
	})

	app.Listen(":8080")
}
