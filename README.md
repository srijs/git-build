git-build
=========

Build a git repository with a `Dockerfile`. Stop worrying about a dirty working directory or the wrong `HEAD`.

## Why?

Frequently building and publishing docker images from different branches in a repository is cumbersome.

You have to:

1. Make sure you checked out the right branch
2. Make sure your working directory is clean (you don't want to publish your secret config or that huge test asset)
3. Upload the current directory to the Docker daemon, and have him build the image
4. Push the image to the designated Docker registry

`git-build` makes this a one liner.

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

Optionally, you can specify a docker registry via the `-publish` flag,
and the image will be published to that registry.

For example:

    git build -publish=index.docker.io/baz master ./bar

This uploads the content of `./bar` in the state of the `master` branch,
builds it according to `./bar/Dockerfile`, and saves the resulting image as `bar:master`.
If the build is successful, the image will be uploaded to the user repository `baz/bar`
in `index.docker.io`.
