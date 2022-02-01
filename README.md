# Clown bot for discord
[![Actions Status](https://github.com/sasha-sem/clown-bot-go/workflows/Go/badge.svg)](https://github.com/sasha-sem/clown-bot-go/actions)  
A bot for discord to play different games on server

## Avialable games
For now it's only **wordle** but more games coming soon...   
The bot now works only on Russian, I'd be glad if someone helps me to translate it
## Usage
Bot supports slash commands  
Supported commands:
- /привет - bot will say hallo to you
- /wordle new - bot will start new session for a wordle game
- /wordle help - bot will explain you how to play wordle
- /wordle guess *word* - bot tell you if your guess is right

## Deployment
You can deploy bot on the heroku
1. Create an application at https://discordapp.com/developers/applications
2. Setup bot
3. Copy bot Token
4. Use this token as BOT_TOKEN in local.env
5. Deploy on heroku with this template  
[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/sasha-sem/clown-bot-go)
6. Invite your bot to the server
7. Bot needs next permisions:
    - Send Messages
    - Embed Links
    - Read Message History
    - Mention Everyone
    - Add Reactions
8. Enjoy your bot