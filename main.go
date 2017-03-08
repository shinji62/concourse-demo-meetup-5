package main

import (
	"fmt"

	_ "github.com/mattes/migrate/driver/mysql"
	"github.com/mattes/migrate/migrate"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	migration = kingpin.Flag("migration", "if true migration db will happen").Default("false").Bool()
	mysqlDSN  = kingpin.Flag("database-url", "mysql dns").Default("mysql://root:root@tcp(localhost:3306)/meetup").String()
)

func Migrate(mysqlDSN string) {
	fmt.Println("Migration triggered")
	allErrors, ok := migrate.UpSync(mysqlDSN, "./migration")
	if !ok {
		fmt.Println("Error During Migration", allErrors)
	}

}

func main() {

	kingpin.Parse()
	if *migration {
		Migrate(*mysqlDSN)
	}

}
