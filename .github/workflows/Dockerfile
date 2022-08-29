name: CI
on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: setup golang
        uses: actions/setup-go@v3
        with:
          go-version: '1.18.0'
      - name: install dependencies
        run: |
          go version
          go get -u golang.org/x/lint/golint
      - name: run build command
        run: go build .
      - name: run vet
        run: go vet .
  deploy:
    runs-on: ubuntu-latest
    needs: [build]
    if: ${{ github.ref == 'refs/heads/main' && github.event_name == 'push' }}
    steps:
        - name: deploy to github pages
          uses: thomps9012/financial-api@v3
          with:
            folder: build
            
