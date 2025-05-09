import {generateWireFramesFor} from "./generator/wire-frames-generator";
import {AcquireCertificateArgs, AcquireCertificateResult} from "@bsv/sdk";

const SubjectPubKeyHex = "025ad43a22ac38d0bc1f8bacaabb323b5d634703b7a774c4268f6a09e4ddf79097";
const CertifierPubKeyHex = "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1";
const TypeBase64 = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0ZXN0LXR5cGU=";
const SerialNumberBase64 = "AAAAAAAAAAAAAAAAAAB0ZXN0LXNlcmlhbC1udW1iZXI=";

export const acquireCertificate: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'acquireCertificate-simple': generateWireFramesFor(
        'acquireCertificate',
        {
            // Result object (WalletCertificate) - JSON format
            type: TypeBase64,
            serialNumber: SerialNumberBase64,
            subject: SubjectPubKeyHex,
            certifier: CertifierPubKeyHex,
            revocationOutpoint: "txid123:0",
            fields: {email: "alice@example.com", name: "Alice"}, // Alphabetical
            signature: "sig-hex",
        } as AcquireCertificateResult,
        {
            // Args object (WalletAcquireCertificateArgs) - JSON format
            type: TypeBase64,
            certifier: CertifierPubKeyHex,
            acquisitionProtocol: "issuance",
            fields: {email: "alice@example.com", name: "Alice"}, // Alphabetical
            serialNumber: SerialNumberBase64,
            revocationOutpoint: "txid123:0",
            signature: "sig-hex",
            certifierUrl: "https://certifier.example.com",
            keyringRevealer: "revealer-key-hex", // Assuming this is a string ID/hex
            keyringForSubject: {field1: "key1", field2: "key2"},
            privileged: false,
        } as AcquireCertificateArgs
    ),
    // Add more cases if needed
}
