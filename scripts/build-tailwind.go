// scripts/build-tailwind.go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"npx", "tailwindcss",
		"-i", "./web/input.css",
		"-o", "./web/static/css/tailwind.css",
		"--minify",
	)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		fmt.Println("Install Tailwind dulu: npm install -D tailwindcss")
		panic(err)
	}
	fmt.Println("Tailwind built!")
}
