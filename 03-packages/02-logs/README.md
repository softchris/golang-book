<https://blog.logrocket.com/five-structured-logging-packages-for-go/>

# Logging

Once you started writing code you realize quite early that you have a need to print things to the screen as well as sometimes to a file or even a log service. What you want to say is usually what type of logging you want to do.

## Reasons to log

There are many reasons to log, here's some reasons:

- **Information**, there might be case where you want provide some type of information that could be of use to the one using the program.
- **Success**. A success message is a little more that just information, it indicates that you succeeded with something.
- **Warning**. When you have a warning something happened that you should be aware of. It's usually not serious enough to shut down the app but it should make you vigilant, it could be that memory is running low for example.
- **Error**. When you get an error you tend to end up in a state where it's no longer a wise choice to continue. 
- **Performance**. It's common to measure how long something takes, for the sake of improving things this information can be useful.
- **Other**. There are also other reasons why you would log something, usually that's connected to your business. 

## What to log

The general rule is the more you can log the better. Especially if it's an error you want to fix you might want to log things like:

- When it happened
- What happened
- Specific error info

For every case you want to log, have a log at how the log message will be used, will a team be logging through these logs, what would help them. See if you can interview someone on that team.

## Using `log`

identifying where to use it

### Standard log `Println()`

### `Fatal`

### Log to a file

## Standard logging or third-party