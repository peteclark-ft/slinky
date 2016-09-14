# Slinky

Quickly test an html page for broken / non-200 links.

```sh
go get -u github.com/peteclark-ft/slinky
```

## Usage

Normal usage.

```sh
slinky --url https://github.com/peteclark-ft/slinky
```

Make it quicker.

```sh
slinky --url https://github.com/peteclark-ft/slinky --threads 400
```

Show even when it's a 200 response.

```sh
slinky --url https://github.com/peteclark-ft/slinky --details
```
