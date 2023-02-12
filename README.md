# solana-indexer-gen-types
Wrapper types used in solana-indexer-gen-go

**Warning** The types in this repo are experimental and subject to change

## Getting Started
Run tests 
```go
go test ./...
```

## TODO
Implement [Scanner]((https://pkg.go.dev/database/sql/?tab=doc#Scanner) ) and [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) Interface
- [ ] github.com/gagliardetto/binary U128
- [ ] github.com/gagliardetto/solana-go UnixTimeSeconds
- [ ] github.com/gagliardetto/solana-go DurationSeconds
- - [ ] github.com/gagliardetto/solana-go AccountMeta