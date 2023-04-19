declare type Callback = (...args: any) => void;

export class Listener {
  callbacks: Callback[];

  constructor() {
    this.callbacks = [];
  }

  notify(...args: any) {
    this.callbacks.forEach((h) => {
      h(...args)
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
