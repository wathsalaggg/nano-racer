# Contributing to Nano-Racer

Thanks for taking a look. Contributions of any size are welcome, and so are
people who are new to Go or new to open source. This is a small project, so
it is a good place to make a first contribution.

## Ways to help

- **Pick an issue.** The ones labelled
  [`good first issue`](https://github.com/wathsalaggg/nano-racer/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
  are small and come with hints.
- **Add an example.** A new race in the `examples` folder is the easiest
  change to make.
- **Report a problem.** If something breaks or looks wrong,
  [open an issue](https://github.com/wathsalaggg/nano-racer/issues/new) and
  say what you ran, what you expected and what happened.
- **Suggest an idea.** Open an issue or start a
  [discussion](https://github.com/wathsalaggg/nano-racer/discussions).

If you plan a larger change, please open an issue first so we can agree on
the approach before you spend time on it.

## Setting up

You need [Go](https://go.dev/dl/) 1.22 or newer. There are no other
dependencies.

```
git clone https://github.com/wathsalaggg/nano-racer.git
cd nano-racer
go test
go run ./examples/basic
```

## Making a change

1. Fork the repository on GitHub and clone your fork.
2. Create a branch: `git checkout -b my-change`
3. Make the change.
4. Run the checks below.
5. Commit, push the branch to your fork, and open a pull request.

## Checks

GitHub runs these on every pull request. Run them yourself first:

```
gofmt -l .           # should print nothing
go vet ./...
go test -race ./...
go build ./...
```

`gofmt -w .` fixes formatting for you.

## Style

- **Comments are written for beginners.** This project is meant to be
  readable by someone who has never used Go. When your change uses a Go
  feature the code has not used before, add a short comment saying what it
  does. Plain words are better than jargon.
- **No dependencies.** Everything uses Go's standard library. Please keep it
  that way unless an issue has agreed otherwise.
- **Small commits with a reason.** One idea per commit. In the message, say
  what changed and why.
- **Add a test** when you change how `Run` or `Print` behaves.

## Where things are

| File | What it does |
|---|---|
| `racer.go` | times a single function |
| `race.go` | lanes, laps, and running them with goroutines |
| `bars.go` | the live progress bars |
| `report.go` | the results table |
| `race_test.go` | the tests |
| `examples/` | small programs that use the library |

## Licence

By contributing you agree that your work is released under the project's
[MIT licence](LICENSE).
