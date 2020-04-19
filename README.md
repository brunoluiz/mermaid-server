# Mermaid Live Server

⚡️ Serve Mermaid (MMD) files with live reloading

- Serves Mermaid
- Live-reload on save

## Install

### MacOS

Use `brew` to install it

```
brew install brunoluiz/tap/mermaid-server
```

### Linux and Windows

[Check the releases section](https://github.com/brunoluiz/mermaid-server/releases) for more information details 

### Docker

The tool is available as a Docker image as well. Please refer to [Docker Hub page](https://hub.docker.com/r/brunoluiz/mermaid-server/tags) to pick a release

### go get

Install using `GO111MODULES=off go get github.com/brunoluiz/mermaid-server/cmd/mermaid-server` to get the latest version. This will place it in your `$GOPATH`, enabling it to be used anywhere in the system.

**⚠️ Reminder**: the command above download the contents of master, which might not be the stable version. [Check the releases](https://github.com/brunoluiz/mermaid-server/releases) and get a specific tag for stable versions.

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
   --address value, -a value  server hostname (default: ":7777")
   --input value, -i value    mmd file path
   --config value, -c value   mmd config json file path
   --help, -h                 show help (default: false)
```
