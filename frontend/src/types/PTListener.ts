import { NotifyPTConfig, NotifyPTTemperature } from '../../wailsjs/go/backend/Events'
import { parameters } from '../../wailsjs/go/models';
import { Listener } from './Listener';

declare type PTCallbackConfig = (PT: parameters.PT) => void;
declare type PTCallbackTemperature = (t: parameters.Temperature) => void;

class ptListener {
  private static _instance: ptListener;
  config: Listener;
  temperature: Listener

  public static get Instance() {
    return this._instance || (this._instance = new this());
  }

  private constructor() {
    this.config = new Listener()
    this.temperature = new Listener()

    NotifyPTConfig().then((ev) => {
      return runtime.EventsOn(ev, (...args: any) => {
        this.handleConfig(...args);
      });
    })

    NotifyPTTemperature().then((ev) => {
      return runtime.EventsOn(ev, (...args: any) => {
        this.handleTemperature(...args);
      });
    })
  }

  private handleTemperature(...args: any) {
    if (args.length != 1) {
      console.log("Expected single element")
      return
    }
    try {
      let t = new parameters.Temperature(args[0])
      this.temperature.notify(t)
    } catch (e) {
      console.log(e)
    }
  }

  private handleConfig(...args: any) {
    try {
      let t = new parameters.PT(args[0])
      this.config.notify(t)
    } catch (e) {
      console.log(e)
    }

  }

  subscribeConfig(cb: PTCallbackConfig) {
    this.config.subscribe(cb)
  }

  unsubscribeConfig(cb: PTCallbackConfig) {
    this.config.unsubscribe(cb)
  }

  subscribeTemperature(cb: PTCallbackTemperature) {
    this.temperature.subscribe(cb)
  }

  unsubscribeTemperature(cb: PTCallbackTemperature) {
    this.temperature.unsubscribe(cb)
  }

}

export var PTListener = ptListener.Instance
