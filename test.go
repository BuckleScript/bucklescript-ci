

package main

import (
		"os/exec"
		"fmt"
		"log"
	"os"
		)

func checkError(err error){
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func testTheme(theme string){
	os.RemoveAll(theme)
	cmd := exec.Command("bsb", "-theme",theme,"-init",theme)
	output, err:= cmd.CombinedOutput()
	
	fmt.Println(string(output))
	checkError(err)
	
	fmt.Println("Started to build ")
	cmd2 := exec.Command("npm", "run", "build")		
	cmd2.Dir = theme
	output2, err:= cmd2.CombinedOutput()
	fmt.Println(string(output2))
	checkError(err)		
	os.RemoveAll(theme)
}
func main(){

	output, err := exec.Command("which", "bsb").Output()
	checkError(err)	
	fmt.Println("bsb", "is", string(output))
	
	// testTheme("basic")
	// testTheme("generator")
	for _, theme := range []string{"basic", "basic-reason","generator","minimal"} {
		fmt.Println("Test theme", theme)
		testTheme(theme)
	}
}		