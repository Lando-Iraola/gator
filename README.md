# Gator

This is a guided project from [Boot.dev](https://www.boot.dev/lessons/55d273f2-5d8f-4acc-9493-bf8be8aea700)  
This is a TUI Blog aggregator. It scrapes RSS feeds and displays the results in the CLI

## Dependencies

This project requires the following to run:

- Go
- Postgresql 18.4

## How to run

At the root of the project run

```bash
go run . {command} {args}
```

## List of commands

### login

#### arguments

- user_name: The name of a registered user

### register

#### arguments

- user_name: The name of a user to register, errors on duplicates

### reset

### users

### agg

- arguments
  - time_between_reqs: Time between scrapes, in the format of 30s, 1m, 1m30s, 1h, etc

### addfeed

#### arguments

- name: The user given name for this feed
- url: The user given url for this feed

### feeds

### follow

This command allow an user to add to their existing feed the interests of other users.

#### arguments

- url: The url of a feed from another user

### following

This command lists the feeds the logged in user is following

### unfollow

This command unfollow a feed from the logged in user

#### arguments

- url: The url of a feed to unfollow

### browse

This command displays the title, publishing date, description and link to the article from the logged in user's feed

#### arguments

- limit: The number of items to be displayed, by default it's 2
