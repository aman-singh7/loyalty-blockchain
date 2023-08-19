@echo off
set CWD=%cd%
set BLOCKCHAIN_PATH=blockchain
set BLOCKCHAIN_DIR=%CWD%\%BLOCKCHAIN_PATH%
cd %BLOCKCHAIN_DIR% && truffle compile && truffle run abigen