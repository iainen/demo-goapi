#!/usr/bin/env zsh

export BASEDIR=$(dirname "$0")/..
export TEMPDIR="$BASEDIR"/temp
export PROTOCOLDIR="$BASEDIR"/temp/protocols

echo "$BASEDIR"
echo "$TEMPDIR"
echo "$PROTOCOLDIR"


function initProtocolHome(){
  if [[ -d "$PROTOCOLDIR" ]]; then
    rm -rf "$PROTOCOLDIR}"
  fi
  mkdir -p "$PROTOCOLDIR"
}

function cleanHistoryCodes(){
  rm -rf "$BASEDIR"/demo
}

function generateCodes(){
  protoc \
  --proto_path="$PROTOCOLDIR"/demo-protocol \
  --proto_path="$PROTOCOLDIR" \
  --go_out="$BASEDIR" \
  --go-grpc_out="$BASEDIR" \
  "$PROTOCOLDIR"/demo-protocol/*/*.proto

  mv "$BASEDIR"/goapi/demo "$BASEDIR"/ \
  && rm -rf "$BASEDIR"/goapi && rm -rf "$TEMPDIR"
  go mod tidy
}

initProtocolHome
cleanHistoryCodes

cp -r "$BASEDIR"/../demo-protocol "$PROTOCOLDIR"/
cp -r "$BASEDIR"/../sky-apm-demo/service/protobuf/google1 "$PROTOCOLDIR"/google

generateCodes
