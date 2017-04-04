package main

import (
	"sync"
)

type Canvas struct {
	Size    int
	Data    []byte
	RWMutex *sync.RWMutex
}

func NewCanvas(size int) *Canvas {
	bi := &Canvas{
		Size:    size,
		Data:    make([]byte, size*size*3),
		RWMutex: &sync.RWMutex{},
	}

	for i, _ := range bi.Data {
		bi.Data[i] = 0
	}

	return bi
}

func (c *Canvas) Set(x, y int, r, g, b byte) error {
	j := (y*c.Size + x) * 3
	c.Data[j] = r
	c.Data[j+1] = g
	c.Data[j+2] = b
	return nil
}
