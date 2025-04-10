import fs from 'fs';
import path from 'path';

export function saveFile(destination: string, filename: string, wireHex: string, jsonObj: object) {
    const obj = {
        json: jsonObj,
        wire: wireHex
    }

    const filepath = path.join(destination, filename + '.json');

    if (!fs.existsSync(destination)) {
        fs.mkdirSync(destination, { recursive: true });
    }
    fs.writeFileSync(filepath, JSON.stringify(obj, null, 2));
}
