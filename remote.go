package main

// Author:Boyn
// Date:2020/12/25
import (
	"github.com/go-git/go-git/v5"
)

func LoadRemoteURI(path string) ([]string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	remotes, err := r.Remotes()
	if err != nil {
		return nil, err
	}
	uriList := make([]string, 0)
	for _, r := range remotes {
		uriList = append(uriList, r.Config().URLs...)
	}
	return uriList, nil
}
