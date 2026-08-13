-- +goose Up
CREATE TABLE feeds (
	id UUID primary key,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null,
	name TEXT not null,
	url TEXT not null unique,
	user_id UUID not null,
	constraint user_feed foreign key(user_id) references users(id) on delete cascade
);

-- +goose Down
DROP TABLE feeds;
