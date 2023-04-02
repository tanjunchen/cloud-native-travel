package main

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SetupMongo() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session, nil
}

type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"pop"`
}

// ExecMongo 演示创建和查询
func ExecMongo() error {
	db, err := SetupMongo()
	if err != nil {
		return err
	}

	conn := db.DB("gocookbook").C("example")

	// 我们可以一次性插入多条
	if err := conn.Insert(&State{"Washington", 7062000}, &State{"Oregon", 3970000}); err != nil {
		return err
	}

	var s State
	if err := conn.Find(bson.M{"name": "Washington"}).One(&s); err != nil {
		return err
	}

	if err := conn.DropCollection(); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)
	return nil
}

func testmonggo() {
	if err := ExecMongo(); err != nil {
		panic(err)
	}
}

func main() {
	testmonggo()
}

type Item struct {
	Name  string
	Price int64
}

// Storage是我们的存储接口 将使用Mongo存储实现它
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}

// MongoStorage实现了storage 接口
type MongoStorage struct {
	*mgo.Session
	DB         string
	Collection string
}

// NewMongoStorage 初始化MongoStorage
func NewMongoStorage(connection, db, collection string) (*MongoStorage, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	ms := MongoStorage{
		Session:    session,
		DB:         db,
		Collection: collection,
	}
	return &ms, nil
}

// GetByName 查询mongodb以获取具有正确名称的item
func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	c := m.Session.DB(m.DB).C(m.Collection)
	var i Item
	if err := c.Find(bson.M{"name": name}).One(&i); err != nil {
		return nil, err
	}
	return &i, nil
}

// Put 添加一个item到mongo 中
func (m *MongoStorage) Put(ctx context.Context, i *Item) error {
	c := m.Session.DB(m.DB).C(m.Collection)
	return c.Insert(i)
}

// PerformOperations 演示创建并存入一个item然后对其进行查询查询
func PerformOperations(s Storage) error {
	ctx := context.Background()
	i := Item{Name: "candles", Price: 100}
	if err := s.Put(ctx, &i); err != nil {
		return err
	}

	candles, err := s.GetByName(ctx, "candles")
	if err != nil {
		return err
	}
	fmt.Printf("Result: %#v\n", candles)
	return nil
}

func test4() {
	if err := ExecMongoStorage(); err != nil {
		panic(err)
	}
}

// Exec 初始化存储 storage然后使用存储接口执行操作
func ExecMongoStorage() error {
	m, err := NewMongoStorage("localhost", "gocookbook", "items")
	if err != nil {
		return err
	}
	if err := PerformOperations(m); err != nil {
		return err
	}

	if err := m.Session.DB(m.DB).C(m.Collection).DropCollection(); err != nil {
		return err
	}
	return nil
}
