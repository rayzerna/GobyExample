package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	_ = grepCmd.Start()
	_, _ = grepIn.Write([]byte("hello grep\ngoobye grep"))
	_ = grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	_ = grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -alh")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
