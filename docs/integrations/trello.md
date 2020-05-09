# Trello Integration

- [Trello Integration](#trello-integration)
  - [Setting up API keys](#setting-up-api-keys)
  - [Configuring boards](#configuring-boards)
  - [More resources](#more-resources)
    - [API Documentation](#api-documentation)
    - [Experimentation with cURL](#experimentation-with-curl)
    - [Other links](#other-links)

## Setting up API keys

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

## Configuring boards

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

## More resources

### API Documentation

Documentation for the Trello API can be found at https://developer.atlassian.com/cloud/trello/rest.

### Experimentation with cURL

An example call to retrieve details about the owner of the access key and token:

```sh
export TRELLO_KEY=__paste_access_key_here__;
export TRELLO_TOKEN=__paste_access_token_here__;
curl 'https://api.trello.com/1/members/me?key=${TRELLO_KEY}&token=${TRELLO_TOKEN}';
```

### Other links

- [https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/)
