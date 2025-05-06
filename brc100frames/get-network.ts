import { generateWireFramesFor } from "./generator/wire-frames-generator";

export const getNetwork: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'getNetwork-simple': generateWireFramesFor(
        'getNetwork',
        {
            network: "mainnet",
        },
        {
            // No args for getNetwork
        }
    ),
}
