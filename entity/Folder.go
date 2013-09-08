package entity

type Folder struct {
	Id         int
	ParentId   int
	UserId     int
	Name       string
	CreateDate int64
}
