// Package nodes provides additional functionality for parsley
package nodes

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/scottkgregory/parsley"
)

// DiceNode is used to perform dice rolls
type DiceNode struct {
	left, right parsley.Node
}

var _ parsley.Node = &DiceNode{}

// NewDiceNode generates a new dice node
func NewDiceNode(left, right parsley.Node) *DiceNode {
	return &DiceNode{left, right}
}

// Eval implements parsley.Node.
func (d *DiceNode) Eval(data map[string]any) (any, error) {
	// Evaluate both sides
	leftVal, leftErr := d.left.Eval(data)
	if leftErr != nil {
		return nil, fmt.Errorf("left side error: %w", leftErr)
	}

	rightVal, rightErr := d.right.Eval(data)
	if rightErr != nil {
		return nil, fmt.Errorf("right side error: %w", rightErr)
	}

	lf, err := parsley.ToFloat64(leftVal)
	if err != nil {
		return nil, fmt.Errorf("left side was not valid float: %w", err)
	}

	rf, err := parsley.ToFloat64(rightVal)
	if err != nil {
		return nil, fmt.Errorf("right side was not valid float: %w", err)
	}

	li := int(math.Round(lf))
	ri := int(math.Round(rf))

	ret := 0
	for range li {
		ret += rand.IntN(ri + 1)
	}

	return ret, nil
}

// String implements parsley.Node.
func (d *DiceNode) String() string {
	return "d"
}
