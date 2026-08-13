-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;


-- name: FeedList :many
SELECT
	feeds.id,
	feeds.name as feed,
	url,
	feeds.created_at,
	feeds.updated_at,
	users.name as user
FROM
	feeds
INNER JOIN
	users
ON
	feeds.user_id = users.id;
