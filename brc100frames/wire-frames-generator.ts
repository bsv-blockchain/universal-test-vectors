import {WalletInterface, WalletWireProcessor, WalletWireTransceiver} from "@bsv/sdk";

export async function generateWireFramesFor<
    Keys extends keyof WalletInterface,
    Method extends WalletInterface[Keys],
    Arguments extends Parameters<Method>,
    ResultType extends Awaited<ReturnType<Method>>,
>(actionName: Keys, resultObj: ResultType, ...args: Arguments) {
    const processor = new WalletWireProcessor({
        [actionName]: () => Promise.resolve(resultObj)
    } as any)

    const originalTransmitToWallet = processor.transmitToWallet;
    const mockTransmitToWallet = {
        call: null as number[] | null,
        result: null as number[] | null,
        async transmitToWallet(message: number[]) {
            mockTransmitToWallet.call = message
            const result = await originalTransmitToWallet.apply(processor, [message]);
            mockTransmitToWallet.result = result
            return result;
        }
    };
    processor.transmitToWallet = mockTransmitToWallet.transmitToWallet;

    const wallet = new WalletWireTransceiver(processor)

    // @ts-ignore
    await wallet[actionName](...args)

    return {
        wire: {
            args: frameNumbersToHex(mockTransmitToWallet.call!),
            result: frameNumbersToHex(mockTransmitToWallet.result!)
        },
        json: {
            args: args[0],
            result: resultObj,
        }

    }
}

function frameNumbersToHex(numbers: number[]): string {
    const body = new Uint8Array(numbers)
    return Array.from(body, byte => byte.toString(16).padStart(2, '0')).join('');
}
