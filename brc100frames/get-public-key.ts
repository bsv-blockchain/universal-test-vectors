import {generateWireFramesFor} from "./generator/wire-frames-generator";
import {PrivateKey, PublicKey} from "@bsv/sdk";

const privKey = PrivateKey.fromString('6a2991c9de20e38b31d7ea147bf55f5039e4bbc073160f5e0d541d1f17e321b8');
const counterpartyKey = PublicKey.fromString("0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1");

export const getPublicKey: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'getPublicKey-simple': generateWireFramesFor(
        'getPublicKey',
        {
            publicKey: privKey.toPublicKey().toString(),
        },
        {
            protocolID: [2, "tests"],
            keyID: "test-key-id",
            counterparty: counterpartyKey.toString(),
            privileged: true,
            privilegedReason: "privileged reason",
            seekPermission: true,
        }
    ),
    // Add more cases if needed
}
