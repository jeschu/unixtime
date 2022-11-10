#!/usr/bin/env bash

function label() {
  label="${1}"
  echo '---------------------------------------------'
  echo "--- ${1}"
  echo '---------------------------------------------'
}

function clean() {
  label 'clean'
  go clean -v -x
  rm -rfv dist
}

function buildFor() {
  echo "--- buildFor ${1} ${2}"
  GOOS="${1}" GOARCH="${2}" go build -o "dist/${1}-${2}/unixtime"
}

function build() {
  clean
  label 'build'

  go build -v -a -o 'dist/unixtime'
  buildFor 'darwin'  'amd64'
  buildFor 'linux'   'amd64'
  buildFor 'linux'   'arm64'
  buildFor 'linux'   'arm'
  buildFor 'windows' 'amd64'
}

build