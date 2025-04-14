package nodes

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley"
	"github.com/scottkgregory/rolo/internal/assert"
)

func Test(tt *testing.T) {
	testCases := []struct {
		left           int
		right          int
		expectedMin    int
		expectedMax    int
		expectedString string
	}{
		{
			left:           1,
			right:          4,
			expectedMin:    0,
			expectedMax:    4,
			expectedString: "d",
		},
		{
			left:           2,
			right:          20,
			expectedMin:    0,
			expectedMax:    40,
			expectedString: "d",
		},
	}
	for _, tc := range testCases {
		tt.Run(fmt.Sprintf("%dd%d", tc.left, tc.right), func(t *testing.T) {
			for range 500 {
				n := NewDiceNode(&mockNode{1, "1"}, &mockNode{4, "4"})
				actual, err := n.Eval(nil)
				assert.Nil(t, err)

				af, err := parsley.ToFloat64(actual)
				assert.Nil(t, err)

				ai := int(af)
				assert.Nil(t, err)
				assert.Equal(t, true, ai >= tc.expectedMin && ai <= tc.expectedMax)

				assert.Equal(t, n.String(), tc.expectedString)
			}
		})
	}
}

type mockNode struct {
	ret any
	str string
}

var _ parsley.Node = &mockNode{}

// Eval implements parsley.Node.
func (m *mockNode) Eval(data map[string]any) (any, error) {
	return m.ret, nil
}

// String implements parsley.Node.
func (m *mockNode) String() string {
	return m.str
}
