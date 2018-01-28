package models

import "gopkg.in/mgo.v2/bson"

// Represents a cluster, we use bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Cluster struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	IpAddress  string        `bson:"ipaddress" json:"ipaddress"`
	Description string        `bson:"description" json:"description"`
}
