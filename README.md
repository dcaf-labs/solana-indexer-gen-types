# solana-indexer-gen-types

Wrapper types used in solana-indexer-gen-go

**Warning** The types in this repo are experimental and subject to change

## Getting Started

Run tests

```go
go test ./...
```

## TODO

Implement [Scanner](https://pkg.go.dev/database/sql/?tab=doc#Scanner) and [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) Interface

- [x] github.com/gagliardetto/binary Uint128
- [x] Add tests for negative values in int128
- [x] github.com/gagliardetto/solana-go AccountMeta
