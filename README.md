# Structure
There are two main modules here: `game` and `bot`.
## Game
The `game` module is responsible for the logic of the game and deciding how
guesses are evaluated, and includes a CLI implementation of playing the game.

The function for running the game accepts a custom "guess decider" function as
an argument which allows for any number of different ways to decide and input
your guesses.

## Bot
The `bot` module is responsible for the creation and usage of a bot to play the
game, filtering down the list of valid guesses based on feedback from the
previous guess.

The function for running the bot game is a little less customizable, but it
does accept an argument for the "filter method", so you can implement your own
method for filtering down the guesses based on previous information. A drawback
for the current implementation is that the previous information available here
does not contain ALL game information up to that point -- in particular, yellow
guesses are wiped after each guess. A smarter bot than mine might like to use
that information, but as of now, it is hidden.

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
