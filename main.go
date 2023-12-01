// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"personal-budget/util"

// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"

// 	"github.com/golang-migrate/migrate/v4"
// 	_ "github.com/golang-migrate/migrate/v4/database/postgres"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// 	_ "github.com/lib/pq"
// )

// // const (
// // 	dbDriver      = "postgres"
// // 	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// // 	serverAddress = "0.0.0.0:8080"
// // )

// func main() {

// 	config, err := util.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal().Msg("Cannot connect to db:")
// 	}
// 	if config.Environment == "development" {
// 		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

// 	}
// 	conn, err := sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal().Msg("Cannot connect to db:")
// 	}

// 	/// run db migration
// 	runDBMigration(config.MigrationURL, config.DBSource)
// 	fmt.Println(conn)
// 	// store := db.NewStore(conn)

// 	// redisOpt := asynq.RedisClientOpt{
// 	// 	Addr: config.RedisAddress,
// 	// }
// 	// // taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
// 	// // go runTaskProcessor(config, redisOpt, store)
// 	// runGinServer

// }

// func runDBMigration(migrationURL string, dbSource string) {
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		log.Fatal().Msgf("cannot create new migrate instance: %w", err)
// 	}
// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal().Msg("failed to run migration up:")
// 	}
// 	log.Info().Msg("db migrated successfully")
// }

// // func runGinServer(config util.Config) {
// // 	server, err := api.NewServer(config)
// // 	if err != nil {
// // 		log.Fatal().Msg("cannot create server")
// // 	}
// // 	err = server.Start(config.HTTPServerAddress)
// // 	if err != nil {
// // 		log.Fatal().Msg("cannot start server")
// // 	}
// // }

package main

import (
	"fmt"
	"regexp"
)

func extractNumerals(input string) string {
	// Create a regular expression to match numerals
	re := regexp.MustCompile("[0-9]+")

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)

	// Concatenate the matched numerals into a single string
	result := ""
	for _, match := range matches {
		result += match
	}

	return result
}

func main() {
	inputString := "dqfournine5four2jmlqcgv"
	numerals := extractNumerals(inputString)

	fmt.Println("Numerals:", numerals)
}
