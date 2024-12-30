package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/SauravNaruka/gator/internal/config"
	"github.com/SauravNaruka/gator/internal/database"
	"github.com/SauravNaruka/gator/internal/gatorapi"
	_ "github.com/lib/pq"
)

type state struct {
	db             *database.Queries
	cfg            *config.Config
	gatorapiClient *gatorapi.Client
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	gatorapiClient := gatorapi.NewClient(5 * time.Second)

	programState := &state{
		db:             dbQueries,
		cfg:            &cfg,
		gatorapiClient: gatorapiClient,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerListFeedFollows)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
