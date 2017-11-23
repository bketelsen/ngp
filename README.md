## ngp - New Go Package

## Description

`ngp` is an opinionated helper utility that creates some boilerplate for a new Go command or package.

Featuring:

- HelloGopher - modified makefile inspired by Cloudflare's HelloGopher
- Docker integration

## Output

Run `ngp` in an empty directory.  *IT WILL OVERWRITE THINGS RIGHT NOW*  In the future, it may move existing files or directories that would have been overwritten.

`ngp` will create:

- Dockerfile for a project with appropriate settings for a Go command
- Makefile suitable for any Go project.  Based on a modified "HelloGopher" makefile by Cloudflare.


## Requirements and Notes


Docker is required for Docker builds.

Doesn't work in Windows without `make` installed.

## Project

Start with an empty directory where you intend to build your project.  This directory should be in your GOPATH.

After running `ngp`, your project will have a Makefile which has everything you need to get started.  

Start with `setup`:

```
make setup
```


## Make Targets

- all - run `test` and `build` targets
- bin/ - install coverage, deps, and imports helpers
- build - make the target binary
- clean - remove bin
- cover - run coverage report
- deps - run `dep ensure` to install dependencies
- docker - build the docker image
- format - format the source code
- list - list build targets
- setup - create the project structure and install tools
- test  - run tests
- tags - list git tags
