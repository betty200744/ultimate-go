package main

type ClientInfo struct {
	Pid int64
}

type ClientInfoSlice []*ClientInfo

func (p ClientInfoSlice) Len() int {
	return len(p)
}
func (p ClientInfoSlice) Less(i, j int) bool {
	if p[i] != nil && p[j] != nil {
		return p[i].Pid < p[j].Pid
	}
	return false
}
func (p ClientInfoSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
