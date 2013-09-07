package entity

type File struct {
	Id         int
	FolderId   int
	UserId     int
	Name       string
	Stuffix    string
	Size       int64
	CreateDate int64
	ModifyDate int64
}
