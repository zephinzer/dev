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

# Installation

Releases are available on Github at [https://github.com/zephinzer/dev/releases](https://github.com/zephinzer/dev/releases).

You can also install it via `go install`:

```sh
go install github.com/zephinzer/dev
```

Test the installation works by running `dev -v`.

# Usage

`dev` loads its configuration from `~/dev.yaml` where `~` is your home directory. It also searches for a local `./dev.yaml` in your current working directory and if found, merges it with the global configuration.

## Software checks

### Setting up software checks

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

```sh
dev check software;
```

## Network checks

### Setting up network checks

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

```sh
dev check networks;
```

## Link directory

### Setting up the link directory

```yaml
links:
- label: internal vpn endpoint
  categories: ["vpn"]
  url: https://openvpn.yourdomain.com
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

# Contributing

# License

Code is licensed under the MIT license. [Click here to view the full text](./LICENSE).
