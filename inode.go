package main

// `In the context of golang 
// let’s try to understand it with an example of os file system. 
// The os file system has files and folders and folders itself contain files and folders. 
// Each file and folder can be represented by an inode interface. 
// inode interface also has the clone() function.`

type inode interface{
	print(string)
	clone() inode
}

