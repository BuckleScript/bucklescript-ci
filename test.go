

package main

import (
		"os/exec"
		"fmt"
		"log"
		"os"
		"sync"
		)

func checkError(err error){
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func testTheme(theme string){
	fmt.Println("Removing",theme)
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
	var wg sync.WaitGroup
	for _, theme := range []string{"basic", "basic-reason","generator","minimal"} {
		fmt.Println("Test theme", theme)
		wg.Add(1)
		go (func(theme string){
			defer wg.Done()
			testTheme(theme)
		})(theme)
	}
	wg.Wait()
}		