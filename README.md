# Caboose

A Redis clone built in Go, because I had nothing better to do this afternoon. Named after [Captain Michael J. Caboose](https://rvb.fandom.com/wiki/Michael_J._Caboose).

## How to use it?

Please don't. It's a Proof of Concept.

Run `go run main.go` in a terminal window, it'll spawn a (_somewhat_) Redis-compatible server on `localhost:6900`. Send commands using `redis-cli`, like in the screenshot:

![sc](https://github.com/aziflaj/caboose/assets/5219775/c762a28b-f900-4cd3-9179-f908484c52ff)

## RESP
TL;DR version of [REdis Serialization Protocol](https://redis.io/docs/reference/protocol-spec/).

> **Disclaimer:**
> 
> I was lazy to read through the whole RESP specs, but I think I am implementing RESP 2.0

- RESP is a req/res protocol, with binary requests and responses (but they're ASCII-compatible)
- All req/res are terminated by CRLF (`\r\n`)
- The 1st byte of the req/res payload defines the data type:
  - Simple (non-binary) Strings start with `+`: e.g. `+OK\r\n`
  - Bulk strings (whatever that is) start with `$`: e.g. `$5\r\nhowdy\r\n` - That initial 5 represents the byte length of the encoded string
  - Errors start with `-`: e.g. `-ERR bruh whachu doin\r\n`
  - Integers start with `:`: e.g. `:420` or `:-69`
  - Arrays start with `*`: each array item should be encoded with the data type as well, so:
    - String Array: `*2\r\n$5\r\nhowdy\r\n$4\r\nyall\r\n\r\n` (notice the initial `*2`, the length, followed by `$5\r\nhowdy` and the `$4\r\nyall` divided by CRLF)
    - Int array: `*2\r\n:420\r\n:-69\r\n`
    - Mixed data types: `*2\r\n:69\r\n$4\r\nwink\r\n` (I'm joking now, but I can see how parsing this will be a pain in my assholes)
  - Nulls start with `_`: e.g. `_\r\n`
  - Bools start with `#`
  - ~~Doubles start with `,` (now you're just making things up)~~ I decided not to bother; this is a proof of concept
  - Maps start with `%` (Aight, I'm not gonna implement all these now, am I?)

