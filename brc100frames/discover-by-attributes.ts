import {generateWireFramesFor} from "./generator/wire-frames-generator";
import type {DiscoverByAttributesArgs, DiscoverCertificatesResult} from "@bsv/sdk";

// Define the keys used in the test cases
const pubKeyHex = "025ad43a22ac38d0bc1f8bacaabb323b5d634703b7a774c4268f6a09e4ddf79097";
const CounterpartyHex = "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1";

export const discoverByAttributes: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'discoverByAttributes-simple': generateWireFramesFor(
        'discoverByAttributes',
        { // Result object definition (reusing structure from discoverByIdentityKey)
            totalCertificates: 1,
            certificates: [
                {
                    // Embedded Certificate fields
                    type: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0ZXN0LXR5cGU=",
                    serialNumber: "AAAAAAAAAAAAAAAAAAB0ZXN0LXNlcmlhbC1udW1iZXI=",
                    subject: pubKeyHex,
                    certifier: CounterpartyHex,
                    revocationOutpoint: "txid123:0",
                    fields: {"name": "Alice", "email": "alice@example.com"},
                    signature: "7369676e61747572652d686578",

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
            attributes: {"email": "alice@example.com", "role": "admin"},
            limit: 5,
            offset: 0,
            seekPermission: false,
        } as DiscoverByAttributesArgs
    ),
    // Add more test cases here if needed
}
