import {generateWireFramesFor} from "./wire-frames-generator";

export const createAction: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'simple': generateWireFramesFor(
        'createAction',
        {
            txid: 'deadbeef20248806deadbeef20248806deadbeef20248806deadbeef20248806',
            tx: [1, 2, 3, 4],
        },
        {
            description: 'Test action description',
            outputs: [
                {
                    lockingScript: '00', // Sample locking script
                    satoshis: 1000,
                    outputDescription: 'Test output',
                    basket: 'test-basket',
                    customInstructions: 'Test instructions',
                    tags: ['test-tag']
                }
            ],
            labels: ['test-label'],
        }
    )
}
