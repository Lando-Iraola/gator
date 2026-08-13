-- +goose Up
CREATE TABLE feed_follows(
	id UUID primary key,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null,
	user_id UUID not null references users(id) on delete cascade,
	feed_id UUID not null references feeds(id) on delete cascade,
	CONSTRAINT uq_feed_user_pair UNIQUE (user_id, feed_id)
);


-- +goose Down
DROP TABLE feed_follows;
