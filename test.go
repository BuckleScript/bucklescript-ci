

package main

import (
		"os/exec"
		"fmt"
		"log"
		)

func checkError(err error){
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
func main(){

	output, err := exec.Command("which", "bsb").Output()
	checkError(err)	
	fmt.Println(
		string(output))
}		