-- name: QueryArticles :many
select * from articles;

-- name: QueryArticlesBySlug :one
select * from articles where slug=?;
