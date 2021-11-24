package main

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

type Incidents struct {
	Id        uuid.UUID
	ChannelId string
	Title     string
	Desc      string
	Sev       string
	Creator   string
	Deleted   bool
}

var schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"incidents": &memdb.TableSchema{
			Name: "incidents",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "ChannelId"},
				},
				"title": &memdb.IndexSchema{
					Name:    "title",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Title"},
				},
			},
		},
	},
}
