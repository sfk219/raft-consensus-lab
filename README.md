# Raft Consensus Lab (Go)

A runnable **educational Raft-inspired cluster simulator**, independently written as a portfolio demonstration. Models synchronous leader elections, majority quorum, state replication, leader failure and recovery. **Not** a complete Raft implementation or production distributed database.

## Run and test
```bash
go test ./...
go run ./cmd/demo
```

**Code:** `raftsim/cluster.go`, `raftsim/cluster_test.go`, `cmd/demo/main.go`.

## Features
- Elect leader only with a majority and a sufficiently up-to-date log.
- Replicate key/value writes among reachable nodes.
- Refuse writes if the majority is unavailable.
- Recover an offline node by synchronizing committed data.

**Tech:** Go, distributed systems, consensus concepts, testing. No client/employer code.