package main

import (
	"testing"
	"reflect"
)

func TestSummarising(t *testing.T) {
	input := []string{"tama","tama","seba"};

	output := summarise(input)

	if !reflect.DeepEqual(output, map[string]int{"tama":2, "seba":1}) {
		t.Error("They aren't the same.", output)
	}
}
