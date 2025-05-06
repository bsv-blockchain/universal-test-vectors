import {generateWireFramesFor} from "./generator/wire-frames-generator";


export const relinquishOutput: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'relinquishOutput-simple': generateWireFramesFor(
        'relinquishOutput',
        {
            relinquished: true,
        },
        {
            basket: "test-basket",
            output: "abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890.2",
        }
    ),
    // Add more cases if needed
}
