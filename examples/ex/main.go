package main

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func newBook(name string, pages int) *book {
	return &book{
		Name:  name,
		Pages: pages,
	}
}
func main() {
	mgm.SetCtxTimeout(10 * time.Second)
	uri := "mongodb://172.16.1.248/test?w=majority&tz_aware=True"
	mongoClient, err := mongo.Connect(mgm.Ctx(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err = mongoClient.Ping(mgm.Ctx(), readpref.Primary()); err != nil {
		panic(err)
	}
	db := mongoClient.Database("terrace")
	mgm.SetMongoDatabase(db)
	m := newBook("boo", 12)
	mgm.Coll(m).Create(m)

}
