package main

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"reflect"
	"testing"
)

func TestJSONUnmarshal(t *testing.T) {
	var got []codes.Code
	want := []codes.Code{codes.OK, codes.NotFound, codes.Internal, codes.Canceled}
	in := `["OK", "NOT_FOUND", "INTERNAL", "CANCELLED"]`
	err := json.Unmarshal([]byte(in), &got)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("json.Unmarshal(%q, &got) = %v; want <nil>.  got=%v; want %v", in, err, got, want)
	}
}

func TestUnmarshalJSON_NilReceiver(t *testing.T) {
	var got *codes.Code
	in := codes.OK.String()
	if err := got.UnmarshalJSON([]byte(in)); err == nil {
		t.Errorf("got.UnmarshalJSON(%q) = nil; want <non-nil>.  got=%v", in, got)
	}
}

func TestUnmarshalJSON_UnknownInput(t *testing.T) {
	var got codes.Code
	for _, in := range [][]byte{[]byte(""), []byte("xxx"), []byte("codes.Code(17)"), nil} {
		if err := got.UnmarshalJSON([]byte(in)); err == nil {
			t.Errorf("got.UnmarshalJSON(%q) = nil; want <non-nil>.  got=%v", in, got)
		}
	}
}
