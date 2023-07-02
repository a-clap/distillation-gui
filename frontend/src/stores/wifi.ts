import {defineStore} from "pinia";
import Parameter from "../types/Parameter";
import {WifiAPList, WifiConnect} from "../../wailsjs/go/backend/Backend";
import {Loader} from "../types/Loader";
import {AppErrorCodes} from "./error_codes";

export const useWIFIStore = defineStore('wifi', {
    state: () => {
        return {
            apList: [] as AP[],
            ssid: "",
            enabled: false,
            busy: false,
            connected: false,
            password: new Parameter("", false, () => {
            }) as Parameter
        }
    },
    actions: {
        init() {
            this.password = new Parameter("", false, this.connect)
        },
        connect(psk: string) {
            Loader.show(AppErrorCodes.WifiConnect, 2000)
            WifiConnect(this.ssid, psk).then(() => {
                Loader.close()
            })
        },
        enable() {
            this.getAP()
        },

        getAP() {
            if (this.enabled) {
                this.apList = []
                let newAps: AP[] = []
                this.busy = true
                WifiAPList().then(
                    aps => {
                        aps.forEach(element => {
                            newAps.push({ssid: element})
                        });
                        this.apList = newAps
                        this.busy = false
                    },
                    error => {
                        console.debug(error)
                        setTimeout(() => {
                            this.getAP()
                        }, 200);
                    })
            }
        }
    }
})


export interface AP {
    ssid: string
}
