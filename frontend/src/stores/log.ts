import { defineStore } from "pinia";
import { i18n } from "../main";

function padTo2Digits(num: number) {
    return num.toString().padStart(2, '0');
}

function formatDate(date: Date) {
    return (
        [
            padTo2Digits(date.getHours()),
            padTo2Digits(date.getMinutes()),
            padTo2Digits(date.getSeconds()),
        ].join(':') +
        ' ' +
        [
            padTo2Digits(date.getDate()),
            padTo2Digits(date.getMonth() + 1),
            date.getFullYear(),
        ].join('-')
    );
}

export interface Column {
    title: string;
    width: number;
}

export const useLogStore = defineStore('logs', {
    state: () => {
        return {
            errors: [] as string[],
            logData: [] as string[][],
            columns: [] as Column[],
        }
    },
    actions: {
        init() {
            this.columns = []
            this.columns.push({ title: i18n.global.t('logs.id'), width: 70 })
            this.columns.push({ title: i18n.global.t('logs.timestamp'), width: 150 })
            this.columns.push({ title: i18n.global.t('logs.message'), width: 600 })

        },
        add(errCode: number, msg: string) {
            let elem: string[] = [
                errCode.toString(),
                formatDate(new Date()),
                msg
            ]
            this.logData.unshift(elem)
        },
    }
})
