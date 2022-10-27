package main

import (
	"github.com/i0Ek3/blogie/pkg/server"
	_ "github.com/i0Ek3/blogie/pkg/setup"
)

// @title blogie
// @version 1.0
// @description A blog backend program developed with Gin.
// @termOfService https://github.com/i0Ek3/blogie
func main() {
	server.Boot()
}
