package slice

import (
	"github.com/lib/pq"
)

type Bootstrap struct {
	Address string
	PeerID  string
}

func dupInSlice(item Bootstrap, list []Bootstrap) bool {
	for _, v := range list {
		if v.PeerID == item.PeerID && v.Address == item.Address {
			return true
		}
	}
	return false
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueStringArray(stringSlice *pq.StringArray) *pq.StringArray {
	keys := make(map[string]bool)
	list := new(pq.StringArray)
	for _, entry := range *stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			*list = append(*list, entry)
		}
	}
	return list
}
