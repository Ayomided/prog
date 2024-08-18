package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ayomided/prog.git/article"
	"github.com/Ayomided/prog.git/sqlite"
	"github.com/Ayomided/prog.git/views"
	"github.com/gorilla/feeds"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed data/*.db
var Data embed.FS

func RssHandler(ctx echo.Context, db *sqlite.Queries) error {
	feed := &feeds.Feed{
		Title:       "David Adediji | blog",
		Link:        &feeds.Link{Href: "https://prog.fly.dev"},
		Description: "Discussing",
		Author:      &feeds.Author{Name: "David Adediji", Email: "idavid.adediji@gmail.com"},
		Created:     time.Now(),
	}

	articles, err := db.QueryArticles(ctx.Request().Context())
	if err != nil {
		return err
	}

	var feedItems []*feeds.Item
	for _, article := range articles {
		feedItems = append(feedItems,
			&feeds.Item{
				Id:      fmt.Sprintf("tag: %v, %v:/articles/%v", "localhost", article.CreatedAt, article.Slug),
				Title:   article.Title,
				Link:    &feeds.Link{Href: fmt.Sprintf("https://prog.fly.dev/articles/%v", article.Slug)},
				Created: article.CreatedAt,
			})
	}
	feed.Items = feedItems
	rss, err := feed.ToRss()
	if err != nil {
		return err
	}
	ctx.Response().Header().Set(echo.HeaderContentType, "application/rss+xml")
	return ctx.String(http.StatusOK, rss)
}

func HomeHandler(ctx echo.Context, db *sqlite.Queries) error {
	articles, err := db.QueryArticles(ctx.Request().Context())
	projects, err := db.QueryProjects(ctx.Request().Context())

	if err != nil {
		return err
	}
	return views.Home(articles, projects).Render(ctx.Request().Context(), ctx.Response())
}

func ArticleHandler(ctx echo.Context, db *sqlite.Queries, parser article.Parser) error {
	slug := ctx.Param("slug")

	article, err := db.QueryArticlesBySlug(ctx.Request().Context(), slug)

	articleContent, err := parser.Parse(slug)
	if err != nil {
		return err
	}
	return views.Article(article.Title, articleContent).Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbConnection := setupDatabaseConnection()

	db := sqlite.New(dbConnection)

	e := echo.New()

	parser := article.NewParser()

	e.Static("/static", "static")

	// Route to handle HTMX requests
	e.GET("/content", func(c echo.Context) error {
		target := c.QueryParam("target")

		var content string
		switch target {
		case "anyone":
			content = `<h1 class="text text-anyone text-[#F3D8BD]">Hello there, I’m a designer who cares about making beautiful things that help people.</h1>`
		case "recruiters":
			content = `<h1 class="text text-recruiters text-[#F3D8BD]">I’m a product designer with 15 years of experience across brand and product, at companies large and small. I’m not actively looking for a new role.</h1>`
		case "design-directors":
			content = `<h1 class="text text-design-directors text-[#F3D8BD]">Content for Design Directors</h1>`
		case "product-designers":
			content = `<h1 class="text text-product-designers text-[#F3D8BD]">Content for Product Designers</h1>`
		case "product-managers":
			content = `<h1 class="text text-product-managers text-[#F3D8BD]">Content for Product Managers</h1>`
		case "engineers":
			content = `<h1 class="text text-engineers text-[#F3D8BD]">Content for Engineers</h1>`
		default:
			content = ``
		}

		return c.HTML(http.StatusOK, content)
	})

	e.GET("/", func(c echo.Context) error {
		return HomeHandler(c, db)
	})
	e.GET("/feed", func(c echo.Context) error {
		return RssHandler(c, db)
	})
	e.GET("/articles/:slug", func(c echo.Context) error {
		return ArticleHandler(c, db, parser)
	})

	e.Logger.Fatal(e, e.Start(":"+port))
}

const (
	dbPath = "data/blog.db"
)

func setupDatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return db
}
