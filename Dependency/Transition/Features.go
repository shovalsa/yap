package Transition

import (
	"math"
	"regexp"
	"strconv"
)

// Code copied from float64 version in math/abs.go
func AbsInt(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

func (c *Configuration) GetProperty(property string) (string, bool) {
	if property != "d" {
		return "", false
	}
	if len(c.Stack) != 0 && len(c.Queue) != 0 {
		return AbsInt(c.Nodes[c.Stack[0]].ElementIndex - c.Nodes[c.Queue[0]].ElementIndex)
	}
}

func (c *Configuration) GetSource(source string) *interface{} {
	switch source {
	case "N":
		return &(c.Queue)
	case "S":
		return &(c.Stack)
	}
	return nil
}

func (c *Configuration) GetLocation(currentTarget *interface{}, location []byte) (*HasProperties, bool) {
	switch t := currentTarget.(type) {
	default:
		return nil, false
	case *[]uint16:
		return c.GetLocationNodeStack(t, location)
	case *DepNode:
		return c.GetLocationDepNode(t, location)
	case *DepArc:
		// currentTarget is a DepArc
		// location remainder is discarded
		// (currently no navigation on the arc)
		return currentTarget.(*HasProperties), true
	}
}

func (c *Configuration) GetLocationNodeStack(stack *[]uint16, location string) (*HasProperties, bool) {
	// currentTarget is a slice
	// location "head" must be an offset
	offset, err := strconv.ParseInt(currentLocation, 10, 0)
	if !err {
		panic("Error parsing location string " + location + " ; " + err.Error())
	}
	// if a referenced location cannot exist
	// return an empty result
	if len(t) <= offset {
		return nil, false
	}
	return c.GetLocation(stack[offset], location[len(currentLocation):])
}

// INCOMPLETE!
func (c *Configuration) GetLocationDepNode(node *DepNode, location string) (*HasProperties, bool) {
	// location "head" can be either:
	// - empty (return the currentTarget)
	// - the leftmost/rightmost (l/r) arc
	// - the k-th leftmost/rightmost (lNNN/rNNN) arc
	// - the head (h)
	// - the k-th head (hNNN)
	if len(location) == 0 {
		return node, true
	}
	locationRemainder := location[1:]
	switch location[0] {
	case "h":
		if len(locationRemainder) == 0 {
			if head, exists := c.GetDepNodeHead(node); exists {
				return head.(*HasProperties), true
			} else {
				return nil, false
			}
		}
		return
	case "l":
		return
	case "r":
		return
	}

}

func (c *Configuration) GetDepNodeHead(node *DepNode) (*DepNode, bool) {
	if node.HeadIndex == -1 {
		return nil, false
	}
	if len(c.Nodes) <= node.HeadIndex {
		panic("Node referenced head out of bounds")
	}
	return c.Nodes[node.HeadIndex], true
}
