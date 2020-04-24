# Dev

The ultimate developer experience CLI tool.

- [Dev](#dev)
- [Usage](#usage)
  - [Table of Canonical Tokens](#table-of-canonical-tokens)
  - [Initialisation](#initialisation)
    - [Database](#database)
  - [Retrieval](#retrieval)
    - [Account Information](#account-information)
      - [Github Account Information](#github-account-information)
      - [Gitlab Account Information](#gitlab-account-information)
      - [Pivotal Tracker Account Information](#pivotal-tracker-account-information)
    - [Notifications](#notifications)
      - [Gitlab Notifications](#gitlab-notifications)
      - [Pivotal Tracker Notifications](#pivotal-tracker-notifications)
    - [Work](#work)
      - [Pivotal Tracker Work](#pivotal-tracker-work)
- [Setting Up](#setting-up)
  - [Configuration](#configuration)
    - [Sample configuration file](#sample-configuration-file)
  - [Platforms](#platforms)
    - [Github](#github)
      - [Setting Up](#setting-up-1)
    - [Gitlab](#gitlab)
    - [PivotalTracker](#pivotaltracker)
      - [Setting Up](#setting-up-2)
- [Development Runbook](#development-runbook)
  - [Getting Started](#getting-started)
  - [References](#references)
- [Licensing](#licensing)

- - -

# Usage

The following is an overview of what can be done:

```sh
# initialise persistent database
dev initialise database;

# retrieve consumed configuration
dev get config;

# retrieve notifications (todos) from gitlab
dev get notifications gitlab;

# retrieve notifications from pivotal tracker
dev get notifications pivotaltracker;

# retrieve your work from pivotal tracker
dev get work pivotaltracker;

# retrieve information about your accounts from github
dev get account github;

# retrieve information about your accounts from gitlab
dev get account gitlab;

# retrieve your accounts from pivotal tracker
dev get account pivotaltracker;
```

## Table of Canonical Tokens

| Concept | Type | Canon | Aliases |
| --- | --- | --- |
| Account | Noun | `account` | `accounts`, `acc`, `accs`, `a` |
| Configuration | Noun | `configuration` | `config`, `conf`, `cf`, `c` |
| Gitlab | Noun | `gitlab` | `gl` |
| Github | Noun | `github` | `gh` |
| Notifications | Noun | `notifications` | `notification`, `notif`, `notifs`, `n` |
| PivotalTracker | Noun | `pivotaltracker` | `pivotal`, `pt` |
| Work | Noun | `work` | `stories`, `tasks`, `tickets`, `w` |
| Get | Verb | `get` | `retrieve`, `g` |
| Initialise | Verb | `initialise` | `initialize`, `init`, `i` |

## Initialisation

Example: `dev init -h`

### Database

Example: `dev init db`

## Retrieval

Example: `dev get -h`

### Account Information

Example: `dev get account -h`

#### Github Account Information

Example: `dev get account github`

#### Gitlab Account Information

Example: `dev get account gitlab`

#### Pivotal Tracker Account Information

Example: `dev get account pivotaltracker`

### Notifications

Example: `dev get notifs -h`

#### Gitlab Notifications

Example: `dev get notifs gitlab`

#### Pivotal Tracker Notifications

Example: `dev get notifs pivotaltracker`

### Work

Example: `dev get work -h`

#### Pivotal Tracker Work

Example: `dev get work pivotaltracker`


- - -

# Setting Up

## Configuration

Configuration is done via YAML.

### Sample configuration file

```yaml
# this defines networks that should be reachable from your machine
networks:
  - name: internal vpn
    check:
      schema: http
      hostname: gitlab.internal.domain.com
      path: /
# this defines software that should be on your machine
software:
  - name: golang
    check:
      command:
        - go
        - version
      exitCode: 0
      stdout: go version go[/d\.]
# this defines platforms that the developer should have access to
platforms:
  pivotaltracker:
    accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    projects:
      - name: work
        projectID: "XXXXXXX"
      - name: personal
        projectID: "XXXXXXX"
        accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  github:
    accounts:
      - name: personal
        accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  gitlab:
    accounts:
      - name: personal
        accessToken: XXXXXXXXXXXXXXXXXXXX
      - name: work-on-prem
        hostname: gitlab.yourdomain.com
        accessToken: XXXXXXXXXXXXXXXXXXXX
```

## Platforms

### Github

#### Setting Up

Retrieve your `accessToken` by generating a new personal access token from [https://github.com/settings/tokens](https://github.com/settings/tokens). You'll need the following scopes:

- repo:status
- repo_deployment
- public_repo
- repo:invite
- read:packages
- read:org
- read:public_key
- read:repo_hook
- notifications
- read:user
- read:discussion
- read:enterprise
- read:gpg_key

### Gitlab

Retrieve your `accessToken` by generating a new personal access token from [https://gitlab.com/profile/personal_access_tokens](https://gitlab.com/profile/personal_access_tokens). You'll need the following scopes:

- api
- read_api

> If you're using an on-premise Gitlab, change `gitlab.com` to your Gitlab's hostname

### PivotalTracker

#### Setting Up

Retrieve your `accessToken` from [https://www.pivotaltracker.com/profile](https://www.pivotaltracker.com/profile).

- - -

# Development Runbook

## Getting Started

1. Clone this repository using `git clone git@gitlab.com:usvc/utils/dev.git`
2. Install dependencies using `make deps`
3. Create a local development configuration file at `./dev.yaml` relative to the project root containing [the sample configuration file](#sample-configuration-file)

## References

- On building a static binary file with libraries requiring CGO: [Golang w/SQLite3 + Docker Scratch Image](https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image)
- On executing CLI commands from Go [Advanced command execution in Go with os/exec](https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html)

# Licensing

MIT
