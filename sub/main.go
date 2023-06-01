package main

import "sub/src/workers"

func main() {
	workers.New(workers.Counter, 5)
}
