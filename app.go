package main

import (
	"os"
	"fmt"
	"os/exec"
	"io"
	"log"
)

func usage() {
	fmt.Printf("Usage: %s <tree-ish> <docker-tag>\n", os.Args[0])
}

func main() {

	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	treeish := os.Args[1]
	tag     := os.Args[2]

	gitArchive := exec.Command("git", "archive", treeish)

	gitArchiveOut, err := gitArchive.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	gitArchiveErr, err := gitArchive.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	dockerBuild := exec.Command("docker", "build", "-t", tag, "-")

	dockerBuildOut, err := dockerBuild.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	dockerBuildErr, err := dockerBuild.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	dockerBuildIn, err := dockerBuild.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = gitArchive.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = dockerBuild.Start()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		io.Copy(os.Stdout, dockerBuildOut)
	}()

	go func() {
		io.Copy(os.Stderr, dockerBuildErr)
	}()

	go func() {
		io.Copy(os.Stderr, gitArchiveErr)
	}()

	io.Copy(dockerBuildIn, gitArchiveOut)

	err = gitArchive.Wait()
	if err != nil {
		log.Fatal(err)
	}

	dockerBuildIn.Close()

	err = dockerBuild.Wait()
	if err != nil {
		log.Fatal(err)
	}

}