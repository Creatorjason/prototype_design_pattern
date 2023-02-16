package main

import "fmt"

type file struct{
	name string
}

func (fl *file) print(indentation string){
	fmt.Println(indentation + fl.name)
}

func (fl *file) clone() inode{
	return &file{name: fl.name + "_clone"}
}