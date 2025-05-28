import {generateWireFramesFor} from "./generator/wire-frames-generator";
import type {DiscoverByIdentityKeyArgs, DiscoverCertificatesResult} from "@bsv/sdk";

// Define the keys used in the test cases
const pubKeyHex = "025ad43a22ac38d0bc1f8bacaabb323b5d634703b7a774c4268f6a09e4ddf79097";
const CounterpartyHex = "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1";

export const discoverByIdentityKey: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'discoverByIdentityKey-simple': generateWireFramesFor(
        'discoverByIdentityKey',
        { // Result object definition
            totalCertificates: 1,
            certificates: [
                {
                    // Embedded Certificate fields
                    type: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0ZXN0LXR5cGU=",
                    serialNumber: "AAAAAAAAAAAAAAAAAAB0ZXN0LXNlcmlhbC1udW1iZXI=",
                    subject: pubKeyHex,
                    certifier: CounterpartyHex,
                    revocationOutpoint: "aec245f27b7640c8b1865045107731bfb848115c573f7da38166074b1c9e475d.0",
                    fields: {"name": "Alice", "email": "alice@example.com"},
                    signature: "3045022100a6f09ee70382ab364f3f6b040aebb8fe7a51dbc3b4c99cfeb2f7756432162833022067349b91a6319345996faddf36d1b2f3a502e4ae002205f9d2db85474f9aed5a",

                    // IdentityCertificate specific fields
                    certifierInfo: {
                        name: "Test Certifier",
                        iconUrl: "https://example.com/icon.png",
                        description: "Certifier description",
                        trust: 5,
                    },
                    publiclyRevealedKeyring: {"pubField": "pubKey"},
                    decryptedFields: {"name": "Alice"},
                }
            ]
        } as DiscoverCertificatesResult,
        { // Args object definition
            identityKey: CounterpartyHex,
            limit: 10,
            offset: 0,
            seekPermission: true,
        } as DiscoverByIdentityKeyArgs
    ),
    // Add more test cases here if needed
}
