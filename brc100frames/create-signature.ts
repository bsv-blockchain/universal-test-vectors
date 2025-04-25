import {generateWireFramesFor} from "./generator/wire-frames-generator";

// Function to convert hex string to number array
const hexToBytes = (hex: string): number[] => {
    const bytes: number[] = [];
    for (let c = 0; c < hex.length; c += 2) {
        bytes.push(Number.parseInt(hex.substr(c, 2), 16));
    }
    return bytes;
};

export const createSignature: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'createSignature-simple': generateWireFramesFor(
        'createSignature',
        {
            // Byte array representation for wire encoding
            signature: hexToBytes("3044022004000100020003000400050006000700080009000a000b000c000d000e000f00101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30310220ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
        },
        {
            protocolID: [1 /* SecurityLevelEveryApp */, 'test-protocol'],
            keyID: 'test-key',
            counterparty: 'self',
            privileged: true,
            privilegedReason: 'test reason',
            seekPermission: true,
            data: [11, 22, 33, 44],
            // hashToDirectlySign: undefined, // Omitting for simple test
        }
    ),
    // Add more cases if needed
}
