package main

import (
	ch "github.com/saromanov/consistenthash"
)
func main(){
	c := ch.New(nil)
	c.Add("1")
}