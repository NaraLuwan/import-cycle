package main

import "github.com/NaraLuwan/import-cycle/a"

func main() {
	a := a.New(3)
	a.Pb.DisplayC()
}
