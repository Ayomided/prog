package main

import (
	"database/sql"
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

func RssHandler(ctx echo.Context, db *sqlite.Queries) error {
	feed := &feeds.Feed{
		Title:       "David Adediji | blog",
		Link:        &feeds.Link{Href: "localhost:8080"},
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
				Link:    &feeds.Link{Href: fmt.Sprintf("localhost:8080/articles/%v", article.Slug)},
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

func setupDatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "blog.db")
	if err != nil {
		panic(err)
	}
	return db
}
