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

### Using software checks

```sh
dev check software;
```

## Network checks

### Setting up network checks

### Using network checks

```sh
dev check networks;
```

## Link directory

### Setting up the link directory

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

#### Using the Pivotal Tracker integration

```sh
dev get pivotal account;
dev get pivotal notifs;
```

### Trello

#### Setting up Trello integration

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
