
# Usage

```sh
# retrieve consumed configuration
dev get config;

# retrieve notifications from pivotal tracker
dev get notifs pivotaltracker;

# retrieve your work from pivotal tracker
dev get work pivotaltracker;
```

# Configuration

## Sample configuration file

```yaml
# this defines software that should be on the developer's machine
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
      accessToken: XXXXXXX-XXXXXXXXXXXX
    - name: work-on-prem
      hostname: gitlab.yourdomain.com
      accessToken: XXXXXXX-XXXXXXXXXXXX
```

# References

- [Golang w/SQLite3 + Docker Scratch Image](https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image)
  
# Licensing

MIT
