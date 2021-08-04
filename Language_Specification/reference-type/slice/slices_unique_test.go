package slice

import (
	"fmt"
	"testing"

	"github.com/lib/pq"
)

func TestUniqueStringArray(t *testing.T) {
	type args struct {
		stringSlice *pq.StringArray
	}
	tests := []struct {
		name string
		args args
		want *pq.StringArray
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringSlice := []string{"1", "5", "3", "6", "9", "9", "4", "2", "3", "1", "5"}
			stringSlice = append(stringSlice, stringSlice...)
			fmt.Println(stringSlice)
			uniqueSlice := unique(stringSlice)
			fmt.Println(uniqueSlice)

			sa := &pq.StringArray{"a", "b"}

			*sa = append(*sa, *sa...)
			fmt.Println(sa)
			su := UniqueStringArray(sa)
			fmt.Println(su)

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
		})
	}
}
