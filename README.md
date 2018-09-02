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
  * [`atlas docker`](#atlas-docker)
  * [`atlas repo`](#atlas-repo)
  * [`atlas git`](#atlas-git)
  * [`atlas issues`](#atlas-issues)
* [Config](#config)

## Introduction
`atlas` is a CLI tool that leverages development applications to make common
tasks more efficient.
One of the main purposes is to make builds simpler across multiple repos while
providing an aesthetic user interface.
```
➜ atlas help
NAME:
   atlas - Make Development Great Again

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
     build, b    execute the application build process
     docker, d   build an application's Dockerfile
     issues, i   open JIRA/Github issue page for current Git project
     project, p  build project (collection of services)
     repo, r     open Git repo in browser
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
Maven and Docker build example using the config file in `examples/java/`:
[![asciicast](https://asciinema.org/a/197066.png)](https://asciinema.org/a/197066)


## Features
Documentation to be completed, and features are in progress but for now:
* Reworked Maven output (as seen in [Preview](#preview)) intended to provide a more aesthetically pleasing UI.
* Config files intended for automonous building of multi-repos.
  * Including application builder (mvn, npm, etc), Docker builds and Docker runs.
* Opening Git repository in browser from the terminal.
* Allowing users to execute Git commands against multiple repos at once (clone, checkout, update).
* Opening GitHub/JIRA issue in browser from the terminal.
* .gitignore generator with language detection. (not yet impl)
* Dockerfile generator with language detection. (not yet impl)

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

You can read more about the configuration file and usage in
the [config section](#config).
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

## `atlas git`
Wrapper around Git that allows the user to execute commands against many repos
at once.
Instead of having to jump into multiple directories, then `git pull` or
`git checkout -b new-feature`,
`atlas` can handle this in one command.

From the usage below you can see the available commands.
One example to run would be `atlas git clone -c atlas.yaml -e auth`.
Which will clone all the applications specified in your config into the `config
root dir`, excluding the `auth` service.
```
➜ atlas help git
NAME:
   atlas git - preform Git actions against service(s)

USAGE:
   atlas git [global options] command [command options] [arguments...]

COMMANDS:
     branch      create a branch in the services' repo(s) defined in config
     clone       clone the services' repo(s) defined in config
     checkout    checkout a branch in services' repo(s) defined in config
     update, up  pull updates from remote, but keep local changes

GLOBAL OPTIONS:
   --help, -h  show help

SUBCOMMAND OPTIONS:
   --config value, -c value   name of config file in ~/.config/atlas (default: "atlas.yaml")
   --exclude value, -e value  exclude certain services defined in config from the command
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

## `atlas issues`
The same concept as `atlas repo` only for issue trackers! When ran from a repo,
the issue tracker will be opened in your browser. The only downside is that
private issue trackers must be supplied using the `--url` flag.
```
➜ atlas help issues
NAME:
   atlas issues - open JIRA/Github issue page for current Git project

USAGE:
   atlas issues [command options] [arguments...]

OPTIONS:
   --url URL, -u URL  private JIRA base URL (without '/issues' or '<team>')
```

# Config
Config files are the backbone of `atlas`' `project` command and aid in one of
the main goals of the project. These file(s) act as a set of requirements or
steps that you, the user, want to happen. The file itself can be passed into
some commands via the `--config [FILE]` flag. `atlas` will firstly look in the
current working dir for the file and then in it's config dir (`~/.config/atlas/`).
The goal of this is too allow easy and quick switching between multiple config
files.

For `atlas` to operate on multiple applications, they must have a common parent
directory. For most users this will likely be your `~/git/` folder and should be
specified in the config as `root`. The config should then contain an array of
applications with metadata and information about the taks you want to preform.
As a core idea, the application information should describe the end state and
shouldn't require in-depth details describing your tasks or the process.

An annoted example describing the config file and it's available fields is in
the project root: [atlas.example.yaml]. Specific relatable examples exist in the
 [examples folder]. For now only a simple Java "microservices" setup is there,
 but hopefully it should give you a rough idea.

For more information on using config, please refer to the project
wiki: [`atlas` config wiki]

[atlas.example.yaml]: ./atlas.example.yaml
[examples folder]: ./examples
[`atlas` config wiki]: https://github.com/ahstn/atlas/wiki/Config
