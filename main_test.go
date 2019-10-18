package main

import "testing"

func TestFizzBuzz(t *testing.T){
	fb := FizzBuzz(1)

	if fb != "1"  {
		t.Errorf("should be 1, %s", fb)
	}
}



