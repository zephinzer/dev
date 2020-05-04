# Telegram Integrations

## Notifications

### Setting up

To receive notifications via Telegram, we'll need to setup a new Telegram bot. You can do this by talking to [The BotFather](https://t.me/BotFather) using `/newbot`. Give your bot a logical name (eg. "My Notification Bot") followed by a bot username (eg. `my_notification_bot`) and you'll receive an access token in a message that looks like the following:

```
Done! Congratulations on your new bot. You will find it at t.me/your_notification_bot. You can now add a description, about section and profile picture for your bot, see /help for a list of commands. By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.

Use this token to access the HTTP API:
1234556789:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
Keep your token secure and store it safely, it can be used by anyone to control your bot.

For a description of the Bot API, see this page: https://core.telegram.org/bots/api
```

The API token should be placed in the configuration YAML at property `dev.client.notifications.telegram.token`. To register yourself, run `dev init telegram notifications` or `dev i tg n` if you're feeling l33t.

You should see logs indicating your bot's name and informing you that `dev` is starting a telegram bot controller. Click on the link that `dev` outputs and hit the **Start** button which should result in a `/start` being sent to the bot.

You should see a YAML snippet containing the configuration you need to merge with your existing configuration.
