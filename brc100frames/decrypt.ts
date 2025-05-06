import {SecurityLevels} from "@bsv/sdk";
import type {WalletDecryptArgs, WalletDecryptResult} from "@bsv/sdk";
import {generateWireFramesFor} from "./generator/wire-frames-generator";

export const decrypt: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'decrypt-simple': generateWireFramesFor(
        'decrypt',
        {
            plaintext: [1, 2, 3, 4],
        } as WalletDecryptResult,
        {
            // Order matching actual Go output from failed test
            protocolID: [SecurityLevels.App, 'test-protocol'],
            keyID: 'test-key',
            counterparty: 'self',
            privileged: true,
            privilegedReason: 'test reason',
            seekPermission: true,
            ciphertext: [1, 2, 3, 4, 5, 6, 7, 8],
        } as WalletDecryptArgs
    ),
    // Add more cases if needed
}
