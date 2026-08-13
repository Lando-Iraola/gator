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

-- name: CreateFeedFollow :one
WITH feed_follow_insert AS (
	INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)
returning * 
)

select
	feed_follow_insert.id,
	feed_follow_insert.created_at,
	feed_follow_insert.updated_at,
	users.name as user,
	feeds.name as feed
from 
	feed_follow_insert
inner join users
on
	feed_follow_insert.user_id = users.id
inner join feeds
on 
	feed_follow_insert.feed_id = feeds.id;


-- name: FeedByURL :one
SELECT
	*
FROM
	feeds
where
	url = $1;

-- name: GetFeedFollowsForUser :many
select
	feeds.name as feed,
	users.name as user
from feed_follows
inner join 
	feeds
on
	feeds.id = feed_follows.feed_id
inner join
	users
on
	users.id = feed_follows.user_id
where
	feed_follows.user_id = $1;

