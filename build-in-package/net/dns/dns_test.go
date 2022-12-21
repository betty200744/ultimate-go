package dns

import (
	"net"
	"testing"
)

func TestLookupAddr(t *testing.T) {
	// dig ptr  -x  8.8.8.8
	// dig ptr  -x  114.114.114.114
	adds := []string{"8.8.8.8", "114.114.114.114"}
	for _, add := range adds {
		names, err := net.LookupAddr(add)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(names)
	}
}

func TestLookupCNAME(t *testing.T) {
	// dig cname www.iana.org
	names := []string{"www.iana.org", "www.iana.org.", "www.google.com"}
	for _, name := range names {
		cname, _ := net.LookupCNAME(name)
		t.Log(cname)
	}
}

func TestLookupHost(t *testing.T) {
	// dig google.com
	names := []string{"google.com"}
	for _, name := range names {
		addrs, _ := net.LookupHost(name)
		t.Log(addrs)
	}
}
