package main

import "fmt" 

func main() {
	forPythonStyle()
}

func forPythonStyle(){
	strings:= []string {"hello","world","nie"}
	for i,s:=range strings {
		fmt.Println(i,s)
	}
}