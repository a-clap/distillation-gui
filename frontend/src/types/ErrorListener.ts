import { NotifyError } from '../../wailsjs/go/backend/Events'
import { Listener } from './Listener';
import { AppErrorCodes } from '../stores/error_codes';

declare type ErrorCallback = (id: number) => void;

class errorListener {
  private static _instance: errorListener;
  error: Listener;

  public static get Instance() {
    return this._instance || (this._instance = new this());
  }

  private constructor() {
    this.error = new Listener()

    NotifyError().then((ev: string) => {
      return runtime.EventsOn(ev, (...args: any) => {
        this.handle(...args);
      });
    })

  }

  private handle(...args: any) {
    try {
      let t = Number(args[0])
      this.error.notify(t)
    } catch (e) {
      console.log(e)
    }
  }

  sendError(nb: AppErrorCodes) {
    this.handle(nb)
  }


  subscribe(cb: ErrorCallback) {
    console.log("subscribing to error callback")
    this.error.subscribe(cb)
  }

  unsubscribe(cb: ErrorCallback) {
    this.error.unsubscribe(cb)
  }

}

export var ErrorListener = errorListener.Instance
