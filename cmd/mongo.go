package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

var (
	dbaddr       *string
	user         *string
	passwd       *string
	accessdbname *string
	dbname       *string
	collection   *string

	session *mgo.Session
)

func FirstInitMongo() error {
	fmt.Println("Start app.FirstInitMongo")
	var err error
	session, err = mgo.DialWithInfo(
		&mgo.DialInfo{
			Addrs:    []string{*dbaddr},
			Source:   *accessdbname,
			Username: *user,
			Password: *passwd,
			Database: *dbname,
		})
	if err != nil {
		return errors.Wrap(err, "Couldn't access mongo server")
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*dbname).C(*collection)
	err = c.DropCollection()
	if err != nil {
		fmt.Printf("%#v\n", errors.Wrap(err, "Didn't drop collection "+*collection))
	}

	// c = session.DB(dbname).C(collection)

	// index := mgo.Index{
	// 	Key:             []string{"$text:fact", "$text:victim", "$text:accused"},
	// 	DefaultLanguage: "russian",
	// }

	// err = c.EnsureIndex(index)
	// if err != nil {
	// 	fmt.Printf("%#v\n", errors.Wrap(err, "Couldn't create index"))
	// }
	return nil
}

// dbaddr = "172.17.0.2:27017"
// user = "admin"
// passwd = "admin"
// accessdbname = "admin"
// dbname = "accounts"
// collection = "criminal2016"
