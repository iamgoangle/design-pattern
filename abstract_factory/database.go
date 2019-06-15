package main

import "errors"

// abstract
type Database interface {
	Connect() error
}

// product
type MongoDB struct {
	conection string
}

func (m *MongoDB) Connect() error {
	return nil
}

type HBase struct {
	conection string
}

func (h *HBase) Connect() error {
	return errors.New("unable to connect hbase")
}

type Databases struct {
	*MongoDB
	*HBase
}

// func GetFactory()

func DBClient(db string) Database {
	switch db {
	case "mongo":
		return &MongoDB{}
	case "hbase":
		return &HBase{}
	default:
		return nil
	}
}

func main() {

}
