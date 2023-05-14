import { defineStore } from "pinia";
import { useDSStore } from "./ds";
import { usePTStore } from "./pt";

export const useNameStore = defineStore('appNames', {
    state: () => {
        return {
            ds: useDSStore(),
            pt: usePTStore(),
        }
    },
    actions: {
        init() {
        },
        name_to_id(name: string): [string, boolean] {
            let ptIDX = this.pt.pt.findIndex(i => i.name.value == name)
            let dsIDX = this.ds.ds.findIndex(i => i.name.value == name)

            // Found pt and DS, not unique name
            if (ptIDX != -1 && dsIDX != -1) {
                return [name, false]
            } else if (ptIDX == -1 && dsIDX == -1) {
                // Found nothing
                return [name, false]
            }

            // PT
            if (ptIDX != -1) {
                return [this.pt.pt[ptIDX].id, true]
            }
            // So DS left
            return [this.ds.ds[dsIDX].id, true]
        },
        id_to_name(id: string): [string, boolean] {
            let ptIDX = this.pt.pt.findIndex(i => i.id == id)
            let dsIDX = this.ds.ds.findIndex(i => i.id == id)

            // Found pt and DS, not unique name
            if (ptIDX != -1 && dsIDX != -1) {
                return [id, false]
            } else if (ptIDX == -1 && dsIDX == -1) {
                // Found nothing
                return [id, false]
            }

            // PT
            if (ptIDX != -1) {
                return [this.pt.pt[ptIDX].name.value.toString(), true]
            }
            // So DS left
            return [this.ds.ds[dsIDX].name.value.toString(), true]
        }

    }
})