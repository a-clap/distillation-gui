declare type Callback = (...args: any) => void;

export class Listener {
  callbacks: Callback[];
  
  constructor(eventName: string) {
    this.callbacks = [];
    runtime.EventsOn(eventName, (...args) => {
        this.#notify(args)
    })  
  }

  #notify(...args: any) {
    this.callbacks.forEach((h) => {
      h(args)
    })
  }

  subscribe(cb: Callback) {
    let found = this.callbacks.some((c) => {
      return c.toString() == cb.toString();
    });

    if (!found) {
      this.callbacks.push(cb)
    } 

  }
  unsubscribe(cb: Callback) {
    this.callbacks = this.callbacks.filter((c) => c != cb);
  }
}
