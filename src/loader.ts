export async function callbackHandler(callback: any) {
    if (callback.post_type == 'message') {
        return callback
    }
    if (callback.post_type == 'request') {
        return callback
    }
    if (callback.post_type == 'notice') {
        return callback
    }
    if (callback.post_type == 'meta_event') {
        return 2000
    }
}

export async function responseHandler(response: any) {
    if (response.status == 'ok') {
        return response
    }
    if (response.status == 'failed') {
        return response
    }
}