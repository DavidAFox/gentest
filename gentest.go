//A Test is made with a slice of actors and optional weights.  The test will choose one of the actors and then run its Do method and store the returned string.  The actors Do methods should generally return a string identifying them and what they did.
package gentest

import (
	"math/rand"
	"time"
	"fmt"
)

//Actor has a Do method that returns a string describing what it did.
type Actor interface {
	Do(interface{})string
}

//Object for running the test.
type Test struct {
	name string
	actors []Actor
	weights []int
	shared interface{}
	commands int
	actions []string
}

//New returns a new Test using actors and shared with equal weights.
func New(name string, actors []Actor, shared interface{}) *Test {
	t := new(Test)
	t.actors = actors
	t.shared = shared
	t.commands = 10000
	t.name = name
	t.actions = make([]string, 0, t.commands)
	weights := make([]int,len(t.actors), len(t.actors))
	for i := range actors {
		weights[i] = 1
	}
	t.weights = weights
	return t
}

//NewWithWeights creats a new test object with the actors and weights.
func NewWithWeights(name string, actors []Actor, shared interface{}, weights []int) *Test{
	if len(actors) != len(weights) {
		panic("weights != actors")
	}
	t := new(Test)
	t.actors = actors
	t.weights = weights
	t.shared = shared
	t.commands = 10000
	t.actions = make([]string, 0, t.commands)
	return t
}

//SetCommands sets the number of commands for the test to execute.  The default is 10000.
func (test *Test) SetCommands(c int){
	test.commands = c
}

//Start runs the test choosing random actors according to their weights and records the string returned in actions.
func (test *Test) Run() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0;i < test.commands; i++ {
		test.actions = append(test.actions, test.Do(test.shared))
	}
}
//Do is a method that executes a random actor and also makes the test itself satisfy the actor interface.
func (test *Test) Do(shared interface{}) string {
	return fmt.Sprintf("%v: %v", test.name, test.actors[choose(test.weights)].Do(shared))
}
//Actions returns the list of actions done.
func (test *Test) Actions() []string{
	return test.actions
}
//choose returns the index of the weight randomly choosen based on the weights.
func choose(weights []int) int {
	var total int = 0
	var done int = 0
	for i := range weights {
		total = total + weights[i]
	}
	rnumber := rand.Int() % total
	for i := range weights {
		if rnumber < weights[i] + done {
			return i
		} else {
			done = done + weights[i]
		}
	}
	return -1 //none were chosen, probably an error
}
