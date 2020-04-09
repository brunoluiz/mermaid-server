# Mermaid Live Server

⚡️ Serve Mermaid (MMD) files with live reloading

- Serves Mermaid
- Live-reload on save

## Install

### MacOS

Use `brew` to install it

```
brew tap brunoluiz/tap
brew install mermaid-server
```

### Linux and Windows

[Check the releases section](https://github.com/brunoluiz/mermaid-server/releases) for more information details 

### Docker

The tool is available as a Docker image as well. Please refer to [Docker Hub page](https://hub.docker.com/r/brunoluiz/mermaid-server/tags) to pick a release

### go get

Please avoid using `go get`. [Check our releases](https://github.com/brunoluiz/mermaid-server/releases) for more details.

## Usage

To run it, use `mermaid-server --input <file>`


```
NAME:
   mermaid-server - ⚡️ Serve Mermaid (MMD) files with live reloading

USAGE:
   mermaid-server [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --address value  server hostname (default: ":3000")
   --input value    mmd input file
   --help, -h       show help (default: false)
```
