# That Dev Tool

[![pipeline status](https://gitlab.com/zephinzer/dev/badges/master/pipeline.svg)](https://gitlab.com/zephinzer/dev/-/commits/master)
![github tag (latest semver)](https://img.shields.io/github/v/tag/zephinzer/dev)
[![maintainability](https://api.codeclimate.com/v1/badges/c679891cbe62072a9447/maintainability)](https://codeclimate.com/github/zephinzer/dev/maintainability)
[![test coverage](https://api.codeclimate.com/v1/badges/c679891cbe62072a9447/test_coverage)](https://codeclimate.com/github/zephinzer/dev/test_coverage)

A CLI tool for improving the developer experience.

**Onboarding**
- Software checks to verify required software is installed
- Network checks to verify required connectivity is established
- Link directory to quickly access project-related URLs

**Development**
- Receive notifications on todos/work from common developer platforms (Gitlab, Pivotal Tracker)
- Receive notifications on your desktop, or on Telegram
- Open repository website using the default browser

**Table of Contents**
- [That Dev Tool](#that-dev-tool)
- [Installation](#installation)
- [Usage](#usage)
  - [Software checks](#software-checks)
    - [Setting up software checks](#setting-up-software-checks)
    - [Using software checks](#using-software-checks)
  - [Network checks](#network-checks)
    - [Setting up network checks](#setting-up-network-checks)
    - [Using network checks](#using-network-checks)
  - [Link directory](#link-directory)
    - [Setting up the link directory](#setting-up-the-link-directory)
    - [Using the link directory](#using-the-link-directory)
  - [Platform integrations](#platform-integrations)
    - [Github](#github)
      - [Setting up Github integration](#setting-up-github-integration)
      - [Using the Github integration](#using-the-github-integration)
    - [Gitlab](#gitlab)
      - [Setting up Gitlab integration](#setting-up-gitlab-integration)
      - [Using the Gitlab integration](#using-the-gitlab-integration)
    - [Pivotal Tracker](#pivotal-tracker)
      - [Setting up Pivotal Tracker integration](#setting-up-pivotal-tracker-integration)
        - [Setting up Pivotal Tracker API keys](#setting-up-pivotal-tracker-api-keys)
        - [Configuring specific Pivotal Tracker projects](#configuring-specific-pivotal-tracker-projects)
      - [Using the Pivotal Tracker integration](#using-the-pivotal-tracker-integration)
    - [Trello](#trello)
      - [Setting up Trello integration](#setting-up-trello-integration)
        - [Setting up Trello API credentials](#setting-up-trello-api-credentials)
        - [Configuring boards](#configuring-boards)
      - [Using the Trello integration](#using-the-trello-integration)
  - [Notification integrations](#notification-integrations)
    - [Telegram](#telegram)
      - [Setting up Telegram notifications integration](#setting-up-telegram-notifications-integration)
  - [Exit Codes](#exit-codes)
  - [Debugging](#debugging)
- [Contributing](#contributing)
- [Changelog](#changelog)
- [License](#license)

# Installation

Releases are available on Github at [https://github.com/zephinzer/dev/releases](https://github.com/zephinzer/dev/releases).

You can also install it via `go install`:

```sh
go install github.com/zephinzer/dev
```

Test the installation works by running `dev -v`.

# Usage

`dev` loads its configuration from the home directory and the current working directory from which `dev` is run. A file is considered a configuration file if it matches a pattern of `^.dev(.labels)+.yaml$`. Some examples of valid configuration file names are:

1. `.dev.yaml`
2. `.dev.github.yaml`
3. `.dev.gitlab.yourdomain.com.yaml`
4. `.dev.some1elses.yaml`

## Software checks

### Setting up software checks

Add a root level `softwares` property. A working example looks like:

```yaml
softwares:
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
```

### Using software checks

To run checks on the software available locally:

```sh
dev check software;
```

## Network checks

### Setting up network checks

Add a root level `networks` property. A working example looks like:

```yaml
networks:
- name: internet
  check:
    url: https://google.com
- name: internal-vpn
  check:
    url: https://gitlab.internal.com
```

### Using network checks

To run a check on your network connectivity:

```sh
dev check networks;
```

## Link directory

### Setting up the link directory

Add a root level `links` property. A working example looks like:

```yaml
links:
- label: official source-of-truth release repository
  categories: ["scm"]
  url: https://gitlab.com/zephinzer/dev
- label: dev tool build pipeline
  categories: ["cicd"]
  url: https://gitlab.com/zephinzer/dev/pipelines
- label: dev tool release pipeline
  categories: ["cicd", "release"]
  url: https://travis-ci.org/github/zephinzer/dev/
- label: dev tool code quality checks
  categories: ["cicd"]
  url: https://codeclimate.com/github/zephinzer/dev
- label: dev tool releases
  categories: ["cicd", "release]
  url: https://github.com/zephinzer/dev/releases
```

### Using the link directory

To activate the link directory from the terminal, run:

```sh
dev goto;
```

## Platform integrations

### Github

#### Setting up Github integration

#### Using the Github integration

```sh
dev get github account;
dev get github notifs;
```

### Gitlab

#### Setting up Gitlab integration

#### Using the Gitlab integration

```sh
dev get gitlab account;
dev get gitlab notifs;
```

### Pivotal Tracker

#### Setting up Pivotal Tracker integration

##### Setting up Pivotal Tracker API keys

1. Retrieve your `accessToken` from [https://www.pivotaltracker.com/profile](https://www.pivotaltracker.com/profile).
2. Enter the `accessToken` as a property at `platforms.pivotaltracker`

Example:

```yaml
# ...
platforms:
  # ...
  pivotaltracker:
    accessToken: ...
# ...
```

##### Configuring specific Pivotal Tracker projects

1. Navigate to the project you want to receive work/notifications from
2. From the URL, extract the project ID (assuming a URL like `https://www.pivotaltracker.com/n/projects/1234567`, the project ID is `1234567`)
3. Add an array item to the property at `platforms.pivotaltracker.projects`. The array item has a structure containing 3 properties:
   1. `name`: an arbitrary label you can use to identify the project
   2. `projectID` the project ID as retrieved from above **as a string** (surround with `"double quotes"` to be sure)

Example:

```yaml
# ...
platforms:
  # ...
  pivotaltracker:
    accessToken: ...
    projects:
    - name: (some arbitrary name you use to identify your project)
      projectID: "1234567"
    # ... other projects ...
# ...
```

#### Using the Pivotal Tracker integration

```sh
dev get pivotal account;
dev get pivotal notifs;
```

### Trello

#### Setting up Trello integration

##### Setting up Trello API credentials

1. Retrieve your `accessKey` from [https://trello.com/app-key](https://trello.com/app-key).
2. Generate your `accessToken` from the link to **Token** from the above link
3. Enter the `accessKey` and `accessToken` as properties at `platforms.trello`

Example:

```yaml
# ...
platforms:
  # ...
  trello:
    accessKey: ...
    accessToken: ...
# ...
```

##### Configuring boards

The `boards` property at `platforms.trello.boards` takes in an array of board shortlinks. You can retrieve a board's shortlink by visiting the board in your browser and extracting it from the URL.

Assuming your board can be found at the URL `https://trello.com/b/xxxxxxxx/lorem-ipsum`, the board shortlink is `xxxxxxxx`.

Example:

```yaml
# ...
platforms:
  # ...
  trello:
    # ... api keys ...
    boards:
    - xxxxxxxx
# ...
```

#### Using the Trello integration

```sh
dev get trello account;
dev get trello notifs;
```

## Notification integrations

### Telegram

`dev` is able to notify you via Telegram if the integration has been set up.

#### Setting up Telegram notifications integration

1. Talk to [The BotFather](https://t.me/BotFather) using `/start` if you haven't talked to it before.
2. Send the `/newbot` command to create a new bot. Give your bot a logical name (eg. "My Notification Bot") followed by a bot username (eg. `my_notification_bot`).
3. You should receive an access token of the form `1234556789:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`, add the access token at the property `dev.client.notifications.telegram.token`.
4. Run `dev init telegram notifs`
5. Talk to your bot and you should see a chat ID appear in `dev`'s logs
6. Copy and paste this chat ID to the property `dev.client.notifications.telegram.id` **as a string** (surround the ID with quotes)
7. To test the integration, ensure your chat ID is in the configuration, and run `dev debug notifications` and confirm you receive a notification in Telegram

Your configuration should look like:

```yaml
# ... other properties ...
dev:
  # ... other properties ...
  notifications:
    telegram:
      token: 1234556789:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
      id: "987654321"
# ... other properties ...
```

## Exit Codes

Exit codes are bitmasks of error types which are docuemnted at [`./internal/constants/exit_codes.go`](./internal/constants/exit_codes.go).

| Type | Value | Indication |
| --- | --- | --- |
| OK | `0` | Success |
| System | `1` | An error occurred at the system level |
| User | `2` | An error occurred because of user actions |
| Input | `4` | An error occurred because of user input |
| Configuration | `8` | An error occurred because of the consumed configuration |
| Application | `16` | There's likely a bug |
| Validation | `32` | Some expected values seem off |

## Debugging

Two global flags are made available to improve debuggability by increasing the amount of logs.

- `--debug`: display DEBUG level logs (this prints start and end messages, and input/output values)
- `--trace`: display TRACE level logs (this prints all sorts of nonsense but could be useful sometimes)


# Contributing

Coming soon!

# Changelog

| Version | Breaking | Description |
| --- | --- | --- |
| v0.1.7 | NO | Added descriptions for `dev check software` |
| v0.1.6 | NO | Made repository selection deterministic when using `dev add repo` |
| v0.1.4 | NO | Removal of unused fields using the `omitempty` struct tag for networks, softwares, links, and repositories, fixed bug where the `dev` configuration wasn't being correctly merged, refined Pivotal Tracker notification messages |
| v0.1.0 | YES | Changing of configuration filename from `dev.yaml` to `.dev.yaml` |

# License

Code is licensed under the MIT license. [Click here to view the full text](./LICENSE).

