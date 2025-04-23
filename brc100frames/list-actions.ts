import {generateWireFramesFor} from "./generator/wire-frames-generator";

export const listActions: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'listActions-simple': generateWireFramesFor(
        'listActions',
        {
            totalActions: 1,
            actions: [{
                txid: "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
                satoshis: 1000,
                status: "completed",
                isOutgoing: true,
                description: "Test transaction 1",
                version: 1,
                lockTime: 10,
                outputs: [{
                    satoshis: 1000,
                    lockingScript: "76a9143cf53c49c322d9d811728182939aee2dca087f9888ac",
                    spendable: true,
                    tags: ["tag1", "tag2"],
                    outputIndex: 1,
                    outputDescription: "Test output",
                    basket: "basket1",
                }]
            }]
        },
        {
            labels: ["test-label"],
            includeOutputs: true,
            limit: 10,
        }
    ),
    // Add more cases if needed
}
