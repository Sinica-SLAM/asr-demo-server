# asr-demo-server

## Structure

```structure
.
├── cmd 
│   └── server 
├── docs (swagger api docs) 
├── internal 
│   └── handler (http handlers)
└── pkg
    ├── chi (chi router)
    ├── ent (ent orm, but not used)
    ├── middleware (middlewares)
    └── youtube (youtube api client)
```

## Build

```bash
go build ./cmd/server
```

or cross compile for linux(amd64)

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/server
```

## Usage

### client files

put client files in `./dist` directory, you can change the directory in `./internal/handler/static.go`

### youtube service secret

put youtube service secret file as `./client_secret.json`, you can change the file path in `./pkg/youtube/client.go`
