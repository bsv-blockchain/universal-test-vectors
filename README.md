# universal-test-vectors

> [!NOTE]
> This is only an MPV implementation - almost all mechanisms are copied from `spv-wallet` codebase.

This repository contains a set of test vectors (JSON files) that can be used to test the implementation in various programming languages.

- You can find the JSON files in [generated](./generated) directory.

## How it works

The idea is to have a readable set of golang functions that can generate test vectors and save them in JSON files. 

- The definitions of the test vectors are in the [vectors](./vectors) directory.

## Available test vectors

- BSV transactions from one sender to one receiver
- BSV transactions from one sender to one receiver and with OP_RETURN output
- Users with its private & public keys and additional information

## 
