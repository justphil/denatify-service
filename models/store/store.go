package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/justphil/denatify-service/models"
)

type Store interface {
	CreateUser(u *models.User) (string, error)
	GetUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(u *models.User) error
	DeleteUserById(id string) error
	Close()
}

const (
	DB        = "nat-busters"
	USERS_COL = "users"
)

type MongoStore struct {
	sess *mgo.Session
}

func NewMongoStore(url string) *MongoStore {
	sess, err := mgo.Dial(url)

	// TODO: add indexes

	if err != nil {
		panic(err)
	}

	return &MongoStore{sess}
}

func (mongo *MongoStore) query(db string, col string, op func(sess *mgo.Collection)) {
	sess := mongo.sess.Copy()
	defer sess.Close()
	op(sess.DB(db).C(col))
}

func (mongo *MongoStore) CreateUser(u *models.User) (string, error) {
	var userId string
	var err error
	mongo.query(DB, USERS_COL, func(c *mgo.Collection) {
		u.Id = bson.NewObjectId()
		err = c.Insert(u)
		if err == nil {
			userId = u.Id.Hex()
		}
	})

	return userId, err
}

func (mongo *MongoStore) GetUsers() ([]*models.User, error) {
	var users []*models.User
	var err error

	mongo.query(DB, USERS_COL, func(c *mgo.Collection) {
		err = c.Find(nil).All(&users)
	})

	return users, err
}

func (mongo *MongoStore) GetUserById(id string) (*models.User, error) {
	user := &models.User{}
	var err error

	mongo.query(DB, USERS_COL, func(c *mgo.Collection) {
		err = c.FindId(bson.ObjectIdHex(id)).One(user)
	})

	return user, err
}

func (mongo *MongoStore) UpdateUser(u *models.User) error {
	var err error

	mongo.query(DB, USERS_COL, func(c *mgo.Collection) {
		err = c.UpdateId(u.Id, u)
	})

	return err
}

func (mongo *MongoStore) DeleteUserById(id string) error {
	var err error

	mongo.query(DB, USERS_COL, func(c *mgo.Collection) {
		err = c.RemoveId(bson.ObjectIdHex(id))
	})

	return err
}

func (mongo *MongoStore) Close() {
	mongo.sess.Close()
}
