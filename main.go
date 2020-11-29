package main

import "fmt"

type lam struct {
	name string
}

func NewLam(name string) *lam{
	return &lam{name: name}
}

type mongoDB struct {
	name string
}
type mySQL struct {
	name string
}

func NewMongoDB(name string) *mongoDB{
	return &mongoDB{name: name}
}

func NewMySQL(name string) *mySQL{
	return &mySQL{name: name}
}

func (db *mongoDB) PrintStorage() string{
	return fmt.Sprintf("MongoDB -> %v ", db.name)
}
func (db *mongoDB) A(){}

func (db *mySQL) PrintStorage() string{
	return fmt.Sprintf("MySQL -> %v ", db.name)
}

type PrintStorageInter interface {
	PrintStorage() string
}

type repo struct {
	storageInter PrintStorageInter
}

func NewRepo(store PrintStorageInter) *repo {
	return &repo{storageInter: store}
}

func (re *repo) PrintRepo() string{
	return fmt.Sprintf("This is data of interface: %v", re.storageInter.PrintStorage())
}


func main() {
	db := "Data"
	store := NewMongoDB(db)
	//lam := NewLam(db)
	//repo := NewRepo(store)
	repo := repo{storageInter: store}


	fmt.Printf(repo.PrintRepo())
}
