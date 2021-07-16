package main

import (
	"io"
	"strconv"
)

type query struct {
	str1  string
	str2  string
	int1    int
	int2    int
	limit int
}

type output []string

// FizzBuzzFunction writes a json encoded list of string that correspond to a fizz buzz list
type FizzBuzzFunction func(query, io.Writer) error

// FizzBuzz generates values with a switch case and all modulo in an explicit way
func FizzBuzz(q query, w io.Writer) error {
	fizz := []byte(q.str1)
	buzz := []byte(q.str2)
	fizzbuzz := []byte(q.str1 + q.str2)
	_, err := w.Write([]byte("["))
	if err != nil {
		return err
	}
	for i := 1; i <= q.limit; i++ {
		if i > 1 {
			_, err := w.Write([]byte(", "))
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("\""))
		if err != nil {
			return err
		}
		switch {
		case i%q.int1 == 0 && i%q.int2 == 0:
			_, err := w.Write(fizzbuzz)
			if err != nil {
				return err
			}
		case i%q.int1 == 0:
			_, err := w.Write(fizz)
			if err != nil {
				return err
			}
		case i%q.int2 == 0:
			_, err := w.Write(buzz)
			if err != nil {
				return err
			}
		default:
			_, err := w.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("\""))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("]"))
	return err
}



