package graph

import (
	"log"
	"reflect"
	"sync"
)

type Weight float64
type Data struct {
	Value interface{}
}

func (d Data) GetType() string {
	return reflect.TypeOf(d.Value).String()
}
func (d Data) GetValue() interface{} {
	return d.Value
}

type Node struct {
	Data Data
	Name string //optional. Should default to a random 5 digit alphanumeric
}

type Graph struct {
	Log   log.Logger
	Edge  map[*Node]map[*Node]Weight
	Nodes []*Node
	mx    *sync.Mutex
}
