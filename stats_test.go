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

func TestNameSplitting(t *testing.T) {
	input := "tama";

	output := ExtractNames(input)
	if !reflect.DeepEqual(output, []string{"tama"}) {
		print(output);
		t.Error("They aren't the same.", output)
	}
}

func TestNameSpacesSplitting(t *testing.T) {
	input := "tama ";

	output := ExtractNames(input)
	if !reflect.DeepEqual(output, []string{"tama"}) {
		print(output);
		t.Error("They aren't the same.", output)
	}
}

func TestTwoNamesSplitting(t *testing.T) {
	input := "tama & seba";

	output := ExtractNames(input)
	if !reflect.DeepEqual(output, []string{"tama", "seba"}) {
		print(output);
		t.Error("They aren't the same.", output)
	}
}

func TestMultipleNamesSplitting(t *testing.T) {
	input := "tama, nuala & seba";

	output := ExtractNames(input)
	if !reflect.DeepEqual(output, []string{"tama", "nuala", "seba"}) {
		print(output);
		t.Error("They aren't the same.", output)
	}
}

func TestNameSurnameSplitting(t *testing.T) {
	input := "Tom Nook";

	output := ExtractNames(input)
	if !reflect.DeepEqual(output, []string{"Tom Nook"}) {
		print(output);
		t.Error("They aren't the same.", output)
	}
}

