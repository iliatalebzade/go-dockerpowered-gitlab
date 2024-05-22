package main

import "gitlaboratory/laboratory"

func main() {
	_, err := laboratory.Initialization()
	if err != nil {
		panic(err.Error())
	}

	_, err = laboratory.RenameBranch("main")
	if err != nil {
		panic(err.Error())
	}

	_, err = laboratory.CommitAll("main")
	if err != nil {
		panic(err.Error())
	}
}
