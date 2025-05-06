import {generateWireFramesFor} from "./generator/wire-frames-generator";
import type {RelinquishCertificateArgs, RelinquishCertificateResult} from "@bsv/sdk";

// Define the keys used in the test cases
const CounterpartyHex = "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1";

export const relinquishCertificate: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'relinquishCertificate-simple': generateWireFramesFor(
        'relinquishCertificate',
        { // Result object definition
            relinquished: true,
        } as RelinquishCertificateResult,
        { // Args object definition
            type: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0ZXN0LXR5cGU=",
            serialNumber: "AAAAAAAAAAAAAAAAAAB0ZXN0LXNlcmlhbC1udW1iZXI=",
            certifier: CounterpartyHex,
        } as RelinquishCertificateArgs
    ),
    // Add more test cases here if needed
}
