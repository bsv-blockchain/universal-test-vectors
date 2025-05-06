import { generateWireFramesFor } from "./generator/wire-frames-generator";

export const getVersion: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'getVersion-simple': generateWireFramesFor(
        'getVersion',
        {
            version: "1.0.0",
        },
        {
            // No args for getVersion
        }
    ),
}
