package models

import (
	"errors"
	"time"

	"github.com/p12s/library-rest-api/logger/pb"
)

const (
	ENTITY_USER = "USER"
	ENTITY_BOOK = "BOOK"

	ACTION_CREATE   = "CREATE"
	ACTION_GET      = "GET"
	ACTION_UPDATE   = "UPDATE"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

var (
	entities = map[string]pb.LoggerRequest_Entities{
		ENTITY_USER: pb.LoggerRequest_USER,
		ENTITY_BOOK: pb.LoggerRequest_BOOK,
	}

	actions = map[string]pb.LoggerRequest_Actions{
		ACTION_CREATE:   pb.LoggerRequest_CREATE,
		ACTION_UPDATE:   pb.LoggerRequest_UPDATE,
		ACTION_GET:      pb.LoggerRequest_GET,
		ACTION_DELETE:   pb.LoggerRequest_DELETE,
		ACTION_REGISTER: pb.LoggerRequest_REGISTER,
		ACTION_LOGIN:    pb.LoggerRequest_LOGIN,
	}
)

// LogItem
type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityId  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

// ToPbEntity - converter
func ToPbEntity(entity string) (pb.LoggerRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

// ToPbAction - converter
func ToPbAction(action string) (pb.LoggerRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
