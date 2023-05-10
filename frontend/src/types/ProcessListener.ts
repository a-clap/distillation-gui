import {
    NotifyGlobalConfig,
    NotifyPhasesConfig,
    NotifyPhasesPhaseConfig,
    NotifyPhasesPhaseCount,
    NotifyPhasesStatus,
    NotifyPhasesValidate,
} from '../../wailsjs/go/backend/Events'
import { distillation, process } from '../../wailsjs/go/models';
import { Listener } from './Listener';

declare type PhaseCallbackConfig = (c: distillation.ProcessPhaseConfig) => void;
declare type PhaseCallbackValidate = (v: distillation.ProcessConfigValidation) => void;
declare type PhaseCallbackPhaseConfig = (n: number, v: distillation.ProcessPhaseConfig) => void;
declare type PhaseCallbackPhaseCount = (v: distillation.ProcessPhaseCount) => void;
declare type PhaseCallbackStatus = (v: distillation.ProcessStatus) => void;
declare type PhaseCallbackGlobalConfig = (v: process.Config) => void;

class processListener {
    private static _instance: processListener;
    config: Listener;
    validate: Listener;
    phaseConfig: Listener;
    phaseCount: Listener;
    status: Listener;
    globalConfig: Listener;

    public static get Instance() {
        return this._instance || (this._instance = new this());
    }

    private constructor() {
        this.config = new Listener()
        this.validate = new Listener()
        this.phaseConfig = new Listener()
        this.phaseCount = new Listener()
        this.status = new Listener()
        this.globalConfig = new Listener()

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
                console.log("NotifyPhasesPhaseCount")
                this.NotifyPhasesPhaseCount(...args);
            });
        })

        NotifyPhasesStatus().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyPhasesStatus(...args);
            });
        })

        NotifyGlobalConfig().then((ev: string) => {
            return runtime.EventsOn(ev, (...args: any) => {
                this.NotifyGlobalConfig(...args);
            });
        })
    }

    subscribeGlobalConfig(cb: PhaseCallbackGlobalConfig) {
        this.globalConfig.subscribe(cb)
    }

    unsubscribeGlobalConfig(cb: PhaseCallbackGlobalConfig) {
        this.globalConfig.unsubscribe(cb)
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
            console.log(args)
            let t = new distillation.ProcessConfigValidation(args[0])
            this.validate.notify(t)
        } catch (e) {
            console.log(e)
        }
    }
    private NotifyPhasesPhaseConfig(...args: any) {
        try {
            let n = Number(args[0])
            let t = new distillation.ProcessPhaseConfig(args[1])
            this.phaseConfig.notify(n, t)
        } catch (e) {
            console.log(e)
        }
    }

    private NotifyPhasesPhaseCount(...args: any) {
        try {
            let t = new distillation.ProcessPhaseConfig(args[0])
            this.phaseCount.notify(t)
        } catch (e) {
            console.log(e)
        }
    }

    private NotifyPhasesStatus(...args: any) {
        try {
            let a: distillation.ProcessStatus = args[0]
            this.status.notify(a)
        } catch (e) {
            console.log(e)
        }
    }

    private NotifyGlobalConfig(...args: any) {
        try {
            let t = new process.Config(args[0])
            this.globalConfig.notify(t)
        } catch (e) {
            console.log(e)
        }
    }

}

export var ProcessListener = processListener.Instance
