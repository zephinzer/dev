# Development Notes

- [Development Notes](#development-notes)
  - [Getting started](#getting-started)
  - [Directory structure](#directory-structure)
  - [Tools used for development](#tools-used-for-development)
    - [Array2go](#array2go)
    - [`rsrc`](#rsrc)
    - [SQLite3 CLI](#sqlite3-cli)
- [Building](#building)
  - [General build notes](#general-build-notes)
  - [Build notes for Linux](#build-notes-for-linux)
  - [Build notes for MacOS](#build-notes-for-macos)
  - [Build notes for Windows](#build-notes-for-windows)
- [Releasing](#releasing)
- [Distribution](#distribution)
- [References](#references)

## Getting started

1. Clone this repository using `git clone git@gitlab.com:usvc/utils/dev.git`
2. Install dependencies using `make deps`
3. Create a local development configuration file at `./dev.yaml` relative to the project root containing [the sample configuration file](#sample-configuration-file)

## Directory structure

This project adheres to the [Go project layout as defined here](https://github.com/golang-standards/project-layout).

## Tools used for development

### Array2go

This tool is used to convert a PNG image to an array that Go can compile and use as the system tray icon. To install it:

```sh
go get github.com/cratonica/2goarray;
```

See the `prepare_icon` in [the `Makefile`](./Makefile) on usage.

### `rsrc`

This tool is used to compile manifest resources for Windows builds. To install it:

```sh
go get github.com/akavel/rsrc;
```

### SQLite3 CLI

The local database uses SQLite3, to install it run the install for your appropriate platform:

| Platform | Command |
| --- | --- |
| Ubuntu | `apt-get install sqlite3` |
| Fedora | `dnf install sqlite` |
| Archlinux | `pacman -S sqlite` |
| MacOS | `brew install sqlite3` |
| Windows | `choco install sqlite` |
| Alpine | `apk add sqlite` |

# Building

As this is a desktop app meant for cross-platform distribution, this gets a little complicated. The instructions assume an Ubuntu build environment.

## General build notes

1. Run `setup_build` to install the required:
   1. `2goarray` for converting a PNG icon into Go code
   2. `rsrc` for compiling Windows application manifests
2. Run `setup_build_linux` if you're building from a linux environment. This installs (package names may vary across distros, these are for Ubuntu 18.04):
   1. `libgtk-3-dev`
   2. `libappindicator3-dev`
   3. `libwebkit2gtk-4.0-dev`

## Build notes for Linux

1. The Linux assets can be found at `./assets/linux` relative to the project root.

## Build notes for MacOS

1. The MacOS assets can be found at `./assets/macos` relative to the project root.

## Build notes for Windows

1. The Windows assets can be found at `./assets/windows` relative to the project root.

# Releasing

This should be done automatically via the CI pipeline.

# Distribution

> TODO

# References

- On building a static binary file with libraries requiring CGO: [Golang w/SQLite3 + Docker Scratch Image](https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image)
- On executing CLI commands from Go [Advanced command execution in Go with os/exec](https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html)
