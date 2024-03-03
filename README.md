[![Redis-Populate CI](https://github.com/NY-Daystar/Populate-Redis/actions/workflows/ci.yaml/badge.svg)](https://github.com/NY-Daystar/Populate-Redis/actions/workflows/ci.yaml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Version](https://img.shields.io/github/tag/LucasNoga/populate-redis.svg)](https://github.com/LucasNoga/populate-redis/releases)
[![Total views](https://img.shields.io/sourcegraph/rrc/github.com/LucasNoga/populate-redis.svg)](https://sourcegraph.com/github.com/LucasNoga/populate-redis)

![GitHub watchers](https://img.shields.io/github/watchers/ny-daystar/populate-redis)
![GitHub forks](https://img.shields.io/github/forks/ny-daystar/populate-redis)
![GitHub Repo stars](https://img.shields.io/github/stars/ny-daystar/populate-redis)
![GitHub repo size](https://img.shields.io/github/repo-size/ny-daystar/populate-redis)
![GitHub language count](https://img.shields.io/github/languages/count/ny-daystar/populate-redis)
![GitHub top language](https://img.shields.io/github/languages/top/ny-daystar/populate-redis) <a href="https://codeclimate.com/github/ny-daystar/populate-redis/maintainability"><img src="https://api.codeclimate.com/v1/badges/715c6f3ffb08de5ca621/maintainability" /></a>  
![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/m/ny-daystar/populate-redis/main)
![GitHub issues](https://img.shields.io/github/issues/ny-daystar/populate-redis)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/ny-daystar/populate-redis)
![GitHub](https://img.shields.io/github/license/ny-daystar/populate-redis)
[![All Contributors](https://img.shields.io/badge/all_contributors-1-blue.svg?style=circular)](#contributors)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

# Populate Redis

Golang tool to feed redis database with fake users data for testing purposes.

Source code analysed with [DeepSource](https://deepsource.com/)

## Summary

-   [Prerequisites](#prerequisites)
    -   [Setup redis database](#setup-redis-database)
-   [Installing go dependencies](#installing-go-dependencies)
-   [How to Use](#how-to-use)
-   [How to Run](#how-to-run)
    -   [Flush parameters](#flush-parameter)
-   [Unit tests]
-   [Contact](#contact) -[Credits](#credits)

## Prerequisites

-   [Golang](https://golang.org/dl/) >= 1.21.0
-   [Docker Engine](https://www.docker.com/products/docker-desktop/) >= 24.0.7

### Setup redis database

1. Pull redis docker image

```bash
docker pull redis
```

2. Launch redis database

```bash
docker run --name redis -p 6479:6379 -d redis
```

3. Test the connection

```bash
redis-cli -p 6479
```

> You can also use [Redis Insight](https://redis.com/fr/redis-enterprise/redisinsight/)

## Installing Go dependencies

In order to install the dependencies, you should run `go get` once you've pulled
this project

## How to use

Using `go run . --help` will display the list of arguments that you can give to
this script

```bash
  -n int        Number of users to generate (default 10)
  -v            Make the process verbose or not
```

Note: Boolean arguments (`-v` and `-f`) will be set to `true` if they are
provided, there is no need to give them a value like : `-v true`.

Example :

`go run . -f -n 10`

Will enable the flush feature, but keep the verbose feature disabled

## How to run

You can run this script with `go run . [args]`

### Flush parameter

The flush parameter will prompt you with the following message before actually
flushing the database, and display the current database it is looking to flush.

```
2021/04/08 15:46:49 Are you SURE you want to flush this database ? localhost:6379 [y/n]:
```

If you decide not to flush the database, it will still go on to generate users.

You can always interrupt the process with `Ctrl+C` at this step if you do not
want to generate users.

## Unit tests

To launch unit tests

```bash
go test ./...
```

## Contact

    To make a pull request: https://github.com/LucasNoga/redis-populate/pulls
    To summon an issue: https://github.com/LucasNoga/redis-populate/issues
    For any specific demand by mail: luc4snoga@gmail.com

## Credits

Made by Lucas Noga.
Licensed under GPLv3.
