import {ElLoading} from "element-plus";
import {i18n} from "../main";
import {AppErrorCodes} from "../stores/error_codes";
import {ErrorListener} from "./ErrorListener";

let defaultText = i18n.global.t('loader.loading')

class loader {
    private loading: typeof ElLoading.service
    private timeoutID: number
    private errCode: AppErrorCodes

    constructor() {
        this.loading = ElLoading.service
        this.timeoutID = 0
        this.errCode = AppErrorCodes.Success
    }

    show(errCode: AppErrorCodes, timeout: number, text: string = defaultText) {
        this.loading = ElLoading.service({
            lock: true,
            text: text,
            background: 'rgba(0, 0, 0, 0.7)',
        })

        this.errCode = errCode

        this.timeoutID = setTimeout(() => {
            this.fail()
        }, timeout)
    }

    fail() {
        ErrorListener.sendError(this.errCode)
        this.close()
    }


    close() {
        if (this.loading != null) {
            clearTimeout(this.timeoutID)
            this.loading.close()
        }
    }
}

export var Loader: loader = new loader()