# Enviroment

- [aqua](https://aquaproj.github.io/)
- [editorconfig](https://editorconfig.org/)

# Versions

| Tools   | version |
| ------- | ------- |
| golang  | 1.21.4  |


# Install Tools

```bash
aqua -i
```

# Run Example

## Chat

```bash
cd ./chat
go build -o chat
./chat
```

## trace test

-cover=カバレッジの表示

```bash
cd ./trace
go test -cover
```
