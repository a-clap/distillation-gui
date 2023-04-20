import {
    NotifyPhasesConfig,
    NotifyPhasesPhaseConfig,
    NotifyPhasesPhaseCount,
    NotifyPhasesStatus,
    NotifyPhasesValidate,
} from '../../wailsjs/go/backend/Events'
import { distillation } from '../../wailsjs/go/models';
import { Listener } from './Listener';

declare type PhaseCallbackConfig = (c: distillation.ProcessPhaseConfig) => void;
declare type PhaseCallbackValidate = (v: distillation.ProcessConfigValidation) => void;
declare type PhaseCallbackPhaseConfig = (n: number, v: distillation.ProcessPhaseConfig) => void;
declare type PhaseCallbackPhaseCount = (v: distillation.ProcessPhaseCount) => void;
declare type PhaseCallbackStatus = (v: distillation.ProcessStatus) => void;

class processListener {
    private static _instance: processListener;
    config: Listener;
    validate: Listener;
    phaseConfig: Listener;
    phaseCount: Listener;
    status: Listener;

    public static get Instance() {
        return this._instance || (this._instance = new this());
    }

    private constructor() {
        this.config = new Listener()
        this.validate = new Listener()
        this.phaseConfig = new Listener()
        this.phaseCount = new Listener()
        this.status = new Listener()

        NotifyPhasesConfig().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyPhasesConfig(...args);
            });
        })

        NotifyPhasesValidate().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyPhasesValidate(...args);
            });
        })

        NotifyPhasesPhaseConfig().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyPhasesPhaseConfig(...args);
            });
        })

        NotifyPhasesPhaseCount().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyPhasesPhaseCount(...args);
            });
        })

        NotifyPhasesStatus().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyPhasesStatus(...args);
            });
        })
    }

    subscribeConfig(cb: PhaseCallbackConfig) {
        this.config.subscribe(cb)
    }

    unsubscribeConfig(cb: PhaseCallbackConfig) {
        this.config.unsubscribe(cb)
    }

    subscribeValidate(cb: PhaseCallbackValidate) {
        this.validate.subscribe(cb)
    }

    unsubscribeValidate(cb: PhaseCallbackValidate) {
        this.validate.unsubscribe(cb)
    }

    subscribePhaseConfig(cb: PhaseCallbackPhaseConfig) {
        this.phaseConfig.subscribe(cb)
    }

    unsubscribePhaseConfig(cb: PhaseCallbackPhaseConfig) {
        this.phaseConfig.unsubscribe(cb)
    }

    subscribePhaseCount(cb: PhaseCallbackPhaseCount) {
        this.phaseCount.subscribe(cb)
    }

    unsubscribePhaseCount(cb: PhaseCallbackPhaseCount) {
        this.phaseCount.unsubscribe(cb)
    }

    subscribeStatus(cb: PhaseCallbackStatus) {
        this.status.subscribe(cb)
    }

    unsubscribeStatus(cb: PhaseCallbackStatus) {
        this.status.unsubscribe(cb)
    }

    private NotifyPhasesConfig(...args: any) {
        try {
            let t = new distillation.ProcessPhaseConfig(args[0])
            this.config.notify(t)
        } catch (e) {
            console.log(e)
        }
    }

    private NotifyPhasesValidate(...args: any) {
        try {
            let t = new distillation.ProcessConfigValidation(args[0])
            this.config.notify(t)
        } catch (e) {
            console.log(e)
        }
    }
    private NotifyPhasesPhaseConfig(...args: any) {
        try {
            let n = Number(args[0])
            let t = new distillation.ProcessConfigValidation(args[1])
            this.config.notify(n, t)
        } catch (e) {
            console.log(e)
        }
    }

    private NotifyPhasesPhaseCount(...args: any) {
        try {
            let t = new distillation.ProcessPhaseConfig(args[0])
            this.config.notify(t)
        } catch (e) {
            console.log(e)
        }
    }

    private NotifyPhasesStatus(...args: any) {
        try {
            let t = new distillation.ProcessStatus(args[0])
            this.config.notify(t)
        } catch (e) {
            console.log(e)
        }
    }

}

export var ProcessListener = processListener.Instance
