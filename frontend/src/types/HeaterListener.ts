import { NotifyHeaters } from '../../wailsjs/go/backend/Events'
import { parameters } from '../../wailsjs/go/models';
import { Listener } from './Listener';

declare type HeaterCallback = (h: parameters.Heater) => void;

class heaterListener {
  private static _instance: heaterListener;
  heaters: Listener;

  public static get Instance() {
    return this._instance || (this._instance = new this());
  }

  private constructor() {
    this.heaters = new Listener()

    NotifyHeaters().then((ev) => {
      return runtime.EventsOn(ev, (...args: any) => {
        this.handle(...args);
      });
    })
  }

  private handle(...args: any) {
    if (args.length != 1) {
      console.log("Expected single element")
      return
    }
    try {
      let t = new parameters.Heater(args[0])
      this.heaters.notify(t)
    } catch (e) {
      console.log(e)
    }
  }

  subscribe(cb: HeaterCallback) {
    this.heaters.subscribe(cb)
  }
  unsubscribe(cb: HeaterCallback) {
    this.heaters.unsubscribe(cb)
  }
}

export var HeaterListener = heaterListener.Instance
