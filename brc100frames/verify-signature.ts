import {generateWireFramesFor} from "./generator/wire-frames-generator";

// Function to convert hex string to number array
const hexToBytes = (hex: string): number[] => {
    const bytes: number[] = [];
    for (let c = 0; c < hex.length; c += 2) {
        bytes.push(Number.parseInt(hex.substr(c, 2), 16));
    }
    return bytes;
};

export const verifySignature: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'verifySignature-simple': generateWireFramesFor(
        'verifySignature',
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
            data: [11, 22, 33, 44],
            signature: hexToBytes("302502204e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd41020101"),
            // hashToDirectlyVerify: undefined, // Omitting for simple test
        }
    ),
    // Add more cases if needed
}
