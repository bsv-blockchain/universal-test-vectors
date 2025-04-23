import {createAction} from "./create-action";
import {abortAction} from "./abort-action";
import {saveFile} from "./generator/save-file";
import {signAction} from "./sign-action";
import {listActions} from "./list-actions";
import {internalizeAction} from "./internalize-action";
import {listOutputs} from "./list-outputs";
import {relinquishOutput} from "./relinquish-output";
import {getPublicKey} from "./get-public-key";
import {revealCounterpartyKeyLinkage} from "./reveal-counterparty-key-linkage";
import {revealSpecificKeyLinkage} from "./reveal-specific-key-linkage";

const allFrames = {
    ...createAction,
    ...signAction,
    ...abortAction,
    ...listActions,
    ...internalizeAction,
    ...listOutputs,
    ...relinquishOutput,
    ...getPublicKey,
    ...revealCounterpartyKeyLinkage,
    ...revealSpecificKeyLinkage,
}

async function generate(destination: string) {
    for (const key in allFrames) {
        const obj = await allFrames[key]

        saveFile(destination, `${key}-args`, obj.wire.args, obj.json.args)
        saveFile(destination, `${key}-result`, obj.wire.result, obj.json.result)
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
