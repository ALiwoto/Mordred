package main

import "packs/security"

func main() {
	_test := security.GenerateStrongString("test")
	_test.Length()

}