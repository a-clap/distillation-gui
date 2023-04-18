import {NotifyDSConfig, NotifyDSTemperature} from '../../wailsjs/go/backend/Events'

declare type DSCallbackConfig = () => void;
declare type DSCallbackTemperature = () => void;

class dsListener {
  private static _instance: dsListener;
  configHandlers: DSCallbackConfig[];
  temperatureHandlers: DSCallbackTemperature[];
  
  public static get Instance()
  {
      // Do you need arguments? Make it a regular static method instead.
      return this._instance || (this._instance = new this());
  }
  private constructor() {
    this.configHandlers = [];
    this.temperatureHandlers = [];
    NotifyDSConfig().then((e) => runtime.EventsOn(e, () => this.#notifyConfig()))  
    NotifyDSTemperature().then((e) => runtime.EventsOn(e, () => this.#notifyTemperature()))  
  }

  #notifyConfig() {
    this.configHandlers.forEach((h) => {
      h()
    })
  }
  
  #notifyTemperature() {
    this.temperatureHandlers.forEach((h) => {
      h()
    })
  }

  subscribe(cb: DSCallback) {
    let found = this.handlers.some((c) => {
      return c.toString() == cb.toString();
    });

    if (!found) {
      console.log("subscribe");
      this.handlers.push(cb)
    } else {
      console.log("already exists");
    }

  }
  unsubscribe(cb: DSCallback) {
    this.handlers = this.handlers.filter((c) => c != cb);
  }
}

export var DSListener = dsListener.Instance
