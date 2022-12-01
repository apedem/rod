// Package main ...
package main

import (
	"fmt"

	"github.com/apedem/rod/lib/launcher"
	"github.com/apedem/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
