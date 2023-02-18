/**@desc delay util that resolve after the timeout provided */
export const sleep = (timeout: number) => {
    return new Promise<void>((resolve)=> setTimeout(resolve,timeout));
}