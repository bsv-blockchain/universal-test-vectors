import {generateWireFramesFor} from "./generator/wire-frames-generator";
// Assuming correct path, adjust if necessary
// import {WalletCreateHmacArgs, WalletCreateHmacResult} from "../../../../ts-sdk/src/Wallet.interfaces";

export const createHmac: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'createHmac-simple': generateWireFramesFor(
        'createHmac',
        {
            hmac: [50, 60, 70, 80, 90, 100, 110, 120,
                50, 60, 70, 80, 90, 100, 110, 120,
                50, 60, 70, 80, 90, 100, 110, 120,
                50, 60, 70, 80, 90, 100, 110, 120],
        },
        {
            protocolID: [1 /* SecurityLevelEveryApp */, 'test-protocol'],
            keyID: 'test-key',
            counterparty: 'self',
            privileged: true,
            privilegedReason: 'test reason',
            seekPermission: true,
            data: [10, 20, 30, 40],
        }
    ),
    // Add more cases if needed
}
