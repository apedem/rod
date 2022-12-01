// Package main ...
package main

import (
	"github.com/apedem/rod/lib/utils"
)

func main() {
	utils.Exec("npx -ys -- cspell@6.14.2 --no-progress **")

	utils.Exec("npx -ys -- eslint@8.7.0 --config=lib/utils/lint/eslint.yml --ext=.js,.html --fix --ignore-path=.gitignore .")

	utils.Exec("npx -ys -- prettier@2.5.1 --loglevel=error --config=lib/utils/lint/prettier.yml --write --ignore-path=.gitignore .")

	utils.Exec("go run github.com/ysmood/golangci-lint@latest")

	lintMustPrefix()

	checkGitClean()
}

func checkGitClean() {
	out := utils.ExecLine(false, "git status --porcelain")
	if out != "" {
		panic("Please run \"go generate\" on local and git commit the changes:\n" + out)
	}
}
