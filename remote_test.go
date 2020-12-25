package main

import (
	"os"
	"testing"
)

// Author:Boyn
// Date:2020/12/25

func TestLoadRemoteURI(t *testing.T) {
	currentDir, _ := os.Getwd()
	t.Logf("current dir is %s", currentDir)
	uri, err := LoadRemoteURI(currentDir)
	if err != nil {
		t.Error(err)
	}
	if len(uri) == 0 {
		t.Errorf("uri list is empty")
	}
}
