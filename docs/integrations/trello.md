# Trello Integration

## Documentation

Documentation for the Trello API can be found at https://developer.atlassian.com/cloud/trello/rest.

## Experimentation

### With cURL

An example call to retrieve details about the owner of the access key and token:

```sh
export TRELLO_KEY=__paste_key_here__;
export TRELLO_TOKEN=__paste_token_here__;
curl 'https://api.trello.com/1/members/me?key=${TRELLO_KEY}&token=${TRELLO_TOKEN}';
```