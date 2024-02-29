package main

import (
	"math/rand"
)

type RandomizedSet struct {
	data  map[int]int
	slice []int
}

func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		data:  make(map[int]int),
		slice: []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.data[val]; exists {
		return false
	}
	rs.slice = append(rs.slice, val)
	rs.data[val] = len(rs.slice) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, exists := rs.data[val]
	if !exists {
		return false
	}

	lastVal := rs.slice[len(rs.slice)-1]
	rs.slice[index] = lastVal
	rs.data[lastVal] = index

	rs.slice = rs.slice[:len(rs.slice)-1]
	delete(rs.data, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randIndex := rand.Intn(len(rs.slice))
	return rs.slice[randIndex]
}
