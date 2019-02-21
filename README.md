# :moneybag: Galaxy Merchant Trading

Your day-to-day intergalactic trading conversion tool. Calculate your tradings instantly and get on with life because time is precious and life is too short to spend on these trivial tasks. :grin:

## Use

Build normally:

```sh
$ go build && ./galaxy-merchant
```

Or run directly:

```sh
$ go run main.go
```

And input your trading queries:

```
glob is I
prok is V
how much is glob prok ?
```

To exit, input `exit` as query, or press `ctrl` + `c`.

## Query Pattern

Supported query patterns:

- `X is Y`

    - `X` - Intergalactic word number to replace / represent roman numeral
    - `Y` - Represented single roman numeral character

    Example: `glob is I`

- `I J is K Credits`

    - `I` - Intergalactic words number separated by space
    - `J` - Trading material
    - `K` - Decimal credit value of said material

    Example: `glob prok Silver is 68 Credits`

- `how much is I ?`

    - `I` - Intergalactic words number separated by space

    Example: `how much is glob prok ?`

- `how many Credits is I J ?`

    - `I` - Intergalactic words number separated by space
    - `J` - Trading material

    Example: `how many Credits is glob prok Silver ?`

## TODO

- [ ] Document system design
- [ ] Improve query pattern matching ?
