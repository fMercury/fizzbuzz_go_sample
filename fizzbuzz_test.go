package main

import (
	"bytes"
	"testing"
	"fmt"
)

var testOutput = `["1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"]`

var testQuery = query{
	str1:  "fizz",
	str2:  "buzz",
	int1:    3,
	int2:    5,
	limit: 15,
}

var testQueries = map[string]query{
	"int1:3 int2:5 limit:15":        testQuery,
	"int1:10 int2:3 limit:100":      query{str1: "fizz", str2: "buzz", int1: 10, int2: 3, limit: 100},
	"int1:1000 int2:1000 limit:500": query{str1: "fizz", str2: "buzz", int1: 100, int2: 100, limit: 500},
}

func testFizzBuzz(t *testing.T, f FizzBuzzFunction, name string) {
	var b bytes.Buffer
	err := f(testQuery, &b)
	if err != nil {
		t.Error(err)
	}
	out := b.String()
	if out != testOutput {
		t.Errorf("Invalid output for %s was\n %s\n expected\n %s", name, out, testOutput)
		t.Fail()
	}
	fmt.Println(" .. .. .. ")
}

func TestFizzBuzz(t *testing.T) {
	testFizzBuzz(t, FizzBuzz, "FizzBuzz")	
	fmt.Println(" . . .")

}
