export class GPIO {
  name_: string;
  activeLevel_: boolean;
  state_: boolean;
  constructor(name: string, activeLow: boolean) {
    this.name_ = name;
    this.activeLevel_ = activeLow;
    this.state_ = false
  }

  get name() {
    return this.name_
  }


  get activeLevel(): boolean {
    return this.activeLevel_
  }

  set activeLevel(value: boolean) {
    this.activeLevel_ = value
    console.log("setting active level " + this.activeLevel_)

  }

  set state(value: boolean) {
    this.state_ = value
    console.log("setting state level " + this.state_)
  }

  get state(): boolean {
    return this.state_
  }

}