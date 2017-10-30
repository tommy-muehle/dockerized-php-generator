# Docker image for [{{ .Name }}]({{ .Url }})

[![Build Status](https://travis-ci.org/dockerized-php/{{ .Name }}.svg?branch=master)](https://travis-ci.org/dockerized-php/{{ .Name }})

The image is based on [Alpine Linux](https://alpinelinux.org/) and built daily.

## Supported tags

- `latest` [(latest/Dockerfile)](https://github.com/dockerized-php/{{ .Name }}/blob/master/latest/Dockerfile)

## How to use this image

### Install

Install the container:

```
docker pull dockerizedphp/{{ .Name }}
```

Or alternatively, pull a specific version:

```
docker pull dockerizedphp/{{ .Name }}:0.1
```

### Usage

We are recommend to use this image as an shell alias to access via short-command.
To use simply *{{ .Name }}* everywhere on CLI, add this line to your ~/.zshrc, ~/.bashrc or ~/.profile.

```
alias {{ .Name }}='docker run -ti -v $PWD:/app --rm dockerizedphp/{{ .Name }}:latest'
```

Otherwise you can use this command directly.
