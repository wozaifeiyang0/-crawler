package main

import thy "github.com/wozaifeiyang0/thylog"

func main() {

	thy.SetLogFile("./data", "error.log")

	thy.Error.Println("aaaaaaaaaaa")

}
