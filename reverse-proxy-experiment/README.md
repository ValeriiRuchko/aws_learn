# General intro

## These must be installed on your system:

- AWS SAM;
- Docker;
- Air (hot-reloading package for golang).

## CD to the lambda-directory you intend to work on, for example:

```BASH
  cd reverse-proxy-experiment/
```

## Run

```BASH
 pnpm install
```

AND

```BASH
  go mod tidy
```

## to install neccessary dependencies

### This will watch your local \*.go files and re-build a Lambda docker-image (API-gateway included):

```BASH
  air
```

### To deploy to AWS:

```BASH
  pnpm run deploy
```
