import {generateWireFramesFor} from "./wire-frames-generator";


export const abortAction: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'abortAction-simple': generateWireFramesFor(
        'abortAction',
        {
            aborted: true,
        },
        {
            reference: "dGVzdA==",
        }
    ),
    // Add more cases if needed
}
