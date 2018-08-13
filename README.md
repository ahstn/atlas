# atlas
> Swiss army-knife for software building and development.

[![GoDoc](https://godoc.org/github.com/ahstn/atlas?status.svg)](https://godoc.org/github.com/ahstn/atlas)
[![Go Report Card](https://goreportcard.com/badge/ahstn/atlas)](https://goreportcard.com/report/ahstn/atlas)
[![CircleCI](https://circleci.com/gh/ahstn/atlas/tree/master.svg?style=shield)](https://circleci.com/gh/ahstn/atlas/tree/master)
[![codecov](https://codecov.io/gh/ahstn/atlas/branch/master/graph/badge.svg)](https://codecov.io/gh/ahstn/atlas)

**WORK IN PROGRESS. Only for dev and experimental use.**

# Table of Contents

* [Introduction](#introduction)
* [Status](#status)
* [Preview](#preview)
* [Features](#features)
* [Commands](#commands)
  * [`atlas project`](#atlas-project)
  * [`atlas repo`](#atlas-repo)
  * [`atlas docker`](#atlas-docker)
* [Config](#config)

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

Feature requests in the form of GitHub issues and comments on current issues are
welcomed and encourged to help shape `atlas`.

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

# Commands
## `atlas project`
The main focal point of atlas. In modern development microservices are the hip
architecture/design pattern and each service can be in it's on isolated dir.
This complicates things when, as a developer, you just want to build and run
your application stack or a segment of it.

`atlas` attempts to solve this by taking a simplistic, flexible config file as
input that refers to the application stack you wish to build and run. From there
you can customise whether tests are run for certain apps, if you want Docker
builds to run, what arguments should be passed when running the application and
much more!

You can read more about the configuration file and usage in the [config section](#config).

```
➜ atlas help project
NAME:
   atlas project - build project (collection of services)

USAGE:
   atlas project [command options] [arguments...]

OPTIONS:
   --config value, -c value  name of config file in ~/.config/atlas (default: "atlas.yaml")
   --skipTests, -s           skip tests
   --verbose, -V             verbose logging rather than progress bars
```

## `atlas repo`
Ever wished you could just run a command that would open your Git repo in the
browser, instead of manually switching to Chrome and navigating GitHub? Well
now you can! :grin:

```
➜ atlas help repo
NAME:
   atlas repo - open Git repo in browser

USAGE:
   atlas repo
```

## `atlas docker`
Essentially the same basic funationality as `docker build` with the regular
stdout replaced in favour of a more elegant output.

```
➜ atlas help docker
NAME:
   atlas docker - build an application's Dockerfile

USAGE:
   atlas docker [command options] [directory containing Dockerfile]

OPTIONS:
   --tag name:tag, -t name:tag    name and tag image in the name:tag format
   --arg arg=value, -a arg=value  build arguments in the arg=value format (space seperated)
   --config value, -c value       name of config file in ~/.config/atlas (default: "atlas.yaml")
   --verbose, -V                  verbose logging rather than progress bars
```

# Config
To be finalised and documented..

In the meantime, the project wiki has some information on using config: [`atlas` wiki](https://github.com/ahstn/atlas/wiki/Config)

[eugenp/tutorials]: https://github.com/eugenp/tutorials/
