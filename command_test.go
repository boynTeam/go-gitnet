package main

import "testing"

// Author:Boyn
// Date:2020/12/25

func TestOpen(t *testing.T) {
	err := Open("https://www.google.com")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("open success")
}
