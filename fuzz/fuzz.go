// fuzz.go
package fuzz

func Fuzz(data []byte) int {
	s := string(data)
	if s == "bad string" {
		panic("bad string found")
	}
	return 0
}
