# Gopherdojo — maintainer spec

This repo is a Ziglings/Rustlings-style Go dojo. Every exercise is a small,
self-contained `package main` in `exercises/NNN_slug/` graded by
`exercises/NNN_slug/exercise_test.go`. The runner (`cmd/gopherdojo`) executes
`go test -count=1 ./exercises/NNN_slug` and stops at the first failure.

The learner edits ONLY `exercises/NNN_slug/main.go`. Nothing else.

## Hard requirements (all are verified by `--verify-solutions`)

For every exercise `NNN_slug`:

1. `exercises/NNN_slug/main.go` compiles on its own and its test suite FAILS
   (at least one test) before the learner edits it.
2. `solutions/NNN_slug/main.go` + `solutions/NNN_slug/exercise_test.go` PASS.
3. `.pristine/NNN_slug/` holds byte-identical copies of the starter
   `main.go` and `exercise_test.go` (plus any extra asset such as `.txt`).
4. Tests are deterministic, hermetic, and fast. No network, no external
   services, no reliance on wall-clock precision, no `os.Exit` inside tests,
   no real signals, no `time.Sleep` over 100ms, no unbounded channel waits.
   Any channel receive in a test must be wrapped in
   `select { case ...: case <-time.After(2*time.Second): t.Fatal("timeout") }`.
5. Standard library only. Do not touch `go.mod`. Windows must work.
6. `gofmt` clean.
7. No `init()` side effects, no global mutable state that leaks between tests.

## main.go conventions

```go
// NNN: Topic Name. One or two lines of concept.
//
// Scenario: a short real-world framing for the task.
//
// Rules:
//  1. ...
//  2. ...
//  3. ...
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt" // only imports that are actually used

// FunctionName does X.
func FunctionName(...) ... {
	// TODO
	return zeroValue
}

func main() {
	// small deterministic demo that calls the functions
}
```

- Starters must COMPILE. Never leave an unused import or unused local variable.
  Returning zero values (`0`, `""`, `nil`, `false`) is the standard placeholder.
- Do not add `func main()` twice. Exactly one per directory.
- The header comment is the "word problem": scenario + 3-4 strict rules.
- Keep the API small: 1-3 functions per exercise, exported names only when the
  test needs them. Tests live in the same package, so unexported is fine.
- `main()` is a tiny demo, not graded. Do not put logic in it that the test needs.

## exercise_test.go conventions

```go
package main

import "testing"

func TestThing(t *testing.T) {
	tests := []struct {
		name string
		in   ...
		want ...
	}{
		{"case", ..., ...},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Thing(tt.in); got != tt.want {
				t.Errorf("Thing(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
```

- Table-driven with `t.Run`, and a helpful failure message showing got vs want.
- Cover the edge cases named in the header rules (empty input, nil, error paths,
  boundary values, copy-vs-alias, etc.).
- Tests are the specification. They must be readable and answer-free.

## solutions

`solutions/NNN_slug/main.go` is the clean, idiomatic reference. Same header
comment as the starter minus the TODO. `main()` may stay identical.
Copy the exercise's `exercise_test.go` into the solution directory verbatim.

## Manifest fragment

Write `maintainer/manifest_<BATCH>.json`: a JSON array, in id order, of:

```json
{
  "id": "007",
  "dir": "007_switch",
  "mode": "test",
  "hint": "One answer-free sentence pointing at the concept.",
  "topic": "Switch"
}
```

`topic` must be the official name from the table below.

## Verification you must run (scope to your own exercises only)

```powershell
gofmt -l exercises/<your dirs> solutions/<your dirs>
go test -count=1 ./solutions/NNN_slug      # must PASS
go test -count=1 ./exercises/NNN_slug      # must FAIL (exit 1)
go build ./exercises/NNN_slug              # must succeed (main package compiles)
```

Do NOT run `go test ./...` or `go build ./...`: other batches may be in flight.

## Topic table (id, dir, official topic, notes)

