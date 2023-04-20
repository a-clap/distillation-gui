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

  set activeLevel(value: boolean) {
    this.activeLevel_ = value
    GPIOSetActiveLevel(this.name, this.activeLevel_ ? 1 : 0)

  }

  set state(value: boolean) {
    this.state_ = value
    GPIOSetState(this.name, this.state_)
  }

  get state(): boolean {
    return this.state_
  }

}