package main

import "github.com/remotehy/go-std-mod/app/tom/di"

func main() {

	c := di.NewService()
	c.HandleRequest("tom")
}
