# General intro

## These must be installed on your system:

- AWS SAM;
- Docker;
- Air (hot-reloading package for golang).

## CD to the lambda-directory you intend to work on, for example:

```BASH
  cd server-lambdas/
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

### To run locally without hot-reload (for cases when there is not enought RAM or it just doesn't make sense for the frequency of changes you add:

```BASH
  pnpm run dev:start
```

### To deploy to AWS:

```BASH
  pnpm run deploy
```
