package main

import (
	"fmt"
	"github.com/OJ/gobuster/v3/cli/cmd"
	"os"
)

//----------------------------------------------------
// Gobuster -- by OJ Reeves
//
// A crap attempt at building something that resembles
// dirbuster or dirb using Go. The goal was to build
// a tool that would help learn Go and to actually do
// something useful. The idea of having this compile
// to native code is also appealing.
//
// Run: gobuster -h
//
// Please see THANKS file for contributors.
// Please see LICENSE file for license details.
//
//----------------------------------------------------
func init(){
	if len(os.Args) == 1{
		fmt.Println("Usage: ")
		fmt.Println(os.Args[0]+" -l d.txt -o out.txt -r dns-ip\n")
		os.Exit(0)
	}
}


func main() {
	cmd.Execute()
}
