import {generateWireFramesFor} from "./generator/wire-frames-generator";
import type {GetHeightResult} from "@bsv/sdk";

export const getHeight: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'getHeight-simple': generateWireFramesFor(
        'getHeight',
        { // Result object definition
            height: 850000,
        } as GetHeightResult,
        { // Args object definition (empty for this call)
        } // as GetHeightArgs - No specific args type
    ),
    // Add more test cases here if needed
}
