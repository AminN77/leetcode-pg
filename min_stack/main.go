package main

import (
	"log"
)

func main() {
	log.Println(Constructor())
}

type MinStack struct {
	data    []int
	minimum int
	top     int
}

func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minimum: 0,
		top:     0,
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.top = val
	if val < this.minimum {
		this.minimum = val
	}

}

func (this *MinStack) Pop() {
	tempTop := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	this.top = this.data[len(this.data)-1]
	if this.minimum == tempTop {
		// update min
		tempMin := this.data[0]
		for i := 0; i < len(this.data); i++ {
			if this.data[i] < tempMin {
				tempMin = this.data[i]
			}
		}

		this.minimum = tempMin
	}
}

func (this *MinStack) Top() int {
	return this.top
}

func (this *MinStack) GetMin() int {
	return this.minimum
}
