package database

import (
	"log"

	. "github.com/djfoley01/test-monitor/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CDatabase struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "clusters"
)

// Establish a connection to database
func (m *CDatabase) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of clusters
func (m *CDatabase) FindAll() ([]Cluster, error) {
	var clusters []Cluster
	err := db.C(COLLECTION).Find(bson.M{}).All(&clusters)
	return clusters, err
}

// Find a cluster by its id
func (m *CDatabase) FindById(id string) (Cluster, error) {
	var cluster Cluster
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&cluster)
	return cluster, err
}

// Insert a cluster into database
func (m *CDatabase) Insert(cluster Cluster) error {
	err := db.C(COLLECTION).Insert(&cluster)
	return err
}

// Delete an existing cluster
func (m *CDatabase) Delete(cluster Cluster) error {
	err := db.C(COLLECTION).Remove(&cluster)
	return err
}

// Update an existing cluster
func (m *CDatabase) Update(cluster Cluster) error {
	err := db.C(COLLECTION).UpdateId(cluster.ID, &cluster)
	return err
}
