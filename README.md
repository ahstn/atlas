# atlas
> Make Development Great Again

[![GoDoc](https://godoc.org/github.com/ahstn/atlas?status.svg)](https://godoc.org/github.com/ahstn/atlas)
[![Go Report Card](https://goreportcard.com/badge/ahstn/atlas)](https://goreportcard.com/report/ahstn/atlas)

# Table of Contents

* [Introduction](#introduction)
* [Status](#status)
* [Preview](#preview)
* [Features](#features)

## Introduction
`atlas` is a CLI tool that leverages development applications to make common tasks more efficient.
One of the main purposes is to make builds simpler across multiple repos while providing an aesthetic user interface.

```
➜ atlas help
NAME:
   atlas - Make Development Great Again

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
     build, b    execute the application build process
     project, p  build project (collection of services)
     repo, r     open Git repo in browser
     docker, d   build an application's Dockerfile
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --skipTests, -s  skip tests
   --verbose, -V    verbose logging rather than progress bars
   --help, -h       show help
   --version, -v    print the version
```

## Status
`atlas` is currently **only for dev and experimental use.** Development work is
currently in progress and in the proof of concept stage. The intention is that
version 0.1.0 will be suitable for early-access testing and 0.2.0 should include
most features while being adequately tested.

## Preview
Simple example using a config file with two projects from [eugenp/tutorials]:
[![asciicast](https://asciinema.org/a/vcZS0r2z15HXiusFTBHGPQtSQ.png)](https://asciinema.org/a/vcZS0r2z15HXiusFTBHGPQtSQ)


## Features
Documentation to be completed, and features are in progress but for now:
* Reworked Maven output (as seen in [Preview](#preview)) intended to provide a more aesthetically pleasing UI.
* Config files intended for automonous building of multi-repos.
  * Including application builder (mvn, npm, etc), Docker builds and Docker runs.
* Opening Git repository in browser from the terminal.
* Opening GitHub/JIRA issue in browser from the terminal.
* .gitignore generator with language detection.
* Dockerfile generator with language detection.

[eugenp/tutorials]: https://github.com/eugenp/tutorials/
