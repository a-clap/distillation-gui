import { defineStore } from "pinia";
import { ErrorListener } from "../types/ErrorListener";
import { i18n } from "../main";

export const useErrorStore = defineStore('errors', {
    state: () => {
        return {
            errors: [] as string[],
            show: false,
            title: "",
            msg: ""
        }
    },
    actions: {
        init() {
            ErrorListener.subscribe(this.errorCallback)
        },

        errorCallback(id: number) {
            let err = `errors.${id}`;
            if (i18n.global.te(err)) {
                err = i18n.global.t(err)
            } else {
                err = i18n.global.t('errors.unknown')
                err += id.toString()
            }
            this.open(i18n.global.t('errors.title'), err)
        },

        open(title: string, msg: string) {
            this.msg = msg
            this.title = title
            this.show = true
        },

        close() {
            this.show = false
        }
    }
})