import {generateWireFramesFor} from "./generator/wire-frames-generator";


export const internalizeAction: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'internalizeAction-simple': generateWireFramesFor(
        'internalizeAction',
        {
            accepted: true,
        },
        {
            tx: [1, 2, 3, 4],
            description: "test transaction",
            labels: ["label1", "label2"],
            seekPermission: true,
            outputs: [{
                outputIndex: 0,
                protocol: "wallet payment",
                paymentRemittance: {
                    derivationPrefix: "prefix",
                    derivationSuffix: "suffix",
                    senderIdentityKey: "sender-key",
                }
            }, {
                outputIndex: 1,
                protocol: "basket insertion",
                insertionRemittance: {
                    basket: "test-basket",
                    customInstructions: "instruction",
                    tags: ["tag1", "tag2"],
                }
            }]
        }
    ),
    // Add more cases if needed
}
