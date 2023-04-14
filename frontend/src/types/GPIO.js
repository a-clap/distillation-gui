import { GetValues } from "../../wailsjs/go/main/App"

class Rectangle {
  constructor(height, width) {
    this.height = height;
    this.width = width;
  }

  callGo() {
    GetValues().then((result) => {
        result.forEach((elem) => {
          console.log(elem.name)
        })
    })
  }

}

export {Rectangle}