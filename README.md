git-build
=========

Build a git repository with a `Dockerfile`. Stop worrying about a dirty working directory or the wrong `HEAD`.

## Installation

You can get binaries for Linux, Windows and OS X from [gobuild.io](http://gobuild.io/github.com/srijs/git-build). Install the binary somewhere in your `$PATH`.

## Usage

From the root of the repository you want to build, call `git build <tree-ish> <docker-tag>`, where `<tree-ish>` is the name of the branch or tag you want to build, and `<docker-tag>` is the tag you want to give to the newly built docker image.
