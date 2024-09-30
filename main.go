package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
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

//go:embed all:article/assets/resume
var resumeFiles embed.FS

func RssHandler(ctx echo.Context, db *sqlite.Queries) error {
	feed := &feeds.Feed{
		Title:       "David Adediji | blog",
		Link:        &feeds.Link{Href: "https://prog.fly.dev"},
		Description: "Hello! I am David, I share my thoughts here",
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
	if err != nil {
		return err
	}
	return views.Home(articles).Render(ctx.Request().Context(), ctx.Response())
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
func AboutHandler(ctx echo.Context, db *sqlite.Queries) error {
	projects, err := db.QueryProjects(ctx.Request().Context())
	if err != nil {
		return err
	}
	return views.About(projects).Render(ctx.Request().Context(), ctx.Response())
}

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	dbConnection := setupDatabaseConnection()

// 	db := sqlite.New(dbConnection)

// 	e := echo.New()

// 	parser := article.NewParser()

// 	if _, err := os.Stat("article/assets/images"); err == nil {
// 		// Local development: serve files from the filesystem
// 		e.Static("/images", "article/assets/images")
// 	} else {
// 		// Production: serve files from the embedded filesystem
// 		fsys, err := fs.Sub(staticFiles, "article/assets/images")
// 		if err != nil {
// 			e.Logger.Fatal(err)
// 		}
// 		assetHandler := http.FileServer(http.FS(fsys))
// 		e.GET("/images/*", echo.WrapHandler(http.StripPrefix("/images/", assetHandler)))
// 	}

// 	e.GET("/", func(c echo.Context) error {
// 		return HomeHandler(c, db)
// 	})
// 	e.GET("/feed", func(c echo.Context) error {
// 		return RssHandler(c, db)
// 	})
// 	e.GET("/articles", func(c echo.Context) error {
// 		return ArticleListHandler(c, db)
// 	})
// 	e.GET("/about", func(c echo.Context) error {
// 		return AboutHandler(c, db)
// 	})
// 	e.GET("/articles/:slug", func(c echo.Context) error {
// 		return ArticleHandler(c, db, parser)
// 	})
// 	e.GET("/resume", func(c echo.Context) error {
// 		return c.Attachment("article/assets/resume/DAVIDADEDIJI-CV.pdf", "DAVIDADEDIJI-CV.pdf")
// 	})

// 	e.Logger.Fatal(e, e.Start(":"+port))
// }

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
