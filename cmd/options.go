package main

import (
	cli "github.com/jawher/mow.cli"
)

var (
	db_type *string
)

func setOptions(app *cli.Cli) {

	db_type = app.String(cli.StringOpt{
		Name:   "t dbtype",
		Value:  "mongodb",
		Desc:   "choose what dbms to use",
		EnvVar: "CCA_DB_TYPE DB_TYPE",
	})
	dbaddr = app.String(cli.StringOpt{
		Name:   "a dbaddr",
		Value:  "127.0.0.1:27017",
		Desc:   "db server address",
		EnvVar: "CCA_DB_ADDR DB_ADDR",
	})
	user = app.String(cli.StringOpt{
		Name:   "u user",
		Value:  "admin",
		Desc:   "db server user name",
		EnvVar: "CCA_DB_USER DB_USER",
	})
	passwd = app.String(cli.StringOpt{
		Name:   "p passwd",
		Value:  "admin",
		Desc:   "db server user password",
		EnvVar: "CCA_DB_PWD DB_PWD",
	})
	accessdbname = app.String(cli.StringOpt{
		Name:   "d accessdbname",
		Value:  "admin",
		Desc:   "user source db name (actual for mongodb)",
		EnvVar: "CCA_DB_SOURCE DB_SOURCE",
	})
	dbname = app.String(cli.StringOpt{
		Name:   "n dbname",
		Value:  "accounts",
		Desc:   "working db name to use",
		EnvVar: "CCA_DB_NAME DB_NAME",
	})
	collection = app.String(cli.StringOpt{
		Name:   "c collection",
		Value:  "criminal",
		Desc:   "main collection name to use",
		EnvVar: "CCA_COLLECTION COLLECTION",
	})

}
