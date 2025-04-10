import {createAction} from "./create-action";
import {abortAction} from "./abort-action";
import {saveFile} from "./save-file";

const allFrames = {
    ...createAction,
    ...abortAction,
}

async function generate(destination: string) {
    for (const key in allFrames) {
        const obj = await allFrames[key]

        saveFile(destination, `${obj.name}-${key}-args`, obj.wire.args, obj.json.args)
        saveFile(destination, `${obj.name}-${key}-result`, obj.wire.result, obj.json.result)
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
