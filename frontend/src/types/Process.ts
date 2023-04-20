import { PhasesSetPhaseCount } from "../../wailsjs/go/backend/Backend";
import Parameter from "./Parameter";

export class ProcessPhase {

}

export class Process {
    phases: ProcessPhase[];
    phaseCount: Parameter;

    constructor() {
        this.phases = []
        this.phaseCount = new Parameter(0, false, this.setPhaseCount)
    }

    private setPhaseCount(cnt: number) {
        console.log("set phase count " + cnt)
        PhasesSetPhaseCount(cnt)
   }
}