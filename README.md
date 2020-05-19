# Dev

[![pipeline status](https://gitlab.com/zephinzer/dev/badges/master/pipeline.svg)](https://gitlab.com/zephinzer/dev/-/commits/master)
![github tag (latest semver)](https://img.shields.io/github/v/tag/zephinzer/dev)

Pimp your developer experience.

This tool exists to improve the day-to-day work experience of software developers. Use this to:

1. Improve onboarding process through defining required software in code
2. Improve day-to-day work experience through defining required network connections in code
3. Improve productivity through moving checking/notification activities to the command line

You can:

- ✅ Check if you have required software - `dev check software`
- ✅ Check if you have required network connectivity - `dev check networks`
- ✅ Check if you have required repositories - `dev check repos`
- ✅ Search for and go to links within your project - `dev goto`
- ✅ Get work assigned to you - `dev get work`
- ✅ Get notifications from developer platforms - `dev get notifications`
- ✅ Get notifications on your desktop
- ✅ Get notifications on your telegram
- ✅ Open the website of repository you're in - `dev open repo`
- ✅ Be l33t by using aliases that become simpler as you use them (example: from `dev get notifications` to `dev get notifs` to `dev g n`)

**Table of Contents**

- [Dev](#dev)
- [Installation](#installation)
  - [Via Git Repository](#via-git-repository)
  - [Other Ways](#other-ways)
- [Usage](#usage)
  - [Overview](#overview)
  - [Table of Canonical Tokens](#table-of-canonical-tokens)
  - [Logs Output](#logs-output)
  - [Configuration](#configuration)
  - [Platforms](#platforms)
    - [Github](#github)
      - [Setting Up](#setting-up)
    - [Gitlab](#gitlab)
    - [PivotalTracker](#pivotaltracker)
    - [Trello](#trello)
    - [Telegram](#telegram)
- [Development Notes](#development-notes)
- [Licensing](#licensing)



- - -



# Installation

## Via Git Repository

Clone this repository and run `make install_local`.

## Other Ways

Coming soon!


- - -


# Usage

## Overview

The following is an overview of what can be done:

```sh
# check stuff
#############
dev check software; # checks if required software is installed
# l33t: dev c sw

dev check networks; # checks if required network access is available
# l33t: dev c nw

dev check repositories; # checks if required repositories are available locally
# l33t: dev c r


# retrieving account information 
################################
dev get github account; # from github
# l33t: dev g gh a

dev get gitlab account; # from gitlab
# l33t: dev g gl a

dev get pivotaltracker account; # from pivotal tracker
# l33t: dev g pt a

dev get trello account; # from trello
# l33t: dev g tr a


# retrieve notifications
########################
dev get gitlab notifications; # from gitlab
# l33t: dev g gl n

dev get pivotaltracker notifications; # pivotal tracker
# l33t: dev g pt n


# retrieve your work 
####################
dev get work pivotaltracker; # from pivotal tracker
# l33t: dev g w pt

dev get config; # retrieve consumed configuration
# l33t: dev g c

# initialise persistent database
################################
dev initialise database;

# initialise telegram notification integration
##############################################
dev initialise telegram notifications;

dev open repository; # opens the website of the repository you're currently in
# l33t: dev o r

dev start client; # starts the desktop client helper application
# l33t: dev s
```



## Table of Canonical Tokens

| Concept | Type | Canon | Aliases |
| --- | --- | --- | --- |
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



## Logs Output

Labelled logs are sent to `stderr` and unlabelled logs are sent to `stdout`. Pipable output is typically sent to `stdout` so you can pipe or stream it to another IO source, logs that indicate inner workings of the application are sent to `stderr`.

- To pipe the **`stdout` logs** only, use the **`>` operator** (`stderr` logs will be sent to terminal).
- To pipe the **`stderr` logs** only, use the **`2>` operator** (`stdout` logs will be sent to terminal).
- To pipe **all logs** use the **`&>` operator**.



## Configuration

See the [documentation on Configuration](./docs/configuration.md) for more information.

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

See the [Pivotal Tracker integration](./docs/integrations/pivotal-tracker.md) for more information.

### Trello

See the [Trello Integration](./docs/integrations/trello.md) for more information.

### Telegram

See the [Telegram Integration](./docs/integrations/telegram.md) for more information.


- - -


# Development Notes

See the [documentation on Development Notes](./docs/development.md) for more information.


- - -


# Licensing

Code is licensed under the MIT license. [Click here to view the full text](./LICENSE).
