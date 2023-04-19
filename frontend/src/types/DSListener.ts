import { NotifyDSConfig, NotifyDSTemperature } from '../../wailsjs/go/backend/Events'
import { parameters } from '../../wailsjs/go/models';
import { Listener } from './Listener';

declare type DSCallbackConfig = (ds: parameters.DS) => void;
declare type DSCallbackTemperature = (t: parameters.Temperature) => void;

class dsListener {
  private static _instance: dsListener;
  config: Listener;
  temperature: Listener

  public static get Instance() {
    return this._instance || (this._instance = new this());
  }

  private constructor() {
    this.config = new Listener()
    this.temperature = new Listener()

    NotifyDSConfig().then((ev: string) => {
      return runtime.EventsOn(ev, (...args: any) => {
        this.handleConfig(...args);
      });
    })

    NotifyDSTemperature().then((ev: string) => {
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
      let t = new parameters.DS(args[0])
      this.config.notify(t)
    } catch (e) {
      console.log(e)
    }

  }

  subscribeConfig(cb: DSCallbackConfig) {
    this.config.subscribe(cb)
  }

  unsubscribeConfig(cb: DSCallbackConfig) {
    this.config.unsubscribe(cb)
  }

  subscribeTemperature(cb: DSCallbackTemperature) {
    this.temperature.subscribe(cb)
  }

  unsubscribeTemperature(cb: DSCallbackTemperature) {
    this.temperature.unsubscribe(cb)
  }

}

export var DSListener = dsListener.Instance
