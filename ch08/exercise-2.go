package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type record struct {
	ID   int
	Name string
	Addr string
}

type list []record
type relation map[int]record

func (obj list) saveTo(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	if err := encoder.Encode(obj); err != nil {
		return err
	}
	return nil
}

func (obj *list) loadFrom(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

func (obj relation) saveTo(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	if err := encoder.Encode(obj); err != nil {
		return err
	}
	return nil
}

func (obj *relation) loadFrom(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

func main() {
	arr := list{
		record{0, "A", "A_addr"},
		record{1, "B", "B_addr"},
		record{2, "C", "C_addr"},
	}
	err := arr.saveTo("exercise-2-list.gob")
	if err != nil {
		fmt.Println(err)
	}

	var emptyList list
	err = emptyList.loadFrom("exercise-2-list.gob")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("list:", arr)
	fmt.Println("emptyList:", emptyList)

	dict := relation{
		13: record{13, "D", "D_addr"},
		15: record{15, "F", "F_addr"},
		96: record{96, "60", "60_addr"},
	}
	err = dict.saveTo("exercise-2-dict.gob")
	if err != nil {
		fmt.Println(err)
	}

	var emptyDict relation
	err = emptyDict.loadFrom("exercise-2-dict.gob")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dict:", dict)
	fmt.Println("emptyDict", emptyDict)
}
