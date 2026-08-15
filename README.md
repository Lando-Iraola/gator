# Gator

This is a guided project from [Boot.dev](https://www.boot.dev/lessons/55d273f2-5d8f-4acc-9493-bf8be8aea700)  
This is a TUI Blog aggregator. It scrapes RSS feeds and displays the results in the CLI

## Dependencies

This project requires the following to run:

- Go
- Postgresql 18.4
- goose

## Prerequisite

Gator requires a running PostgreSQL database instance.  
Gator requires a .gatorconfig.json be present in the home directory of the OS user to point to the database instance.

```json
{
  "db_url": "postgres://postgres:@localhost:5432/gator?sslmode=disable",
  "current_user_name": "gator_db_user"
}
```

Gator requires goose for managing database migrations, to install it run the following:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

After the installation of goose we must run the database migrations  
Before we can do this, we must clone this repository, once done, inside the repository root directory we must run the following:

```bash
goose postgres postgres://postgres:@localhost:5432/gator -dir ./sql/schema up
```

## How to install

Once all the prerequisites have been fulfilled we can install gator:

```bash
go install github.com/Lando-Iraola/gator
```

## List of commands

| Command     | Arguments             | Description                                           |
| :---------- | :-------------------- | :---------------------------------------------------- |
| `login`     | `<user_name>`         | The name of a registered user                         |
| `register`  | `<user_name>`         | The name of a user to register (errors on duplicates) |
| `reset`     | _None_                | Resets state/database                                 |
| `users`     | _None_                | Lists all registered users                            |
| `agg`       | `<time_between_reqs>` | Time between scrapes (e.g., `30s`, `1m`, `1h`)        |
| `addfeed`   | `<name> <url>`        | Name and URL for the feed to add                      |
| `feeds`     | _None_                | Lists all feeds                                       |
| `follow`    | `<url>`               | URL of a feed to follow                               |
| `following` | _None_                | Lists all feeds followed by the logged-in user        |
| `unfollow`  | `<url>`               | URL of a feed to unfollow                             |
| `browse`    | `<limit>`             | Number of items to display (optional, default: `2`)   |

### login

Sets the currently logged in user

#### arguments

- user_name: The name of a registered user

##### example

gator login lando

### register

Register an users and logs them in

#### arguments

- user_name: The name of a user to register, errors on duplicates

##### example

gator register lando

### reset

Wipes the database from all its feeds, users and everything inbetween.

##### example

gator reset

### users

Lists the registered users in the database

##### example

gator users

### agg

Scrapes the registered feeds for later browsing, waiting with time between requests.

#### arguments

- time_between_reqs: Time between scrapes, in the format of 30s, 1m, 1m30s, 1h, etc

##### example

gator agg 30s

### addfeed

Adds a feed and the currently logged in user will automatically follow that new feed

#### arguments

- name: The user given name for this feed
- url: The user given url for this feed

##### example

gator addfeed "Test feed" "https://blog.boot.dev/index.xml"

### feeds

Lists all available feeds

##### example

```bash
gator feeds
```

### follow

This command allow an user to add to their existing feed the interests of other users.

#### arguments

- url: The url of a feed from another user

##### example

```bash
gator follow "https://blog.boot.dev/index.xml"
```

### following

This displays the feeds the currently logged in user is following

##### example

```bash
gator following
```

### unfollow

This command unfollow a feed from the logged in user

#### arguments

- url: The url of a feed to unfollow

##### example

```bash
gator unfollow "https://blog.boot.dev/index.xml"
```

### browse

This command displays the title, publishing date, description and link to the article from the logged in user's feed

#### arguments

- limit: The number of items to be displayed, by default it's 2

##### example

```bash
gator browse 5
```
