# Telegram Integrations

## Notifications

### Setting up `dev` to notify you via Telegram

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
