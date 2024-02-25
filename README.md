# Installation
If you want to install the binary for whatever reason, you can run
```
go install github.com/tymcgee/wordbot@v1.0.0
```

# Usage
To play the game, run `wordbot`.

To play around with the bot, there are a few more options:
| Flag            | Meaning                                                                |
|-----------------|------------------------------------------------------------------------|
| -a {string}     | Choose a specific five-letter answer to play against                   |
| -bot            | Choose to play with the bot, instead of interactively                  |
| -first {string} | Choose the first guess for the bot to play with (only works with -bot) |
| -n {integer}    | Choose the number of games for the bot to play (only works with -bot)  |

So for example, if you want to test the bot against the answer `hello`, you can
run `wordbot -bot -a hello`. If you want to run the bot 10000 times using the
first guess `slate` with random answers every time, you would run
`wordbot -bot -first slate -n 10000`.

## License
MIT
