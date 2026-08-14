-- +goose Up
CREATE TABLE POSTS (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TiMESTAMP NOT NULL,
	title TEXT NOT NULL,
	url TEXT NOT NULL,
	description TEXT,
	published_at TIMESTAMP,
	feed_id UUID references feeds(id) not null,
	UNIQUE (title, url, feed_id)
);

-- +goose Down
DROP TABLE POSTS;
