#!/bin/sh
CWD=$(pwd)
BLOCKCHAIN_PATH="blockchain"
BLOCKCHAIN_DIR="$CWD/$BLOCKCHAIN_PATH"

cd $BLOCKCHAIN_DIR && truffle compile && truffle run abigen
