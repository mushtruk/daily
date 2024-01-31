package main

import (
	"math/rand"
)

type RandomizedSet struct {
	data  map[int]int
	slice []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:  make(map[int]int),
		slice: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.data[val]; exists {
		return false
	}
	this.slice = append(this.slice, val)
	this.data[val] = len(this.slice) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.data[val]
	if !exists {
		return false
	}

	lastVal := this.slice[len(this.slice)-1]
	this.slice[index] = lastVal
	this.data[lastVal] = index

	this.slice = this.slice[:len(this.slice)-1]
	delete(this.data, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := rand.Intn(len(this.slice))
	return this.slice[randIndex]
}
