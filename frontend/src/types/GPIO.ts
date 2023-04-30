import { GPIOSetState } from "../../wailsjs/go/backend/Backend";
import { GPIOSetActiveLevel } from "../../wailsjs/go/backend/Backend";
export class GPIO {
  name: string;
  private activeLevel_: boolean;
  private state_: boolean;
  constructor(name: string, activeLevel: number, state: boolean = false) {
    this.name = name;
    this.activeLevel_ = activeLevel > 0;
    this.state_ = state
  }

  get activeLevel(): boolean {
    return this.activeLevel_
  }

  set activeLevel(value: any) {
    let v: number = 0
    if (typeof value === 'string') {
      v = value == "true" ? 1 : 0
    } else if(typeof value === 'boolean') {
      v = value ? 1 : 0
    } else {
      v = value
    }

    this.activeLevel_ = v > 0
    GPIOSetActiveLevel(this.name, v)

  }

  set state(value: boolean) {
    this.state_ = value
    GPIOSetState(this.name, this.state_)
  }

  get state(): boolean {
    return this.state_
  }

}