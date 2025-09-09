package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: spawner <student-name>")
		return
	}

	student := os.Args[1]
	password := student + "123" // simple password rule

	cmd := exec.Command("docker", "run", "-d",
		"--name", student,
		"--network=pwd-setup_pwdnet",
		"-e", "PASSWORD="+password,
		"-v", "./students/"+student+":/config",
		"lscr.io/linuxserver/code-server:latest")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error creating container:", err)
		return
	}

	fmt.Printf("Student: %s\nURL: http://%s.pwd.asmryz.com\nPassword: %s\n",
		student, student, password)
}
