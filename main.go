package main

import (
	_ "github.com/mattn/go-sqlite3"
)

// func ArticleHandler(ctx echo.Context, db *sqlite.Queries, parser article.Parser) error {
// 	slug := ctx.Param("slug")

// 	article, err := db.QueryArticlesBySlug(ctx.Request().Context(), slug)

// 	articleContent, err := parser.Parse(slug)
// 	if err != nil {
// 		return err
// 	}
// 	return views.Article(article.Title, articleContent).Render(ctx.Request().Context(), ctx.Response())
// }
// func ArticleListHandler(ctx echo.Context, db *sqlite.Queries) error {
// 	articles, err := db.QueryArticles(ctx.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	return views.ArticleList(articles).Render(ctx.Request().Context(), ctx.Response())
// }
// func AboutHandler(ctx echo.Context, db *sqlite.Queries) error {
// 	projects, err := db.QueryProjects(ctx.Request().Context())
// 	if err != nil {
// 		return err
// 	}
// 	return views.About(projects).Render(ctx.Request().Context(), ctx.Response())
// }

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
