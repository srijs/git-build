package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

func usage() {
	fmt.Printf("Usage: %s [options...] <tree-ish> <path>\n\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage

	registryPtr := flag.String("publish", "", "publish the image to a docker registry")

	flag.Parse()

	if flag.NArg() != 2 {
		usage()
		os.Exit(2)
	}

	treeish := flag.Arg(0)
	buildpath := flag.Arg(1)
	registry := *registryPtr

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	wd := path.Join(cwd, buildpath)

	name := path.Base(wd)
	if len(registry) > 0 {
		name = registry + "/" + name
	}

	fmt.Printf("Building tree '%s' in %s as '%s:%s'...\n", treeish, wd, name, treeish)

	gitArchive := exec.Command("git", "archive", treeish)

	gitArchive.Dir = wd

	gitArchiveOut, err := gitArchive.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	gitArchiveErr, err := gitArchive.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	dockerBuild := exec.Command("docker", "build", "-t", name+":"+treeish, "-")

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

	if len(registry) == 0 {
		os.Exit(0)
	}

	fmt.Printf("Publishing to %s...\n", *registryPtr)

	dockerPush := exec.Command("docker", "push", name)

	dockerPushOut, err := dockerPush.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	dockerPushErr, err := dockerPush.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		io.Copy(os.Stdout, dockerPushOut)
	}()

	go func() {
		io.Copy(os.Stderr, dockerPushErr)
	}()

	dockerPush.Run()

}
