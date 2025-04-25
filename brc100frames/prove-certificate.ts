import {generateWireFramesFor} from "./generator/wire-frames-generator";
import {ProveCertificateArgs, ProveCertificateResult} from "@bsv/sdk";

// Define the keys used in the test cases
const pubKeyHex = "025ad43a22ac38d0bc1f8bacaabb323b5d634703b7a774c4268f6a09e4ddf79097";
const CounterpartyHex = "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1";
const VerifierHex = "03b106dae20ae8fca0f4e8983d974c4b583054573eecdcdcfad261c035415ce1ee";

export const proveCertificate: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'proveCertificate-simple': generateWireFramesFor(
        'proveCertificate',
        { // Result object definition
            keyringForVerifier: {"name": "name-key"},
        } as ProveCertificateResult,
        { // Args object definition
            certificate: {
                // Embedded Certificate fields
                type: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0ZXN0LXR5cGU=",
                serialNumber: "AAAAAAAAAAAAAAAAAAB0ZXN0LXNlcmlhbC1udW1iZXI=",
                subject: pubKeyHex,
                certifier: CounterpartyHex,
                revocationOutpoint: "txid123:0",
                fields: {"name": "Alice", "email": "alice@example.com"},
                signature: "7369676e61747572652d686578", // hex for "signature-hex"
            },
            // ProveCertificateArgs specific fields
            fieldsToReveal: ["name"],
            verifier: VerifierHex,
            privileged: false,
            privilegedReason: "prove-reason",
        } as ProveCertificateArgs
    ),
    // Add more test cases here if needed
}
