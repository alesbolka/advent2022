package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type argFunction func(old int) int
type operatorFunction func(aa, bb int) int

type operation struct {
	arg1    argFunction
	arg2    argFunction
	operand operatorFunction
}

func ParseOperation(input string) (*operation, error) {
	opArgs := strings.Split(input, " ")
	if len(opArgs) != 3 {
		return nil, fmt.Errorf("cannot parse operation input '%s'", input)
	}

	switch opArgs[1] {
	case "+":
		return newSum(opArgs[0], opArgs[2])
	case "*":
		return newMultiply(opArgs[0], opArgs[2])
	case "-":
		return newSubtract(opArgs[0], opArgs[2])
	case "/":
		return newDivision(opArgs[0], opArgs[2])
	default:
		return nil, fmt.Errorf("unhandled operand '%s'", opArgs[1])
	}
}

func newSum(arg1, arg2 string) (*operation, error) {
	op, err := newOp(arg1, arg2)
	if err != nil {
		return nil, err
	}

	op.operand = func(aa, bb int) int { return aa + bb }

	return op, nil
}

func newSubtract(arg1, arg2 string) (*operation, error) {
	op, err := newOp(arg1, arg2)
	if err != nil {
		return nil, err
	}

	op.operand = func(aa, bb int) int { return aa - bb }

	return op, nil
}

func newMultiply(arg1, arg2 string) (*operation, error) {
	op, err := newOp(arg1, arg2)
	if err != nil {
		return nil, err
	}

	op.operand = func(aa, bb int) int { return aa * bb }

	return op, nil
}

func newDivision(arg1, arg2 string) (*operation, error) {
	op, err := newOp(arg1, arg2)
	if err != nil {
		return nil, err
	}

	op.operand = func(aa, bb int) int { return aa / bb }

	return op, nil
}

func newOp(arg1, arg2 string) (*operation, error) {
	op := &operation{
		arg1: func(old int) int { return old },
		arg2: func(old int) int { return old },
	}
	if arg1 != "old" {
		val, err := strconv.Atoi(arg1)
		if err != nil {
			return nil, err
		}
		op.arg1 = func(old int) int { return val }
	}
	if arg2 != "old" {
		val, err := strconv.Atoi(arg2)
		if err != nil {
			return nil, err
		}
		op.arg1 = func(old int) int { return val }
	}

	return op, nil
}
