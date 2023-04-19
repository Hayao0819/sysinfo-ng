package ioreg

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	"howett.net/plist"
)

type I *map[string]any

func ioreg(args ...string)(I, error){
	c := exec.Command("ioreg", append(args, "-a")...)
	b, err :=  c.Output()
	if err !=nil{
		return nil, err
	}
	decoder := plist.NewDecoder(bytes.NewReader(b))
	i:= map[string]any{}
	err = decoder.Decode(&i)
	if err !=nil{
		return nil, err
	}
	return &i, nil
}

func GetAll()(I, error){
	i , _:= ioreg()
	f, _ := os.Create("ioreg.json")

	j, _ :=json.Marshal(*i)
	f.Write(j)
	return i, nil
}


