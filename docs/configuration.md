# Configuration

Configuration is done via YAML. By default, `dev` looks for a file at `${HOME}/dev.yaml` and uses that as the base. Next, `dev` will look for a file at `$(pwd)/dev.yaml` and if that is found, merges (add-only) it with the base. If a base configuration is not found, the local `$(pwd)/dev.yaml` will be used. If both are not found and is required by the command you are trying to invoke, an error will be raised.

## Application

Working example:

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
```

## Networks

Working example:

```yaml
# this defines networks that should be reachable from your machine
## run `dev check networks` to run checks on these
networks:
- name: internet
  check:
    url: https://google.com
- name: internal-vpn
  check:
    url: https://gitlab.internal.com
```

## Softwares

Working example:

```yaml
# this defines software that should be on your machine
## run `dev check software` to run checks on these
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

## Repositories

Working example:

```yaml
# this defines repositories that a developer should have on their machines
# run `dev check repos` to verify their integrity
repositories:
- name: frontend
  description: our website
  cloneURL: git@gitlab.com:yourns/frontend.git
  workspaces: ["dev"]
- name: api
  description: our api
  cloneURL: git@gitlab.com:yourns/api.git
  workspaces: ["dev"]
- name: backend
  description: our backend stuff
  cloneURL: git@gitlab.com:yourns/backend.git
  workspaces: ["dev", "ops"]
```

## Links

Working example:

```yaml
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

## Platform Integrations

### Github

Working example:

```yaml
platforms:
  ## this defines the github integration  
  github:
    accounts:
    - name: personal
      accessToken: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

### Gitlab

Working example:

```yaml
platforms:
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
```

### Pivotal Tracker

Working example:

```yaml
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
```

### Trello

Working example:

```yaml
platforms:
  ## this defines the trello integration  
  trello:
    accessKey: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    accessToken: yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy
    boards: # these will be output in your work
    - shortLink: xXxXxXxX
      list: xxxxxxxx
```


