package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
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

//go:embed article/assets/images*
var staticFiles embed.FS

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
func ArticleListHandler(ctx echo.Context, db *sqlite.Queries) error {
	articles, err := db.QueryArticles(ctx.Request().Context())
	if err != nil {
		return err
	}
	return views.ArticleList(articles).Render(ctx.Request().Context(), ctx.Response())
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

	// Route to handle HTMX requests
	e.GET("/content", func(c echo.Context) error {
		target := c.QueryParam("target")

		var content string
		switch target {
		case "anyone":
			content = `<h1 class="text text-anyone">Hello there, I am an explorer and enjoy understanding how things work and how to build better.</h1>`
		case "recruiters":
			content = `<h1 class="text text-recruiters">Quick learning, emerging technologies, and rapid development.</h1>`
		case "engineers":
			content = `<h1 class="text text-engineers">\t\t\t\t vs "    " -> 👀</h1>`
		case "startup-founders":
			content = `<h1 class="text text-design-directors">I don't know who you are (yet), I don't know what your product is (yet). If you are looking for a doctor I can tell you I am not your person, but I have a very particular set of skills. Skills that I have acquired over a very long career. Skills that make me an asset for people like you. If you shoot me a message that will be the start of an amazing thing. I will be here, waiting for you.</h1>`
		case "product-managers":
			content = `<h1 class="text text-product-designers">Managers</h1>`
		case "ai-engineers":
			content = `<h1 class="text text-product-managers">GPT-5 Will be the best - <span class="italic align-middle text-sm">everyone</span></h1>`
		default:
			content = ``
		}

		return c.HTML(http.StatusOK, content)
	})

	if _, err := os.Stat("article/assets/images"); err == nil {
		// Local development: serve files from the filesystem
		e.Static("/images", "article/assets/images")
	} else {
		// Production: serve files from the embedded filesystem
		fsys, err := fs.Sub(staticFiles, "article/assets/images")
		if err != nil {
			e.Logger.Fatal(err)
		}
		assetHandler := http.FileServer(http.FS(fsys))
		e.GET("/images/*", echo.WrapHandler(http.StripPrefix("/images/", assetHandler)))
	}

	e.GET("/", func(c echo.Context) error {
		return HomeHandler(c, db)
	})
	e.GET("/feed", func(c echo.Context) error {
		return RssHandler(c, db)
	})
	e.GET("/articles", func(c echo.Context) error {
		return ArticleListHandler(c, db)
	})
	e.GET("/articles/:slug", func(c echo.Context) error {
		return ArticleHandler(c, db, parser)
	})
	e.GET("/resume", func(c echo.Context) error {
		return c.Attachment("article/assets/images/DAVIDADEDIJI-CV.pdf", "DAVIDADEDIJI-CV.pdf")
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
