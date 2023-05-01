"use strict"
import * as fs from 'fs'
import * as path from 'path'

import YAML from "yaml"

import { fileURLToPath } from 'url';
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

import { SandboxMessages } from "./msg.js"

export namespace SandboxEnvError {

    export class error extends Error {

        constructor(message) {
            super(message);
            this.name = "env-error";
        }
    }

    export function is(e) {
        return e && e.stack && e.message && (e.name == "env-error");
    }
}

export class SandboxEnv {

    msg: SandboxMessages = new SandboxMessages()

    constructor() { 

    	this.register(
    		"APP_ROOT", 
    		path.resolve(__dirname, "../")
        )
    }

    private register(key: string, value: string) {
        process.env[key] = value
    }

    source(config_path: string, verbose = false) {

    	/*  check env config_path */
        if (!fs.existsSync(config_path)) {
         
            return new SandboxEnvError.error(`config file does not exist ${config_path}`)
        }

        /* attempt to source config */
        let env
		try {
	
			env = YAML.parse(fs.readFileSync(config_path, 'utf8'))

		} catch (err) {

			return new SandboxEnvError.error(err.message)
		}

		/*  register environmet */
        for (let key of Object.keys(env)) {

            if (["APP_ROOT"].includes(key)) {

                return new SandboxEnvError.error(`Cannot assign protected variable ${key}`)
            }
            
            this.register(key, env[key])
        }
    }

    getEnv(key : string) {

        return process.env[key]
    }
}
