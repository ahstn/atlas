# atlas
> Make Development Great Again

[![GoDoc](https://godoc.org/github.com/ahstn/atlas?status.svg)](https://godoc.org/github.com/ahstn/atlas)
[![Go Report Card](https://goreportcard.com/badge/ahstn/atlas)](https://goreportcard.com/report/ahstn/atlas)

**WORK IN PROGRESS. Only for dev and experimental use.**

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
     project, p  Build Project (Collection of Services)
     repo, r     Open Git repo in browser
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -s, --skipTests  skip tests
   -V, --verbose    verbose logging rather than progress bars
   --help, -h       show help
   --version, -v    print the version
```

## Overview
Very basic example for now:
[![asciicast](https://asciinema.org/a/vcZS0r2z15HXiusFTBHGPQtSQ.png)](https://asciinema.org/a/vcZS0r2z15HXiusFTBHGPQtSQ)
