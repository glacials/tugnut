# Tugnut

Tugnut is a [LiveSplit][1] parser written in Go.

[1]: http://livesplit.org/

## Building locally

```bash
go install ./...
tugnut
```

Now Tugnut is listening on port 8000 for requests. In another terminal you can give it a LiveSplit file:

```bash
curl localhost:8080/parse/livesplit -F file=@/path/to/livesplit/file.lss
```

and you should receive a JSON interpretation of it back.

### Running tests
```bash
go test ./...
```

## Using Tugnut headlessly

You can use Tugnut's parsing code in your own Go project by importing `github.com/glacials/tugnut/parser`.
