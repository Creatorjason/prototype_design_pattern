package main

import "fmt"

type folder struct{
	children []inode
	name string
}

func (fd *folder) print(indentation string){
	fmt.Println(indentation + fd.name)
	for _, i := range fd.children{
		i.print(indentation + indentation)
	}
}

func (fd *folder) clone() inode{
	cloneFolder := &folder{name: fd.name + "_clone"}
	var tempChildren []inode
	for _, i := range fd.children{
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}