import {createAction} from "./create-action";
import {abortAction} from "./abort-action";
import {saveFile} from "./generator/save-file";

const allFrames = {
    ...createAction,
    ...abortAction,
}

async function generate(destination: string) {
    for (const key in allFrames) {
        const obj = await allFrames[key]

        const prefix = `${obj.name}-${key}`

        saveFile(destination, `${prefix}-args`, obj.wire.args, obj.json.args)
        saveFile(destination, `${prefix}-result`, obj.wire.result, obj.json.result)
    }
}

if (require.main === module) {
    const destination = process.argv[2];
    if (!destination) {
        console.error("Please provide a destination argument.");
        process.exit(1);
    }
    generate(destination).catch(err => {
        console.error(err);
        process.exit(1);
    });
}
