# universal-test-vectors

This repository contains a set of test vectors (JSON files) that can be used to test the implementation in various programming languages.

- You can find the JSON files in [generated](./generated) directory.

## How it works

The idea is to have a readable set of definitions in golang that can generate test vectors and save them in JSON files. 

- The definitions are in the [vectors](./vectors) directory - Check it out, how it's done.
- The JSON generation is not needed to be done manually - there is a github action which does it automatically (`go generate ./...`).

Example json with a transaction fixture looks like this:
```json lines
{
  "sender": {}, //details about the sender,
  "recipient": {}, //details about the recipient,
  "tx_id": "278784aa4052bb87e85de1d436870297d4d523c6806ca38a61ef9d0657f0f020",
  "raw_hex": "0100000002d93ed6<...rest of raw_tx_hex>",
  "beef_hex": "0100beef02fde903010100<...rest of beef_tx_hex>",
  "ef_hex": "010000000000000000ef02d9<...rest of ef_tx_hex>",
}
```

## Available test vectors

- BSV transactions from one sender to one receiver
- BSV transactions from one sender to one receiver and with OP_RETURN output
- Users with its private & public keys and additional information
- Wire (binary) and JSON frames for the communication (for now `createAction` only)

## How to add new test vectors

### For transaction specification

- Add a new definition in the [vectors](./vectors) directory (analogy to the existing ones).
- Generate the JSON files by running `go generate ./...` in the root directory.
- ... or make a PR with the new definition and the JSON files will be re-generated automatically.

### For BRC-100 frame specification
- Add a new definition in the [brc100frames](./brc100frames) directory (analogy to the existing ones - see [brc100frames/create-action.ts](brc100frames/create-action.ts)).
- Generate the JSON files by running `go generate ./...` in the root directory.
- ... or make a PR with the new definition and the JSON files will be re-generated automatically.

In case of other-than-tx test vectors, you would need to add a new generator tool and define a JSON model ([models](./vectors/models)) and a mapper ([mappers](./vectors/mappers)).

> [!NOTE]
> Of course, the `GivenTx` tool has many more features, but this is just a start - to show how it could work.
