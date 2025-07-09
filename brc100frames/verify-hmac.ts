import {generateWireFramesFor} from "./generator/wire-frames-generator";

export const verifyHmac: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'verifyHmac-simple': generateWireFramesFor(
        'verifyHmac',
        {
            valid: true,
        },
        {
            protocolID: [1 /* SecurityLevelEveryApp */, 'test-protocol'],
            keyID: 'test-key',
            counterparty: 'self',
            privileged: true,
            privilegedReason: 'test reason',
            seekPermission: true,
            data: [10, 20, 30, 40],
            hmac: [50, 60, 70, 80, 90, 100, 110, 120,
                50, 60, 70, 80, 90, 100, 110, 120,
                50, 60, 70, 80, 90, 100, 110, 120,
                50, 60, 70, 80, 90, 100, 110, 120],
        }
    ),
    // Add more cases if needed
}
