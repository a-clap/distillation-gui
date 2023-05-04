import { defineStore } from "pinia";
import { ErrorListener } from "../types/ErrorListener";
import { i18n } from "../main";
import { useLogStore } from "./log";

export const useErrorStore = defineStore('errors', {
    state: () => {
        return {
            errors: [] as string[],
            show: false,
            title: "",
            msg: "",
            last_code: 0 as number,
            skipped: [] as number[],
            log: useLogStore()
        }
    },
    actions: {
        init() {
            ErrorListener.subscribe(this.errorCallback)
            this.skipped = []
        },

        errorCallback(id: number) {
            let err = `errors.${id}`;
            if (i18n.global.te(err)) {
                err = i18n.global.t(err)
            } else {
                err = i18n.global.t('errors.unknown')
                err += id.toString()
            }

            this.log.add(id, err)
            if (this.skipped.includes(id)) {
                return
            }
            this.last_code = id
            this.open(i18n.global.t('errors.title'), err)
        },

        open(title: string, msg: string) {
            this.msg = msg
            this.title = title
            this.show = true
        },

        close() {
            this.show = false
        },
        skip() {
            this.skipped.push(this.last_code)
            this.close()
        },
        reset_skipped() {
            this.skipped = []
        }
    }
})