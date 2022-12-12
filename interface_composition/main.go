package main

//the purpose of this script is to illustrate creating an interface
//that is composed of two other interfaces
//in this case, the Storage interface implements the methods
//of a UserStore and an ItemStore

import (
	"encoding/json"
	"fmt"
)

//create types
type User struct {
	Name string
	Age  int
}

type UserStore interface {
	AddUser(users []*User, user *User) []*User
}

type ustore struct{}

type Item struct {
	Name  string
	Usage string
}

type ItemStore interface {
	AddItem(items []*Item, item *Item) []*Item
}

type istore struct{}

type Storage interface {
	UserStore
	ItemStore
}

type myStore struct {
	UserStore
	ItemStore
}

//constructors
func NewUserStore() UserStore {
	return &ustore{}
}

func NewItemStore() ItemStore {
	return &istore{}
}

func NewStorage() Storage {
	u := NewUserStore()
	i := NewItemStore()

	s := &myStore{
		UserStore: u,
		ItemStore: i,
	}

	return s
}

//methods
func (is istore) AddItem(items []*Item, item *Item) []*Item {

	items = append(items, item)

	return items
}

func (us ustore) AddUser(users []*User, user *User) []*User {

	users = append(users, user)

	return users
}

func main() {

	s := NewStorage()

	var items []*Item
	var users []*User

	u := &User{
		Name: "Eric",
		Age:  34,
	}

	i := &Item{
		Name:  "Phone",
		Usage: "texting",
	}

	items = s.AddItem(items, i)
	users = s.AddUser(users, u)

	//embed to json
	ji, _ := json.Marshal(items)
	ju, _ := json.Marshal(users)

	fmt.Println(string(ji))
	fmt.Println(string(ju))
}
