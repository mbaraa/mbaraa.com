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
	"github.com/hashicorp/go-memdb"
)

type Blog struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var (
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"blog": {
				Name: "blog",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
	db *memdb.MemDB
)

func main() {
	initDB()
	startServer()
}

func initDB() {
	var err error
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

func getWriteDB() *memdb.Txn {
	return db.Txn(true)
}

func getReadDB() *memdb.Txn {
	return db.Txn(false)
}

func newBlog(b Blog) error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	b.ID = getSnakeCase(b.Name)

	fb, _ := getBlog(getSnakeCase(b.Name))
	if fb.Name == b.Name {
		b.ID = fb.ID + "-" + uuid.New().String()[:8]
	}

	t := getWriteDB()
	err := t.Insert("blog", &b)
	if err != nil {
		return err
	}
	t.Commit()

	return nil
}

func updateBlog(newBlog Blog) error {
	b, err := getBlog(newBlog.ID)
	if err != nil {
		return err
	}

	if newBlog.Name != b.Name {
		b.Name = newBlog.Name
	}

	if newBlog.Description != b.Description {
		b.Description = newBlog.Description
	}

	b.UpdatedAt = time.Now()

	err = deleteBlog(b.ID)
	if err != nil {
		return err
	}

	t := getWriteDB()
	err = t.Insert("blog", &b)
	if err != nil {
		return err
	}
	t.Commit()

	return nil
}

func deleteBlog(id string) error {
	t := getWriteDB()
	t.Delete("blog", Blog{
		ID: id,
	})
	t.Commit()
	return nil
}

func getBlog(id string) (Blog, error) {
	row, err := getReadDB().First("blog", "id", id)
	if err != nil || row == nil {
		return Blog{}, err
	}
	return *row.(*Blog), nil
}

func getBlogs() ([]Blog, error) {
	blogs := make([]Blog, 0)

	it, err := getReadDB().Get("blog", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		blogs = append(blogs, *obj.(*Blog))
	}

	sort.Slice(blogs, func(i, j int) bool {
		return blogs[i].CreatedAt.Before(blogs[j].CreatedAt)
	})

	return blogs, nil
}

func getSnakeCase(s string) string {
	sb := new(strings.Builder)

	for i := 0; i < len(s); i++ {
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
		if token != os.Getenv("auth") {
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

		f.WriteString("name,description\n")

		for _, b := range blogs {
			fmt.Fprintf(f, "%s,%s\n", b.Name, b.Description)
		}

		err = f.Close()
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendStatus(200)
	})

	app.Listen(":8080")
}
