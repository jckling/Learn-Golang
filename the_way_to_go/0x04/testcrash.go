package main

func main() {
	var p *int = nil
	*p = 0	// panic: runtime error: invalid memory address or nil pointer dereference
}