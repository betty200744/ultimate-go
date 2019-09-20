package main

import "fmt"

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

func main() {
	duplicateList := []Bootstrap{
		{
			Address: "1",
			PeerID:  "1",
		},
		{
			Address: "2",
			PeerID:  "2",
		},
		{
			Address: "1",
			PeerID:  "1",
		},
		{
			Address: "2",
			PeerID:  "3",
		},
	}
	uniqueList := []Bootstrap{}

	for _, value := range duplicateList {
		if !dupInSlice(value, uniqueList) {
			uniqueList = append(uniqueList, value)
		}
	}

	fmt.Println(uniqueList)

}
