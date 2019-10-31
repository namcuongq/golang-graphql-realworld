package container 

import (
	"gopkg.in/mgo.v2"
    "github.com/BurntSushi/toml"
	"time"
)

type Config struct{
    Listen string

    MongoServer []string
	MongoUser string
	MongoPassword string
	MongoAuthDB string
    MongoDatabase string
}

type Container struct{
    Config      Config
    MongoClient *mgo.Database
}

var container  = new(Container)

func Get() *Container{
    return container
}

func Init(pathConfig string) error {
    if _, err := toml.DecodeFile(pathConfig, &container.Config); err != nil {
		return err
	}

    err := container.loadMongo()
    if err != nil{
        return err
    }

    return nil

}

func (self *Container) loadMongo() error {
	var err error
    info := &mgo.DialInfo{
		Addrs:    container.Config.MongoServer,
		Timeout:  30 * time.Second,
		Username: container.Config.MongoUser,
		Password: container.Config.MongoPassword,
		Database: container.Config.MongoAuthDB,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return err
	}

	container.MongoClient = session.DB(container.Config.MongoDatabase)
	container.MongoClient.Session.SetSocketTimeout(30 * time.Second)
	return nil
}