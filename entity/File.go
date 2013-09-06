package entity

type File struct {
	Id         uint
	FolderId   uint
	UserId     uint
	Name       string
	Stuffix    string
	Size       float32
	CreateDate uint
	ModifyDate uint
}
