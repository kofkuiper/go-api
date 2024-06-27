# Echo API

## To run
```sh
$ air
```

## To generate abi
- Don't forget to install Abigen
```sh
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

```sh
$ abigen --abi ./abis/pluto_abi.json -pkg repositories --type Pluto --out ./repositories/pluto_abi.go
```