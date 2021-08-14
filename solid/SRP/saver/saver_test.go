package saver

import (
	"testing"

	"gobyexample/solid/SRP"
	"gobyexample/solid/SRP/client"
)

func TestSaveToFile(t *testing.T) {
	c0 := client.NewClient(1, 1)
	e1 := client.NewEncoder(c0)
	type args struct {
		fileName string
		data     SRP.Encoder
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SaveToFile",
			args: args{
				fileName: "client.txt",
				data:     e1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveToFile(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
