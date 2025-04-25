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
import {encrypt} from "./encrypt";
import {decrypt} from "./decrypt";
import {createHmac} from "./create-hmac";
import {verifyHmac} from "./verify-hmac";
import {createSignature} from "./create-signature";
import {verifySignature} from "./verify-signature";
import {acquireCertificate} from "./acquire-certificate";
import {listCertificates} from "./list-certificates";
import {proveCertificate} from "./prove-certificate";
import {relinquishCertificate} from "./relinquish-certificate";
import {discoverByIdentityKey} from "./discover-by-identity-key";
import {discoverByAttributes} from "./discover-by-attributes";

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
    ...encrypt,
    ...decrypt,
    ...createHmac,
    ...verifyHmac,
    ...createSignature,
    ...verifySignature,
    ...acquireCertificate,
    ...listCertificates,
    ...proveCertificate,
    ...relinquishCertificate,
    ...discoverByIdentityKey,
    ...discoverByAttributes,
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
