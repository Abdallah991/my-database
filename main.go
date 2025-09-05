package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	fmt.Println("this is test")
}

// updating files in place
// Using the file system as a KV
// ! doesnt handle the concurrent writer scenarios
func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(data)
	return err
}

// write file in data with atomicity
// will count for concurrent writer scenarios
// ! this still dont solve for power loss scenarios
func SaveData2(path string, data []byte) error {
	// create a unique file name with format of //?path.tmp.rand
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	// open file with tmp name
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()
	// 1- write data to file
	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}
	// 2- sync data to files
	err = fp.Sync()
	if err != nil {
		return err
	}
	// 3- rename to path when its successful
	err = os.Rename(tmp, path)
	return err
}
