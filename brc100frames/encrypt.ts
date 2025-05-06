import {SecurityLevels, WalletEncryptArgs, WalletEncryptResult} from "@bsv/sdk";
import {generateWireFramesFor} from "./generator/wire-frames-generator";

export const encrypt: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'encrypt-simple': generateWireFramesFor(
        'encrypt',
        {
            ciphertext: [1, 2, 3, 4, 5, 6, 7, 8],
        } as WalletEncryptResult,
        {
            // Order matching actual Go output from failed test
            protocolID: [SecurityLevels.App, 'test-protocol'],
            keyID: 'test-key',
            counterparty: 'self',
            privileged: true,
            privilegedReason: 'test reason',
            seekPermission: true,
            plaintext: [1, 2, 3, 4],
        } as WalletEncryptArgs
    ),
    // Add more cases if needed
}
