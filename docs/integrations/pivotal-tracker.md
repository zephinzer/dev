# Pivotal Tracker Integration

- [Pivotal Tracker Integration](#pivotal-tracker-integration)
  - [Setting up API keys](#setting-up-api-keys)
  - [Configuring projects](#configuring-projects)

## Setting up API keys

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

## Configuring projects

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
