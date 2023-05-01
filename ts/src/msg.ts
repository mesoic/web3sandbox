"use strict"

import logSymbols from "log-symbols";

export class SandboxMessages {

    constructor() {}

    verified(message: string) {
        console.log(" ", logSymbols.success, "verified:", message)
    }

    warning(msg: string) {
        console.log(" ", logSymbols.warning, " warning:", msg)
    }

    error(msg: string) {
        console.log(" ", logSymbols.error, "   error:", msg)
    }

    critical(message: string) {
        console.log(logSymbols.error, logSymbols.error, "critical:", message)
    }

    info(msg: string) {
        console.log("\r", logSymbols.info, "     info:", msg)
    }

    message(msg: string) {
        console.log("\r", logSymbols.info, msg)        
    }

    log(msg: string) {
        console.log("> ", msg)
    }

    newline() {
        console.log()
    }
}
