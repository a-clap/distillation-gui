export class GPIO {
  constructor(name, activeLow) {
    this.name_ = name;
    this.activeLevel_ = activeLow;
    this.state_ = false
  }

  get name() {
    return this.name_
  }


  get activeLevel() {
    return this.activeLevel_
  }

  set activeLevel(value) {
    this.activeLevel_ = value
    console.log("setting active level " + this.activeLevel_)

  }

  set state(value) {
    this.state_ = value
    console.log("setting state level " + this.state_)
  }

  get state() {
    return this.state_
  }

}