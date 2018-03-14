package main

import (
	"github.com/sean-/seed"

	docker_prompt "github.com/zbiljic/docker-prompt/cmd"
)

func init() {
	/* #nosec */
	seed.Init()
}

func main() {
	docker_prompt.Main()
}
