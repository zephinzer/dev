<center>
  <img src="./assets/icon/512-dark.png">
</center>

# That Dev Tool

[![build status](https://travis-ci.org/zephinzer/dev.svg?branch=master)](https://travis-ci.org/zephinzer/dev) [![pipeline status](https://gitlab.com/zephinzer/dev/badges/master/pipeline.svg)](https://gitlab.com/zephinzer/dev/-/commits/master) ![github tag (latest semver)](https://img.shields.io/github/v/tag/zephinzer/dev) [![maintainability](https://api.codeclimate.com/v1/badges/c679891cbe62072a9447/maintainability)](https://codeclimate.com/github/zephinzer/dev/maintainability) [![test coverage](https://api.codeclimate.com/v1/badges/c679891cbe62072a9447/test_coverage)](https://codeclimate.com/github/zephinzer/dev/test_coverage)

A CLI tool for improving the developer experience.

> Codebase: [https://gitlab.com/zephinzer/dev](https://gitlab.com/zephinzer/dev)  
> Mirror: [https://github.com/zephinzer/dev](https://github.com/zephinzer/dev)
> 
> Please **file bugs/issues/etc on the GITLAB repository** at [https://gitlab.com/zephinzer/dev/-/issues](https://gitlab.com/zephinzer/dev/-/issues)

**Onboarding**
- Software checks to verify required software is installed
- Network checks to verify required connectivity is established
- Link directory to quickly access project-related URLs

**Development**
- Initialise repositories based on a defined template repository
- Receive notifications on todos/work from common developer platforms (Gitlab, Pivotal Tracker)
- Receive notifications on your desktop, or on Telegram
- Open repository website using the default browser

**Table of Contents**
- [That Dev Tool](#that-dev-tool)
- [Installation](#installation)
  - [Install via Binary](#install-via-binary)
  - [Install via Go](#install-via-go)
  - [Install via Make](#install-via-make)
- [Usage](#usage)
  - [Repository Management](#repository-management)
    - [Configuring repositories](#configuring-repositories)
    - [Checking repositories setup](#checking-repositories-setup)
    - [Cloning all listed repositories](#cloning-all-listed-repositories)
    - [Exporting workspaces from repositories](#exporting-workspaces-from-repositories)
      - [Exporting workspace directly to a file](#exporting-workspace-directly-to-a-file)
      - [Exporting workspace in a different format](#exporting-workspace-in-a-different-format)
  - [Software Management](#software-management)
    - [Configuring softwares](#configuring-softwares)
    - [Checking software setup](#checking-software-setup)
  - [Network Management](#network-management)
    - [Configuring networks](#configuring-networks)
    - [Checking network setup](#checking-network-setup)
  - [Link Directory](#link-directory)
    - [Configuring links](#configuring-links)
    - [Using the link directory](#using-the-link-directory)
  - [Template Repositories](#template-repositories)
    - [Configuring repository templates](#configuring-repository-templates)
    - [Creating a repository from a template](#creating-a-repository-from-a-template)
  - [Platform integrations](#platform-integrations)
    - [Github](#github)
      - [Setting up Github integration](#setting-up-github-integration)
      - [Using the Github integration](#using-the-github-integration)
    - [Gitlab](#gitlab)
      - [Setting up Gitlab integration](#setting-up-gitlab-integration)
      - [Using the Gitlab integration](#using-the-gitlab-integration)
    - [Pivotal Tracker](#pivotal-tracker)
      - [Setting up Pivotal Tracker integration](#setting-up-pivotal-tracker-integration)
      - [Configuring specific Pivotal Tracker projects](#configuring-specific-pivotal-tracker-projects)
      - [Using the Pivotal Tracker integration](#using-the-pivotal-tracker-integration)
    - [Trello](#trello)
      - [Setting up Trello integration](#setting-up-trello-integration)
      - [Configuring Trello boards](#configuring-trello-boards)
      - [Using the Trello integration](#using-the-trello-integration)
  - [Notification integrations](#notification-integrations)
    - [Telegram](#telegram)
      - [Setting up Telegram notifications integration](#setting-up-telegram-notifications-integration)
  - [Other Notes](#other-notes)
    - [On output](#on-output)
    - [On exit codes](#on-exit-codes)
    - [On debugging](#on-debugging)
- [Contributing](#contributing)
  - [Development Flow](#development-flow)
  - [Using the Makefile](#using-the-makefile)
  - [Other resources](#other-resources)
- [Changelog](#changelog)
- [License](#license)

# Installation

## Install via Binary

Releases are available on Github at [https://github.com/zephinzer/dev/releases](https://github.com/zephinzer/dev/releases).

## Install via Go

You can also install it via `go install`:

```sh
go install github.com/zephinzer/dev
```

## Install via Make

Clone this repository and run `make install`:

```sh
git clone git@gitlab.com:zephinzer/dev.git;
cd dev;
make install;
```

Test the installation works by running `dev -v`.

# Usage

`dev` loads its configuration from the home directory and the current working directory from which `dev` is run. A file is considered a configuration file if it matches a pattern of `^.dev(.labels)+.yaml$`. Some examples of valid configuration file names are:

1. `.dev.yaml`
2. `.dev.github.yaml`
3. `.dev.gitlab.yourdomain.com.yaml`
4. `.dev.some1elses.yaml`

## Repository Management

To automate this with some degree of sanity, some strongly-opinionated assumptions are made:

1. Your repositories shall be stored in your home directory using the path `<hostname>/<username>/<path-to-repo>`. For example, this repo will be stored at `/home/${USER}/github.com/zephinzer/dev`
2. You operate your IDE using workspaces to reference the repositories so that the exact repository location on your hard drive does not matter
3. You use SSH keys to authenticate with your source control management platform

### Configuring repositories

Repositories are configured using the `repositories` root level property in the configuration file as such:

```yaml
# ...
repositories:
- description: working repository for the dev tool
  name: dev (do work here)
  url: git@gitlab.com/zephinzer/dev.git
  workspaces: [productivity, development]
- description: public repository for the dev tool
  name: dev (install from here)
  url: git@github.com/zephinzer/dev.git
  workspaces: [productivity, production]
# ...
```

The structure for each repository item can be found at [`./pkg/repository/repository.go`](./pkg/repository/repository.go). In summary the fields are:

| Field         | Required | Default                      | Description                                                                    |
| ------------- | -------- | ---------------------------- | ------------------------------------------------------------------------------ |
| `description` | No       | `""`                         | An arbitrary description displayed in `dev` metadata only                      |
| `name`        | No       | `""`                         | An arbitrary name displayed in `dev` metadata only                             |
| `path`        | No       | `~/${HOSTNAME}/${REPO_PATH}` | The path on your local machine to put this repository                          |
| `url`         | Yes      | `-`                          | The URL to this repository. `dev` will always convert this to an SSH clone URL |
| `workspaces`  | No       | `[]`                         | A list of workspace names that this repository should belong to                |

### Checking repositories setup

To verify that all repositories have been installed locally, use:

```sh
dev check repositories;
```

### Cloning all listed repositories

If a repository is listed in the configuration but is not available locally, it's possible to clone it using this tool using:

```sh
dev install repos;
```

> Known issue #1: You won't be able to clone a bare repository.

> Known issue #2: If you have a long list of repositories, some might fail because too many requests are being made to the server.

### Exporting workspaces from repositories

If you've defined the `workspaces` property for your repositories, you can export repositories related to a workspace by using:

```sh
# get the workspace ${WORKSPACE_NAME}
dev get workspace ${WORKSPACE_NAME};
```

#### Exporting workspace directly to a file

Use the `--output-directory`/`-o` flag to export it directly to the current directory (a file at `./${WORKSPACE_NAME}.code-workspace` will be created):

```sh
# export to the current directory
dev get workspace ${WORKSPACE_NAME} --output-directory .;
```

To export it directly to `/some/path/at` (a file at `/some/path/at/${WORKSPACE_NAME}.code-workspace` will be created):

```sh
# export to the directory at /some/path/at
dev get workspace ${WORKSPACE_NAME} -o /some/path/at;
```

If the file already exists, `dev` will complain. Use the `--overwrite`/`-O` flag to tell it to overwrite the file:

```sh
# export to the directory at /some/path/at, overwriting it if the target file exists
dev get workspace ${WORKSPACE_NAME} -o /some/path/at --overwrite;
```

#### Exporting workspace in a different format

To change the formatting of the workspace, use the `--format`/`-f` flag:

```sh
dev get workspace ${WORKSPACE_NAME} -f vscode;
```

The following workspace formats are supported:

1. Visual Studio Code/Codium (format name: `vscode`)



## Software Management

The software management feature allows developers to check whether required software have been installed and if not, to install them.

### Configuring softwares

Software is configured using the root level `softwares` property in the configuration file as such:

```yaml
# ...
softwares:
- name: brew
  description: used for system dependency installations
  platforms: [macos]
  check:
    command: [brew, --version]
    stdout: ^Homebrew
  install:
    link: https://brew.sh/
- name: node
  description: used for the primary projects
  check:
    command: ["node", "-v"]
    stdout: ^v\d+\.\d+\.\d+$
  install:
    link: https://github.com/nvm-sh/nvm#installing-and-updating
- name: terraform
  description: used for bringing up our infrastructure
  check:
    command: ["terraform", "version"]
    stdout: ^Terraform v\d+\.\d+\.\d+$
  install:
    link: https://www.terraform.io/downloads.html
# ...
```

The structure for each software item is detailed in the [`./pkg/software` directory](./pkg/software) and a summary follows:

| Field            | Required | Default | Description                                                                                                                                                                                                  |
| ---------------- | -------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `name`           | No       | `""`    | An arbitrary name displayed in `dev` outputs when metadata is required                                                                                                                                       |
| `check`          | Yes      | -       | Declarative instructions describing how to check for the software's presence                                                                                                                                 |
| `check.command`  | Yes      | -       | A list of strings that form the command to check for the software. For example, if you run `go version` in your CLI to check for Go, the list will look like `["go", "version"]`                             |
| `check.stdout`   | No       | `""`    | A regex-compatible string to match with the command's output on `stdout`                                                                                                                                     |
| `check.stderr`   | No       | `""`    | A regex-compatible string to match with the command's output on `stderr`                                                                                                                                     |
| `check.exitCode` | No       | `0`     | An integer exit code to match with the command's exit code                                                                                                                                                   |
| `description`    | No       | `""`    | An arbitrary description of how the software is used/why it's needed displayed in `dev` outputs when metadata is required                                                                                    |
| `install`        | No       | -       | Declarative instructions describing how to install the software if it's not found                                                                                                                            |
| `install.link`   | No       | `""`    | A link to direct users to if the software is not found installed on the current machine                                                                                                                      |
| `platforms`      | No       | -       | A list of platform strings that define which operating systems this check is valid for. Valid lists are subsets of `{linux, windows, macos}`. When not specified, does not check for platform compatibility. |

### Checking software setup

To run checks on the software available locally:

```sh
dev check software;
```

## Network Management

Traditionally, teams/organisations typically have internal networks (intranets) for developers to access where privately owned code is pulled from/pushed to. This feature assists the developer to see what networks exist and to verify they have access to the networks they should have access to.

### Configuring networks

Networks are configured using the root level `networks` property in the configuration file as such:

```yaml
# ...
networks:
- name: internet
  check:
    url: https://google.com
- name: internal-vpn
  registrationUrl: https://openvpn.internal.com
  check:
    url: https://gitlab.internal.com
# ...
```

The structure of each network can be found in the [`./pkg/network` directory](./pkg/network), a summary is as follows:

| Field                | Required | Default   | Description                                                                                                                                    |
| -------------------- | -------- | --------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `name`               | No       | `""`      | An arbitrary name for the network to be displayed in `dev` metadata when required                                                              |
| `registrationUrl`    | No       | `""`      | A URL for users to request access to this network if applicable                                                                                |
| `check`              | Yes      | -         | Declarative instructions on what to check for                                                                                                  |
| `check.url`          | Yes      | -         | A URL to ping that should be accessible if network connectivity to this network has been established                                           |
| `check.method`       | No       | `"GET"`   | The HTTP method to use to ping the provided `check.url`                                                                                        |
| `check.statusCode`   | No       | `^2\d\d$` | The expected HTTP status code of the response to the ping                                                                                      |
| `check.headers`      | No       | `{}`      | A dictionary of expected headers                                                                                                               |
| `check.responseBody` | No       | `""`      | A regex-compatible string that the response body should be expected to match with. Does not apply if not specified/an empty string is provided |

### Checking network setup

To run a check on your required network connectivity:

```sh
dev check networks;
```

## Link Directory

This feature assists with awareness of team resources by providing a list of resources that a developer can use to explore the team's resources.

### Configuring links

The link directory is configured using the root level `links` property in the configuration file as such:

```yaml
links:
- label: official source-of-truth release repository
  categories: [scm]
  url: https://gitlab.com/zephinzer/dev
- label: dev tool build pipeline
  categories: [cicd]
  url: https://gitlab.com/zephinzer/dev/pipelines
- label: dev tool release pipeline
  categories: [cicd, release]
  url: https://travis-ci.org/github/zephinzer/dev/
- label: dev tool code quality checks
  categories: [cicd]
  url: https://codeclimate.com/github/zephinzer/dev
- label: dev tool releases
  categories: [cicd, release]
  url: https://github.com/zephinzer/dev/releases
```

The structure for each link object can be found in the [`./internal/link` directory](./internal/link) and a summary follows.

| Field        | Required | Default | Description                                                      |
| ------------ | -------- | ------- | ---------------------------------------------------------------- |
| `label`      | No       | `""`    | An arbitrary label for this link to be included in link metadata |
| `categories` | No       | `[]`    | A hashtag-based way of searching to be included in link metadata |
| `url`        | Yes      | -       | The URL to open if this link is selected                         |

### Using the link directory

To activate the link directory from the terminal, run:

```sh
dev goto;
```

## Template Repositories

Being able to create new repositories based on a pre-defined list of templates makes it easy to do the right thing for developers by providing an easy way to kickstart a new service in a organisation/team-approved manner.

### Configuring repository templates

Repository templates can be configured using the configuration key at `dev.repository/.templates` which should be an array of template objects as such:

```yaml
# ...
dev:
  # ...
  repository:
    templates:
    - name: test for bare repository
      url: https://github.com/zephinzer/template-bare
    - name: golang repository
      url: https://gitlab.com/zephinzer/template-go-package
    # ...
  # ...
# ...
```

The structure for each template object can be found at [`./internal/config/dev.go`](./internal/config/dev.go) and a summary follows.

| Field  | Required | Default | Description                                                                  |
| ------ | -------- | ------- | ---------------------------------------------------------------------------- |
| `name` | NO       | `""`    | An arbitrary label used to identify this repository template                 |
| `url`  | YES      | -       | One of the template repository's HTTP URL, HTTPS clone URL, or SSH clone URL |


### Creating a repository from a template

```sh
dev init repo ./path/to/new/repo;

## output:
## > choose a repository template to use
## > 1. ...
## > 2. ...
##
## > your selection (enter 0 to skip):
```

## Platform integrations

### Github

The Github integration is not refined for use yet, try it at your own risk!

#### Setting up Github integration

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

#### Using the Github integration

```sh
dev get github account;
dev get github notifs;
```

### Gitlab

The Gitlab integration is not refined for use yet, try it at your own risk!

#### Setting up Gitlab integration

Retrieve your `accessToken` by generating a new personal access token from [https://gitlab.com/profile/personal_access_tokens](https://gitlab.com/profile/personal_access_tokens). You'll need the following scopes:

- api
- read_api

> If you're using an on-premise Gitlab, change `gitlab.com` to your Gitlab's hostname


#### Using the Gitlab integration

```sh
dev get gitlab account;
dev get gitlab notifs;
```

### Pivotal Tracker

#### Setting up Pivotal Tracker integration

Retrieve your `accessToken` from [https://www.pivotaltracker.com/profile](https://www.pivotaltracker.com/profile) and enter it as the `accessToken` as a property at `platforms.pivotaltracker`

Example:

```yaml
# ...
platforms:
  # ...
  pivotaltracker:
    accessToken: ...
# ...
```

#### Configuring specific Pivotal Tracker projects

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

The Trello integration is not refined for use yet, try it at your own risk!

#### Setting up Trello integration

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

#### Configuring Trello boards

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

## Other Notes

### On output

**All output that's prefixed with a log level is being sent to `stderr`.**

To display just `stderr` output, append `2>/dev/null` to your command.

**All output that appears as-is without a log level is being sent to `stdout`**

To display just `stdout` output, append `1>/dev/null` to your command.

### On exit codes

Exit codes are bitmasks of error types which are docuemnted at [`./internal/constants/exit_codes.go`](./internal/constants/exit_codes.go).

| Type             | Value     | Indication                                                                                                                                                              |
| ---------------- | --------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| OK               | `0`       | Success                                                                                                                                                                 |
| System           | `1`       | An error occurred at the system level                                                                                                                                   |
| User             | `2`       | An error occurred because of user actions                                                                                                                               |
| Input            | `4`       | An error occurred because of user input                                                                                                                                 |
| Configuration    | `8`       | An error occurred because of the consumed configuration                                                                                                                 |
| Application      | `16`      | There's likely a bug                                                                                                                                                    |
| Validation       | `32`      | Some expected values seem off                                                                                                                                           |
| Number of errors | `128-255` | When such an exit code is encountered it's a count of failed iterations from a requested operation, take `$((256 - $?))` to get the number of errors from the operation |

### On debugging

Two global flags are made available to improve debuggability by increasing the amount of logs.

- `--debug`: display DEBUG level logs (this prints start and end messages, and input/output values)
- `--trace`: display TRACE level logs (this prints all sorts of nonsense but could be useful sometimes)


# Contributing

## Development Flow

1. Clone the repository at [https://gitlab.com/zephinzer/dev](https://gitlab.com/zephinzer/dev) (NOTE: not Github, use Gitlab).
2. Make a fork of it
3. Raise an issue at [https://gitlab.com/zephinzer/dev/-/issues](https://gitlab.com/zephinzer/dev/-/issues)
4. Link to the relevant issue in [Github issues](https://github.com/zephinzer/dev/issues) if there's a linked issue
5. Make your changes and raise a Merge Request
6. Once tests pass and the MR is merged to `master`, the repository will be synced to Github and will be automatically released

## Using the Makefile

5. Run `make deps` to retrieve Go dependencies
6. Run `make setup_build` to retrieve system dependencies
7. Run `make build` to run a test build with caching
8. Run `make build_production` to run a full non-cached build
7. Run `make build_static` to run a test build with static linking and caching
8. Run `make build_static_production` to run a full non-cached build with static linking
9.  Run `make test` to run Go test suites
10. Run `make compress` to test whether compression works
11. Run `make image` to test the Docker image build
12. Run `make test_image` to run tests on the built Docker image

## Other resources

- [The Dockerhub page](https://hub.docker.com/r/zephinzer/dev)
- [The Github page](https://github.com/zephinzer/dev)
- [The Gitlab page](https://gitlab.com/zephinzer/dev)
- [The webpage](https://getthat.dev)

# Changelog

| Version | Breaking | Description                                                                                                                                                                                                                      |
| ------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| v0.1.24 | NO | Added error handling for triggering of `dev add -` commands without an existing configuration file |
| v0.1.18 | NO       | Added command to initialising a repository using a template (`dev init repo <path>`)                                                                                                                                             |
| v0.1.7  | NO       | Added descriptions for `dev check software`                                                                                                                                                                                      |
| v0.1.6  | NO       | Made repository selection deterministic when using `dev add repo`                                                                                                                                                                |
| v0.1.4  | NO       | Removal of unused fields using the `omitempty` struct tag for networks, softwares, links, and repositories, fixed bug where the `dev` configuration wasn't being correctly merged, refined Pivotal Tracker notification messages |
| v0.1.0  | YES      | Changing of configuration filename from `dev.yaml` to `.dev.yaml`                                                                                                                                                                |

# License

Code is licensed under the [MIT license (click to view the full text)](./LICENSE).

