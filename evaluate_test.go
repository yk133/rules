package rules

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateBasic(t *testing.T) {
	res, err := Evaluate(`x eq "abc"`, map[string]interface{}{
		"x": "abc",
	})
	require.NoError(t, err)
	require.True(t, res)

	res, err = Evaluate(`x eq abc`, map[string]interface{}{
		"x": "abc",
	})
	require.Error(t, err)
	require.False(t, res)
}

func TestEvaluateWithErrInfo(t *testing.T) {
	res, err := EvaluateWithErrInfo(`x == "abc" or xyz == "abc"  `, map[string]interface{}{
		"x": "abc",
	})
	if err != nil {
		fmt.Printf("err: %+v", err)
	}
	fmt.Println("res:", res)

}
