package repo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

//GetClient gets a mongo client
func GetClient() *mongo.Client {

	c, err := readConf()

	clientOptions := options.Client().ApplyURI(c.Conf.Atlas)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Bad URI: ", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// NOTE: Fields in private structure MUST BE EXPORTED; upper case field names required here
type myData struct {
	Conf struct {
		Atlas string `yaml:"atlas"`
	}
}

func readConf() (*myData, error) {
	filename := "config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	c := &myData{}
	err = yaml.Unmarshal(buf, c) // Careful about structure name mapping
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
