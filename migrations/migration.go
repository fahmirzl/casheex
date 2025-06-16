package migrations

import (
	"casheex/configs"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func DBMigrate() {
    migrations := &migrate.EmbedFileSystemMigrationSource{
       FileSystem: dbMigrations,
       Root:       "sql_migrations",
    }

    n, errs := migrate.Exec(configs.DB, "mysql", migrations, migrate.Up)
    if errs != nil {
       panic(errs)
    }

    fmt.Println("Migration success, applied", n, "migrations!")
}