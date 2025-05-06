import {generateWireFramesFor} from "./generator/wire-frames-generator";


export const revealSpecificKeyLinkage: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'revealSpecificKeyLinkage-simple': generateWireFramesFor(
        'revealSpecificKeyLinkage',
        {
            encryptedLinkage: [1, 2, 3, 4],
            encryptedLinkageProof: [5, 6, 7, 8],
            prover: "02e14bb4fbcd33d02a0bad2b60dcd14c36506fa15599e3c28ec87eff440a97a2b8",
            verifier: "03b106dae20ae8fca0f4e8983d974c4b583054573eecdcdcfad261c035415ce1ee",
            counterparty: "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1",
            protocolID: [2, "tests"],
            keyID: "test-key-id",
            proofType: 1,
        },
        {
            counterparty: "0294c479f762f6baa97fbcd4393564c1d7bd8336ebd15928135bbcf575cd1a71a1",
            verifier: "03b106dae20ae8fca0f4e8983d974c4b583054573eecdcdcfad261c035415ce1ee",
            protocolID: [2, "tests"],
            keyID: "test-key-id",
            privileged: true,
            privilegedReason: "test-reason"
        }
    ),
    // Add more cases if needed
}
