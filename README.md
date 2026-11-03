# Gopherdojo

Small, runnable Go katas in the Ziglings/Rustlings style: one topic at a time,
solved by you, verified by the machine.

This repo is the companion to our 85-topic Go curriculum — Go by Example depth,
Gopherlings mechanics. Exercises are added one per session, in order.

## Current exercise

| # | Topic | Directory | Status |
|---|-------|-----------|--------|
| 001 | Hello World | `exercises/001_hello_world` | ★ in progress |

## Solve it

```sh
cd gopherdojo
go run ./exercises/001_hello_world    # instant PASS/FAIL feedback
go run ./cmd/gopherdojo --only 1      # the real gate (runs go test)
```

Edit only `exercises/001_hello_world/main.go`. When `go test` is green you'll see:

```
✔ [1/1] exercises/001_hello_world passed.
```

Then paste your code to your mentor for review before moving to Topic 2.

## Runner commands

```sh
go run ./cmd/gopherdojo                     # check in order, stop at first failure
go run ./cmd/gopherdojo --only 1            # check one exercise (id or dir)
go run ./cmd/gopherdojo --hint              # print the hint when something fails
go run ./cmd/gopherdojo --list              # all exercises with status
go run ./cmd/gopherdojo --reset 1           # restore exercise 1 from .pristine/
go run ./cmd/gopherdojo --watch             # re-check on save (poll 500ms)
go run ./cmd/gopherdojo --verify-solutions  # maintainer: solutions pass, exercises fail
```

Exit code is non-zero on failure. `NO_COLOR=1` disables colors.

## Exercise 001 rules

You are building the deploy-notifier for a platform team. Its parser is strict.

1. `greet(app)` returns `Hello, <app>!`; an empty `app` returns `Hello, stranger!`.
2. `quoted(s)` returns `s` in double quotes, escaped the Go way (`%q`), so
   interior quotes become `\"` and printable Unicode like `café` stays intact.
3. `banner(app, version)` returns exactly three lines joined by `\n`, with no
   trailing newline:
   ```
   === <app> <version> ===
   Deploying to production
   Have a nice day!
   ```
4. `main()` and the `check(...)` calls are test harness — leave them alone.

## Layout

```
cmd/gopherdojo/       runner
exercises/NNN_topic/  the file you edit (+ exercise_test.go)
solutions/NNN_topic/  reference answer (peek only after passing)
.pristine/NNN_topic/  untouched originals for --reset
exercises.json        manifest: id, dir, mode, hint, topic
```

## Notes

- Requires Go 1.27+ (only the standard library is used).
- `go test ./...` will fail while an exercise is unsolved — that is by design.
- After solving, `go run ./cmd/gopherdojo` prints the next exercise when it exists.
