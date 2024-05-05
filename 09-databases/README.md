# Databases

The goal of this exercise is to practice implementing a persistance layer for an existing HTTP API.

## Steps

### Docker Compose

We will be using PostgreSQL as our RDBMS.

1. Postgres will have to be added to our compose file.
    - Postgres has it's own [image](https://hub.docker.com/_/postgres) on Docker Hub.
2. [Optional] To enable persistance, you will have to mount a volume for the Postgres to store the data in.
    - Search for `PGDATA` in the Docker Hub page.

### Persistance

For our database driver, we will be using the [pgx](https://github.com/jackc/pgx) library in conjunction with the standard `databases/sql` package interfaces.

1. A connection to the database needs to be created when initialization the app.
    - That should generally be done in the `main` function.
    - Consider retrieving the database URL using a environmental variable.
    - Change the `Repository` struct structure and wire the connection into it.
2. Reimplement the repository methods to use the database.
3. [Optional] Test out that persisting the data actually works by shutting the app down and up again.

This is a [example usage](https://github.com/jackc/pgx?tab=readme-ov-file#example-usage) from the [pgx library](https://github.com/jackc/pgx) to get you started:

```
func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
```
