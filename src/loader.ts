import * as examplePlugin from "../plugin/example"

export async function callbackHandler(callback: any) {
    if (callback.post_type == 'message') {
        return pluginRegistry("message", callback)
    }
    if (callback.post_type == 'request') {
        return pluginRegistry("request", callback)
    }
    if (callback.post_type == 'notice') {
        return pluginRegistry("notice", callback)
    }
    if (callback.post_type == 'meta_event') {
        return pluginRegistry("meta_event", callback)
    }
}

async function pluginRegistry(type: string, callback: any) {
    await examplePlugin.handler(type, callback)
}