#!/usr/bin/env zx
import { SandboxEnv, SandboxEnvError } from "../dist/env.js"
import { SandboxMessages } from "../dist/msg.js"

/* process environment */
var env = new SandboxEnv()
var msg = new SandboxMessages()

let err = await env.source(`${__dirname}/config.yaml`)
if(SandboxEnvError.is(err)) {
    console.log(err.message)
    process.exit(1)
}

/* simple zx runner */
if (process.argv[3] == "start") {

    await $ `node ${env.getEnv('APP_ROOT')}/dist/run.js`

} else if (process.argv[3] == "env") {

	msg.message("env:")
	msg.message(`  APP_ROOT : ${env.getEnv('APP_ROOT')}`)
	msg.message(`  APP_ROOT : ${env.getEnv('APP_HOST')}`)
	msg.message(`  APP_ROOT : ${env.getEnv('APP_PORT')}`)

} else {

    msg.message("usage:")
    msg.message("  zx run.mjs run")
    msg.message("  zx run.mjs env")
}
