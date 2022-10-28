package onstrings

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputsConvertNumToWords = []string{"1", "10", "25", "01", "88", "91", "10", "0", "5002", "50002", "100", "1234567", "123456789", "100000000", "1000000000"}
var outputsConvertNumToWords = []string{"one", "ten", "twenty-five", "one", "eighty-eight", "ninety-one", "ten", "zero", "five-thousand-and-two", "fifty-thousand-and-two", "one-hundred", "twelve-lakh-thirty-four-thousand-five-hundred-and-sixty-seven", "twelve-crore-thirty-four-lakh-fifty-six-thousand-seven-hundred-and-eighty-nine", "ten-crore", "one-arab"}

func TestGetNumberInWords(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	assert := assert.New(t)
	for i, input := range inputsConvertNumToWords {
		assert.Equal(outputsConvertNumToWords[i], getNumberInWords(input), fmt.Sprintf("the number in words is not as expected for %s", input))
	}
}
