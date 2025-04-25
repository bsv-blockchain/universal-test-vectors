import { generateWireFramesFor } from "./generator/wire-frames-generator";

export const getHeaderForHeight: Record<string, ReturnType<typeof generateWireFramesFor>> = {
    'getHeaderForHeight-simple': generateWireFramesFor(
        'getHeaderForHeight',
        {
            header: "0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c",
        },
        {
            height: 850000,
        }
    ),
}
