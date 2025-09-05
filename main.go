package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // Postgres driver

	"github.com/sanskarchoudhry/gator/internal/cli"
	"github.com/sanskarchoudhry/gator/internal/config"
	"github.com/sanskarchoudhry/gator/internal/database"
)

func main() {
	// Load config (.gatorconfig.json must include ?sslmode=disable in db_url)
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Open DB connection
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// Wrap db with generated queries
	dbq := database.New(db)

	// App state
	s := &cli.State{
		Cfg: &cfg,
		DB:  dbq,
	}

	// Register commands
	cmds := cli.Commands{Handlers: make(map[string]func(*cli.State, cli.Command) error)}
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", cli.HandlerAddFeed)
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", cli.HandlerFollow)
	cmds.Register("following", cli.HandlerFollowing)

	// Parse args
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	// Run command
	if err := cmds.Run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
