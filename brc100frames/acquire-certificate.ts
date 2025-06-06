import {generateWireFramesFor} from "./generator/wire-frames-generator";
import {AcquireCertificateArgs, AcquireCertificateResult} from "@bsv/sdk";

const SubjectPubKeyHex = "025ad43a22ac38d0bc1f8bacaabb323b5d634703b7a774c4268f6a09e4ddf79097";
const CertifierPubKeyHex = "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1";
const TypeBase64 = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB0ZXN0LXR5cGU=";
const SerialNumberBase64 = "AAAAAAAAAAAAAAAAAAB0ZXN0LXNlcmlhbC1udW1iZXI=";
const RevocationOutpoint = "aec245f27b7640c8b1865045107731bfb848115c573f7da38166074b1c9e475d.0";

export const acquireCertificate: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'acquireCertificate-simple': generateWireFramesFor(
        'acquireCertificate',
        {
            // Result object (WalletCertificate) - JSON format
            type: TypeBase64,
            serialNumber: SerialNumberBase64,
            subject: SubjectPubKeyHex,
            certifier: CertifierPubKeyHex,
            revocationOutpoint: RevocationOutpoint,
            fields: {email: "alice@example.com", name: "Alice"}, // Alphabetical
            // signature pk = 95c5931552e547d72a292e9d6f59eef2b9f7e1576d8c7b49731b505117c0cdfa, msg = test message
            signature: "3045022100a6f09ee70382ab364f3f6b040aebb8fe7a51dbc3b4c99cfeb2f7756432162833022067349b91a6319345996faddf36d1b2f3a502e4ae002205f9d2db85474f9aed5a",
        } as AcquireCertificateResult,
        {
            // Args object (WalletAcquireCertificateArgs) - JSON format
            type: TypeBase64,
            certifier: CertifierPubKeyHex,
            acquisitionProtocol: "issuance",
            fields: {email: "alice@example.com", name: "Alice"}, // Alphabetical
            serialNumber: SerialNumberBase64,
            revocationOutpoint: RevocationOutpoint,
            signature: "3045022100a6f09ee70382ab364f3f6b040aebb8fe7a51dbc3b4c99cfeb2f7756432162833022067349b91a6319345996faddf36d1b2f3a502e4ae002205f9d2db85474f9aed5a",
            certifierUrl: "https://certifier.example.com",
            keyringRevealer: "025ad43a22ac38d0bc1f8bacaabb323b5d634703b7a774c4268f6a09e4ddf79097",
            keyringForSubject: {field1: "key1", field2: "key2"},
            privileged: false,
        } as AcquireCertificateArgs
    ),
    // Add more cases if needed
}