| id | dir | topic | implementation notes |
|----|-----|-------|----------------------|
| 001 | 001_hello_world | Hello World | ALREADY DONE — do not modify |
| 002 | 002_values | Values | int/float/string/bool values, arithmetic, formatting |
| 003 | 003_variables | Variables | `var`, `:=`, multiple assignment, swap |
| 004 | 004_constants | Constants | `const`, typed vs untyped, `iota` block |
| 005 | 005_for | For | classic for, while-style, infinite + break |
| 006 | 006_if_else | If/Else | if with init statement, else-if chain |
| 007 | 007_switch | Switch | tagless switch, multi-value case, `fallthrough` |
| 008 | 008_arrays | Arrays | `[N]T` value semantics, copy on assignment |
| 009 | 009_slices | Slices | append, len/cap, copy, aliasing |
| 010 | 010_maps | Maps | literal, comma-ok, delete, count words |
| 011 | 011_functions | Functions | params, return, function calling function |
| 012 | 012_multiple_returns | Multiple Return Values | divmod / minmax |
| 013 | 013_variadics | Variadic Functions | `...T`, spreading a slice |
| 014 | 014_closures | Closures | counter factory, per-closure state |
| 015 | 015_recursion | Recursion | factorial, fibonacci, base cases |
| 016 | 016_range_builtins | Range over Built-in Types | slice index/value, map, string runes |
| 017 | 017_pointers | Pointers | `&`, `*`, mutate through pointer, nil check |
| 018 | 018_strings_runes | Strings and Runes | byte vs rune length, reverse runes, UTF-8 |
| 019 | 019_structs | Structs | literal, field access, value copy |
| 020 | 020_methods | Methods | value vs pointer receiver mutation |
| 021 | 021_interfaces | Interfaces | implicit satisfaction, interface slice |
| 022 | 022_enums | Enums | `type X int` + `iota` + `String()` |
| 023 | 023_struct_embedding | Struct Embedding | promoted fields/methods |
| 024 | 024_generics | Generics | type params, constraints, generic func |
| 025 | 025_range_iterators | Range over Iterators | `iter.Seq[T]`, yield func |
| 026 | 026_errors | Errors | sentinel error, return + check, `errors.Is` |
| 027 | 027_custom_errors | Custom Errors | custom type implementing `error`, `errors.As` |
| 028 | 028_goroutines | Goroutines | `go`, WaitGroup or done channel |
| 029 | 029_channels | Channels | unbuffered send/receive |
| 030 | 030_channel_buffering | Channel Buffering | `make(chan T, n)`, len/cap |
| 031 | 031_channel_sync | Channel Synchronization | done channel handshake |
| 032 | 032_channel_directions | Channel Directions | `chan<-` / `<-chan` params |
| 033 | 033_select | Select | ready case vs never-ready case |
| 034 | 034_timeouts | Timeouts | `select` + `time.After`, tiny durations |
| 035 | 035_nonblocking_channels | Non-Blocking Channel Operations | `select` + `default` |
| 036 | 036_closing_channels | Closing Channels | close, comma-ok, double-close care |
| 037 | 037_range_channels | Range over Channels | range until close, then ok=false |
| 038 | 038_timers | Timers | `time.NewTimer`/`Reset`, ms durations |
| 039 | 039_tickers | Tickers | `time.NewTicker`, collect N, Stop |
| 040 | 040_worker_pools | Worker Pools | fixed workers, jobs/results channels |
| 041 | 041_waitgroups | WaitGroups | `Add/Done/Wait`, shared counter |
| 042 | 042_rate_limiting | Rate Limiting | ticker limiter, small N, deterministic order |
| 043 | 043_atomic_counters | Atomic Counters | `sync/atomic` Int64, fixed goroutines |
| 044 | 044_mutexes | Mutexes | `sync.Mutex` protecting a map/counter |
| 045 | 045_stateful_goroutines | Stateful Goroutines | request/reply structs over channels |
| 046 | 046_sorting | Sorting | `sort.Ints`, `sort.Strings`, check sorted |
| 047 | 047_sorting_functions | Sorting by Functions | `sort.Slice`, custom less, stable |
| 048 | 048_panic | Panic | explicit panic; test recovers |
| 049 | 049_defer | Defer | LIFO order, defer mutation |
| 050 | 050_recover | Recover | recover in deferred func, convert panic to error |
| 051 | 051_string_functions | String Functions | Contains/Count/Index/Join/Repeat/Replace/Split/Trim |
| 052 | 052_string_formatting | String Formatting | `%v %+v %#v %T %d %.2f %q %s` |
| 053 | 053_text_templates | Text Templates | `text/template`, Execute to buffer |
| 054 | 054_regular_expressions | Regular Expressions | MatchString, FindStringSubmatch, ReplaceAllString |
| 055 | 055_json | JSON | marshal/unmarshal, struct tags, omitempty |
| 056 | 056_xml | XML | marshal/unmarshal, xml tags |
| 057 | 057_time | Time | time.Date, Add, Sub, Before/After, UTC |
| 058 | 058_epoch | Epoch | Unix/UnixMilli/UnixNano roundtrip |
| 059 | 059_time_formatting | Time Formatting / Parsing | layout `2006-01-02 15:04`, Parse/Format |
| 060 | 060_random_numbers | Random Numbers | `rand.New(rand.NewSource(seed))`, Intn range |
| 061 | 061_number_parsing | Number Parsing | `strconv.Atoi`, ParseFloat, ParseBool, errors |
| 062 | 062_url_parsing | URL Parsing | `url.Parse`, Host, Path, Query().Get |
| 063 | 063_sha256 | SHA256 Hashes | `sha256.Sum256`, hex encoding |
| 064 | 064_base64 | Base64 Encoding | StdEncoding + URLEncoding roundtrip |
| 065 | 065_reading_files | Reading Files | `os.ReadFile` on a path passed in; tests use `t.TempDir` |
| 066 | 066_writing_files | Writing Files | `os.WriteFile` + read back; `t.TempDir` |
| 067 | 067_line_filters | Line Filters | `bufio.Scanner`, `io.Reader`/`io.Writer` params |
| 068 | 068_file_paths | File Paths | `filepath.Join/Base/Dir/Ext`, OS-aware expectations |
| 069 | 069_directories | Directories | `os.MkdirAll`, `os.ReadDir`, `t.TempDir` |
| 070 | 070_temp_files | Temporary Files and Directories | `os.CreateTemp`, `os.MkdirTemp`, cleanup |
| 071 | 071_embed | Embed Directive | `//go:embed data/greeting.txt`; asset copied to all 3 dirs |
| 072 | 072_testing_benchmarking | Testing and Benchmarking | testable func + BenchmarkXxx in test file |
| 073 | 073_cmdline_arguments | Command-Line Arguments | parse `[]string` parameter; main uses `os.Args` |
| 074 | 074_cmdline_flags | Command-Line Flags | `flag.NewFlagSet`, parse `[]string` |
| 075 | 075_cmdline_subcommands | Command-Line Subcommands | dispatch `[]string` to add/list handlers |
| 076 | 076_env_vars | Environment Variables | `os.Getenv` + defaults; tests use `t.Setenv` |
| 077 | 077_logging | Logging | `log.New(buf, "", 0)`, deterministic output |
| 078 | 078_http_client | HTTP Client | `httptest.NewServer`, status/body |
| 079 | 079_http_server | HTTP Server | handler + `httptest.NewRecorder`, query params |
| 080 | 080_tcp_server | TCP Server | listen `127.0.0.1:0`, echo once, deadlines |
| 081 | 081_context | Context | `context.WithTimeout`, cancellation wins |
| 082 | 082_spawning_processes | Spawning Processes | `exec.Command("go", "version").Start/Wait` |
| 083 | 083_execing_processes | Exec'ing Processes | `CombinedOutput` of `go env GOOS`, compare `runtime.GOOS` |
| 084 | 084_signals | Signals | pure `signalLabel(os.Signal)`; injected channel, no real signals |
| 085 | 085_exit | Exit | pure `exitCode(err) int` mapping; main calls `os.Exit` |

## Topics not to accidentally duplicate

Keep each exercise focused on its row. If a later concept is needed to make a
test work (e.g. `errors.As` in 027), it is fine, but do not pre-teach the whole
curriculum in one file.
