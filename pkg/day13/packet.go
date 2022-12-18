package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Packet struct {
	isList  bool
	intVal  int
	listVal []*Packet
	parent  *Packet
}

func NewPacket(input string) *Packet {
	val := Packet{
		isList: input[0] == '[' && input[len(input)-1] == ']',
	}

	numStr := ""
	currentNode := &val

	for ii := 1; ii < len(input)-1; ii++ {
		if input[ii] >= '0' && input[ii] <= '9' {
			numStr += string(input[ii])
			continue
		}

		if input[ii] == ',' || input[ii] == ']' {
			if len(numStr) > 0 {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatalf("Failed processing number from '%s' at index %d", input, ii)
				}
				numStr = ""
				if !currentNode.isList {
					log.Fatalf("trying to append to nonlist via input '%s' at index %d", input, ii)
				}

				singleVal := &Packet{
					intVal: num,
				}
				singleVal.listVal = []*Packet{singleVal}
				currentNode.listVal = append(currentNode.listVal, singleVal)
			}

			if input[ii] == ']' {
				currentNode = currentNode.parent
				if currentNode == nil {
					log.Fatalf("Have no parent node in input '%s' at position %d", input, ii)
				}
			}
			continue
		}

		if input[ii] == '[' {
			newNode := Packet{
				isList: true,
				parent: currentNode,
			}
			currentNode.listVal = append(currentNode.listVal, &newNode)
			currentNode = &newNode
			continue
		}
	}

	if len(numStr) > 0 {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("Failed processing number from '%s' at the end", input)
		}
		numStr = ""
		if !currentNode.isList {
			log.Fatalf("trying to append to nonlist via input '%s' at the end", input)
		}
		singleVal := &Packet{
			intVal: num,
		}
		singleVal.listVal = []*Packet{singleVal}
		currentNode.listVal = append(currentNode.listVal, singleVal)
	}

	return &val
}

func (left *Packet) Compare(right *Packet) int {
	if !left.isList && !right.isList {
		if left.intVal < right.intVal {
			return 1
		} else if left.intVal == right.intVal {
			return 0
		}
		return -1
	}

	// Since the listVal of an intiger type value will contain a reference to itself in its list, we can just treat it the same
	for ii, rightVal := range right.listVal {
		if ii >= len(left.listVal) {
			// left list is shorter, order is OK
			return 1
		}

		leftVal := left.listVal[ii]
		cmpVal := leftVal.Compare(rightVal)

		if cmpVal == 0 {
			// These two values could not determine order
			continue
		}

		// the lists have a conclusion
		return cmpVal
	}

	if len(left.listVal) == len(right.listVal) {
		return 0
	}
	// Left side is longer than right side with no other decisions
	return -1
}

func (val Packet) String() string {
	if !val.isList {
		return fmt.Sprintf("%d", val.intVal)
	}
	res := []string{}
	for _, innerVal := range val.listVal {
		res = append(res, innerVal.String())
	}
	return "[" + strings.Join(res, ",") + "]"
}
