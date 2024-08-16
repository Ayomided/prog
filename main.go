package main

import (
	"database/sql"

	"github.com/Ayomided/prog.git/article"
	"github.com/Ayomided/prog.git/sqlite"
	"github.com/Ayomided/prog.git/views"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

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

func main() {
	dbConnection := setupDatabaseConnection()

	db := sqlite.New(dbConnection)

	e := echo.New()

	parser := article.NewParser()

	e.GET("/", func(c echo.Context) error {
		return HomeHandler(c, db)
	})
	e.GET("/articles/:slug", func(c echo.Context) error {
		return ArticleHandler(c, db, parser)
	})

	e.Logger.Fatal(e, e.Start(":8080"))
}

func setupDatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "blog.db")
	if err != nil {
		panic(err)
	}
	return db
}
