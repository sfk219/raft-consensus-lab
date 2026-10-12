# Raft Consensus Lab (Go)

A deterministic educational simulator of **Raft-inspired distributed coordination** built as an independent portfolio project.

**Features:** simplified leader election, majority-quorum checks, replicated key/value state, simulated leader failure and recovery, and tests for stale candidates.

**Tech stack:** Go 1.22+, distributed systems, fault tolerance, automated testing.

## Run locally
```bash
go run ./cmd/demo
go test ./... -v
```

> This is a learning-oriented simulator, not a production Raft implementation. It does not reproduce employer or client code.

[Return to my GitHub profile](https://github.com/sfk219)
