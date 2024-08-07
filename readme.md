A quick follow-along of:

https://getanteon.com/blog/exploring-function-tracing-with-ebpf-and-uprobes/#kprobes-and-tracepoints-4aa6fae0-868e-439b-846a-6b79bb9cd271

I'm running on ubuntu 24.x LTS, kernel 6.8.0-39-generic.

```
go build bpf.go
go build main.go
clang -target bpf -O2 -g -o tracker.o -c bpf.c
```

Check that the main.Greet symbol exists and isn't inlined:
```
$ nm main|grep Greet
0000000000493520 T main.Greet
```

In one terminal, run:
`go run -exec sudo bpf.go`

in the other, run:
`./main # Note, not "go run main.go" as the program location won't match`

Sample output:

`2024/08/07 16:15:18 COUNT: map[Kere1:3 Lucas:4 Mauro:3]`

I think the
```
struct greet_event {
    char param[6];
};
```
limits the name.  I need to learn more about the call structure parse that correctly
