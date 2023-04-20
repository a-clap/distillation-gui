import { NotifyGPIO } from '../../wailsjs/go/backend/Events'
import { parameters } from '../../wailsjs/go/models';
import { Listener } from './Listener';

declare type GPIOCallback = (ds: parameters.GPIO) => void;

class gpioListener {
  private static _instance: gpioListener;
  config: Listener;
  temperature: Listener

  public static get Instance() {
    return this._instance || (this._instance = new this());
  }

  private constructor() {
    this.config = new Listener()
    this.temperature = new Listener()

    NotifyGPIO().then((ev: string) => {
      return runtime.EventsOn(ev, (...args: any) => {
        this.handle(...args);
      });
    })

  }

  private handle(...args: any) {
    try {
      let t = new parameters.GPIO(args[0])
      this.config.notify(t)
    } catch (e) {
      console.log(e)
    }

  }

  subscribe(cb: GPIOCallback) {
    this.config.subscribe(cb)
  }

  unsubscribe(cb: GPIOCallback) {
    this.config.unsubscribe(cb)
  }

}

export var GPIOListener = gpioListener.Instance
