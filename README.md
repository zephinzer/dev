# Dev

Pimp your developer experience.

- ✅ Open website of repository you're in - `dev open repo`
- ✅ Check if you have required software - `dev check software`
- ✅ Check if you have required network connectivity - `dev check networks`
- ✅ Get work assigned to you on - `dev get work`
- ✅ Get notifications from developer platforms - `dev get notifications`
- ✅ Intuitive aliases that become simpler as you use them (example: from `dev get notifications` to `dev get notifs` to `dev g n`)

This tool exists to improve the day-to-day work experience of software developers. Use this to:

1. Improve onboarding process through defining required software in code
2. Improve day-to-day work experience through defining required network connections in code
3. Improve productivity through moving checking/notification activities to the command line

- [Dev](#dev)
- [Installation](#installation)
  - [Via Git Repository](#via-git-repository)
  - [Other Ways](#other-ways)
- [Usage](#usage)
  - [Table of Canonical Tokens](#table-of-canonical-tokens)
- [Setting Up](#setting-up)
  - [Configuration](#configuration)
    - [Sample configuration file](#sample-configuration-file)
  - [Platforms](#platforms)
    - [Github](#github)
      - [Setting Up](#setting-up-1)
    - [Gitlab](#gitlab)
    - [PivotalTracker](#pivotaltracker)
      - [Setting Up](#setting-up-2)
    - [Trello](#trello)
      - [Selecting Trello Boards](#selecting-trello-boards)
      - [Other references for setting up Trello](#other-references-for-setting-up-trello)
    - [Telegram](#telegram)
- [Development Runbook](#development-runbook)
  - [Getting Started](#getting-started)
  - [Development](#development)
    - [Tools used for development](#tools-used-for-development)
      - [SQLite3 CLI](#sqlite3-cli)
  - [Building](#building)
    - [General build notes](#general-build-notes)
    - [Build notes for Linux](#build-notes-for-linux)
    - [Build notes for MacOS](#build-notes-for-macos)
    - [Build notes for Windows](#build-notes-for-windows)
  - [Releasing](#releasing)
  - [Distribution](#distribution)
  - [References](#references)
- [Licensing](#licensing)



- - -



# Installation

## Via Git Repository

Clone this repository and run `go install ./cmd/dev`.

## Other Ways

Coming soon!



- - -



# Usage

The following is an overview of what can be done:

```sh
# check stuff
dev check software; # checks if required software is installed

# retrieving account information 
dev get account github; # from github
dev get account gitlab; # from gitlab
dev get account pivotaltracker; # from pivotal tracker
dev get account trello; # from trello

# retrieve consumed configuration
dev get config;

# retrieve notifications (todos)
dev get notifications github; # from github
dev get notifications gitlab; # from gitlab
dev get notifications pivotaltracker; # pivotal tracker
dev get notifications trello; # trello

# retrieve your work 
dev get work pivotaltracker; # from pivotal tracker

# initialise persistent database
dev initialise database;

# open stuff
dev open repository; # the repository you're currently in

# start stuff
dev start client; # starts the desktop client helper application
```

## Table of Canonical Tokens

| Concept | Type | Canon | Aliases |
| --- | --- | --- |
| Account | Noun | `account` | `accounts`, `acc`, `accs`, `a` |
| Configuration | Noun | `configuration` | `config`, `conf`, `cf`, `c` |
| Database | Noun | `database` | `db` |
| Gitlab | Noun | `gitlab` | `gl` |
| Github | Noun | `github` | `gh` |
| Network | Noun | `network` | `networks`, `net`, `nets`, `nw` |
| Notifications | Noun | `notifications` | `notification`, `notif`, `notifs`, `n` |
| Repository | Noun | `repository` | `repo`, `rp`, `r` |
| PivotalTracker | Noun | `pivotaltracker` | `pivotal`, `pt` |
| Software | Noun | `software` | `sw`, `s`, `apps` |
| Work | Noun | `work` | `stories`, `tasks`, `tickets`, `w` |
| Check | Verb | `check` | `c`, `verify` |
| Get | Verb | `get` | `retrieve`, `g` |
| Initialise | Verb | `initialise` | `initialize`, `init`, `i` |
| Open | Verb | `open` | `o` |
| Start | Verb | `start` | `st`, `s` |



- - -



# Setting Up

## Configuration

Configuration is done via YAML. By default, `dev` looks for a file at `${HOME}/dev.yaml` and uses that as the base. Next, `dev` will look for a file at `$(pwd)/dev.yaml` and if that is found, merges (add-only) it with the base. If a base configuration is not found, the local `$(pwd)/dev.yaml` will be used. If both are not found and is required by the command you are trying to invoke, an error will be raised.

### Sample configuration file

```yaml
# this defines configurations for dev itself
dev:
  client:
    ## persistent datastore configurations
    database:
      ## this defines the path for the local sqlite3 database
      path: ./dev.db
    ## notifications configurations
    notifications:
      telegram:
        ## the bot's api token as provided by the BotFather
        token: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        ## chat id of your personal account according to your bot
        id: "123456789"
# this defines networks that should be reachable from your machine
## run `dev check networks` to run checks on these
networks:
- name: internet
  check:
    url: https://google.com
- name: internal-vpn
  check:
    url: https://gitlab.internal.com
# this defines software that should be on your machine
## run `dev check software` to run checks on these
software:
- name: golang
  check:
    command: ["go", "version"]
    stdout: ^go version go\d\.\d+(\.\d+)? [a-zA-Z0-9]+\/[a-zA-Z0-9]+$
- name: node
  check:
    command: ["node", "-v"]
    stdout: ^v\d+\.\d+\.\d+$
- name: terraform
  check:
    command: ["terraform", "version"]
    stdout: ^Terraform v\d+\.\d+\.\d+$
- name: terragrunt
  check:
    command: ["terragrunt", "-v"]
    stdout: ^terragrunt version v\d+\.\d+\.\d+$
# this defines platforms that the developer should have access to
platforms:
  ## this defines the pivotal tracker integration
  pivotaltracker:
    accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    projects: # these will be output in your work
    - name: work
      projectID: "XXXXXXX"
    - name: personal
      projectID: "XXXXXXX"
      accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    # ... add as you wish ...
  ## this defines the github integration  
  github:
    accounts:
    - name: personal
      accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    # ... add as you wish ...
  ## this defines the gitlab integration  
  gitlab:
    accounts:
    - name: personal
      description: gitlab cloud
      accessToken: XXXXXXXXXXXXXXXXXXXX
    - name: work-on-prem
      hostname: gitlab.yourdomain.com
      description: office gitlab
      accessToken: XXXXXXXXXXXXXXXXXXXX
      # public exposes this account to the server when it is run,
      # accessToken is ALWAYS redacted but you can share the hostname,
      # when not specified, defaults to false
      public: true
    # ... add as you wish ...
  ## this defines the trello integration  
  trello:
    accessKey: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    accessToken: yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy
    boards: # these will be output in your work
    - shortLink: xXxXxXxX
      list: xxxxxxxx
# this defines links that a developer should have access to
## run `dev goto` to trigger the link search gui in the terminal
links:
  - label: internal vpn endpoint
    categories: ["vpn"]
    url: https://openvpn.yourdomain.com
  - label: official source-of-truth release repository
    categories: ["scm"]
    url: https://gitlab.com/usvc/utils/dev
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


### Trello

Retrieve your `accessToken` from [https://trello.com/app-key](https://trello.com/app-key).

#### Selecting Trello Boards

The `boards` property takes in an array of board shortlinks. You can retrieve a board's shortlink by visiting the board in your browser and extracting it from the URL.

Assuming your board can be found at the URL `https://trello.com/b/xxxxxxxx/lorem-ipsum`, the board shortlink is `xxxxxxxx`.

#### Other references for setting up Trello

- [https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/)


### Telegram

See the [Telegram Integration](./docs/integrations/telegram.md) for more information.

- - -


# Development Runbook

## Getting Started

1. Clone this repository using `git clone git@gitlab.com:usvc/utils/dev.git`
2. Install dependencies using `make deps`
3. Create a local development configuration file at `./dev.yaml` relative to the project root containing [the sample configuration file](#sample-configuration-file)

## Development

### Tools used for development

#### SQLite3 CLI

On Ubuntu install with `sudo apt install sqlite3`.

## Building

As this is a desktop app meant for cross-platform distribution, this gets a little complicated. The instructions assume an Ubuntu build environment.

### General build notes

1. Run `setup_build` to install the required:
   1. `2goarray` for converting a PNG icon into Go code
   2. `rsrc` for compiling Windows application manifests
2. Run `setup_build_linux` if you're building from a linux environment

### Build notes for Linux

1. The Linux assets can be found at `./assets/linux` relative to the project root.

### Build notes for MacOS

1. The MacOS assets can be found at `./assets/macos` relative to the project root.

### Build notes for Windows

1. The Windows assets can be found at `./assets/windows` relative to the project root.
2. Confirm that the `rsrc` binary is available in your path to compile the manifest. You can get this using `make setup_build`

## Releasing

This should be done automatically via the CI pipeline.

## Distribution

> TODO

## References

- On building a static binary file with libraries requiring CGO: [Golang w/SQLite3 + Docker Scratch Image](https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image)
- On executing CLI commands from Go [Advanced command execution in Go with os/exec](https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html)

# Licensing

Code is licensed under the MIT license. [Click here to view the full text](./LICENSE).
