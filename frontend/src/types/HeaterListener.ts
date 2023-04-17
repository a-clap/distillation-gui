import {NotifyHeaters} from '../../wailsjs/go/backend/Events'

declare type HeaterCallback = () => void;

class heaterListener {
  private static _instance: heaterListener;
  handlers: HeaterCallback[];
  
  public static get Instance()
  {
      // Do you need arguments? Make it a regular static method instead.
      return this._instance || (this._instance = new this());
  }
  private constructor() {
    this.handlers = [];
    NotifyHeaters().then((e) => runtime.EventsOn(e, () => this.#notify()))  
  }

  #notify() {
    this.handlers.forEach((h) => {
      h()
    })
  }

  subscribe(cb: HeaterCallback) {
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
  unsubscribe(cb: HeaterCallback) {
    this.handlers = this.handlers.filter((c) => c != cb);
  }
}

export var HeaterListener = heaterListener.Instance
