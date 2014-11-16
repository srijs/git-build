git-build
=========

Build a git repository with a `Dockerfile`. Stop worrying about a dirty working directory or the wrong `HEAD`.

## Installation

You can get binaries for Linux, Windows and OS X from [gobuild.io](http://gobuild.io/github.com/srijs/git-build).
Install the binary somewhere in your `$PATH`.

## Usage

From inside the git repository you want to build, call `git build <tree-ish> <path>`,
where `<tree-ish>` is the name of the branch or tag you want to build,
and `<path>` is the path you want to include in the build (`.` for the current directory).

The files in the specified tree will be uploaded to the Docker daemon,
which will build them and store the resulting image as `name:tag`,
where `name` is the base name of the path, and tag is the specified tree.

For example:

    git build master ./bar

This uploads the content of `./bar` in the state of the `master` branch,
builds it according to `./bar/Dockerfile`, and saves the resulting image as `bar:master`.
