package onstrings

import (
	"testing"
)

var inputsConvertToPalindrome = []string{"ivjviw", "phmjjmap", "ivjwvi", "epyyevdadveyype", "abcba", "abecbea"}
var outputsConvertToPalindrome = []int{1, 0, 1, 1, 1, 0}

func TestConvToPalindrome(t *testing.T) {

	for i, input := range inputsConvertToPalindrome {
		res := convToPalindrome(input)
		if res != outputsConvertToPalindrome[i] {
			t.Error("expected output", outputsConvertToPalindrome[i], "got", res, "for input", input)
		}
	}
}
