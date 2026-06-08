# silo-plugin-requests-arr

A Silo host plugin that fulfills Silo content requests (movies and series)
against one or more Sonarr/Radarr instances. It implements the
`request_router.v1` capability defined in the Silo plugin SDK, letting Silo
route an approved request to the correct instance (HD vs. 4K, anime overrides,
per-instance root folder / quality profile / tags) and trigger the add + search.

The plugin is stateless: it stores nothing. All credentials and per-connection
configuration (service endpoint, API key, root folder, quality profile, tags,
default/4K/anime flags, etc.) are supplied by the Silo host on every call.

## Build

```sh
make build
```

This produces a `plugin` binary. Use `make build-all` to cross-compile the
release matrix (linux/amd64, linux/arm64, darwin/arm64) into `dist/`.

## Development note

The SDK dependency currently resolves through a machine-local `replace`
directive in `go.mod` pointing at a side-by-side checkout of
`silo-plugin-sdk`. Before release, swap this for a published, versioned
`github.com/Silo-Server/silo-plugin-sdk` require (CI rejects committed local
`replace` directives).
