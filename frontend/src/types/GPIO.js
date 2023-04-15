export class GPIO {
  name = "";

  constructor(name, activeLow) {
    this.name = name;
    this.activeLow = activeLow;
  }

  setActiveLevel() {
    console.log("setting active level " + this.activeLow)
  }

  setState(newState) {
    console.log("setting new state" + newState)
  }

}

// export {GPIO}