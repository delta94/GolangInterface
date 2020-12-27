package main

import "log"

//type lam struct {
//	name string
//}
//
//func NewLam(name string) *lam{
//	return &lam{name: name}
//}
//
//type mongoDB struct {
//	name string
//}
//type mySQL struct {
//	name string
//}
//
//func NewMongoDB(name string) *mongoDB{
//	return &mongoDB{name: name}
//}
//
//func NewMySQL(name string) *mySQL{
//	return &mySQL{name: name}
//}
//
//func (db *mongoDB) PrintStorage() string{
//	return fmt.Sprintf("MongoDB -> %v ", db.name)
//}
//func (db *mongoDB) A(){}
//
//func (db *mySQL) PrintStorage() string{
//	return fmt.Sprintf("MySQL -> %v ", db.name)
//}
//
//type PrintStorageInter interface {
//	PrintStorage() string
//}
//
//type repo struct {
//	storageInter PrintStorageInter
//}
//
//func NewRepo(store PrintStorageInter) *repo {
//	return &repo{storageInter: store}
//}
//
//func (re *repo) PrintRepo() string{
//	return fmt.Sprintf("This is data of interface: %v", re.storageInter.PrintStorage())
//}
//
//
//func main() {
//	db := "Data"
//	store := NewMongoDB(db)
//	//lam := NewLam(db)
//	repo := NewRepo(store)
//	//repo := repo{storageInter: store}
//
//
//	fmt.Printf(repo.PrintRepo())
//}
func main() {
	lam := 2
	haha := test(&lam)
	var hehe *int
	hehe = haha
	var hoho int
	hoho = *hehe

	var null *int
	//*null = 10
	unNull := 10
	null = &unNull
	point := null
	var ok int
	ok = *point
	log.Println("*haha ", *haha)
	log.Println("haha ", haha)
	log.Println("&haha ", &haha)
	log.Println("hehe ", hehe)
	log.Println("*hehe ", *hehe)
	log.Println("&hehe ", &hehe)
	log.Println("hoho ", hoho+1)
	log.Println("point ", point)
	log.Println("ok ", ok)
	log.Println("*null ", *null)

	log.Println("")
	log.Println("lam ", lam)
	log.Println("&lam ", &lam)
}

func test(p *int) *int {
	var a int
	a = *p + 1
	log.Println("a ", a)
	log.Println("&a ", &a)
	return &a
}
