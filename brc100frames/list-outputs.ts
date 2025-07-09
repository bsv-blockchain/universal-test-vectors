import {generateWireFramesFor} from "./generator/wire-frames-generator";


export const listOutputs: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'listOutputs-simple': generateWireFramesFor(
        'listOutputs',
        {
            totalOutputs: 2,
            BEEF: [1, 2, 3, 4],
            outputs: [{
                satoshis: 1000,
                spendable: true,
                outpoint: "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef.0",
            }, {
                satoshis: 5000,
                spendable: true,
                outpoint: "abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890.2",
            }]
        },
        {
            basket: "test-basket",
            tags: ["tag1", "tag2"],
            tagQueryMode: "any",
            include: "locking scripts",
            includeTags: true,
            limit: 10,
        }
    ),
    // Add more cases if needed
}
