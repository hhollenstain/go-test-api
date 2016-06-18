package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You can't even use this api right you meatbags\n")
}

func Bender(w http.ResponseWriter, r *http.Request) {
	lines, err := readLines("bender_lines.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Fprint(w, lines[rand.Intn(len(lines))])
}
