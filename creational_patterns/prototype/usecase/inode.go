package usecase

type Inode interface {
    Print(string)
    Clone() Inode
}