# gogenshin

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white) ![Last Commit](https://img.shields.io/github/last-commit/randomchaffee/gogenshin)

A simple Discord bot written in Go.

## what it can do

- responds to `!ping` with `Pong!`

## requirements

- Go 1.26.5 or any compatible version
- A Discord application and the bot token
- If you already have experience creating discord bots, the proceeding processes should be straightforward

dependencies:
- [DiscordGo](https://github.com/bwmarrin/discordgo).

## run

install deps and start the bot:

```bash
go mod download
go run .
```

Once the bot is running, send `!ping` in a channel where it can read and send
messages.

## configuration

right now the application reads `DISCORD_BOT_TOKEN` directly from the environment; it does not load `.env` files automatically.