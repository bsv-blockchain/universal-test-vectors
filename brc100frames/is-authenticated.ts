import {generateWireFramesFor} from "./generator/wire-frames-generator";
import type {AuthenticatedResult} from "@bsv/sdk";

export const isAuthenticated: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'isAuthenticated-simple': generateWireFramesFor(
        'isAuthenticated',
        { // Result object definition
            authenticated: true,
        } as AuthenticatedResult,
        { // Args object definition (empty for this call)
        } // as IsAuthenticatedArgs - No specific args type
    ),
    // Add more test cases here if needed
}
