/// <reference path="../protocol.d.ts"/>

export class Model implements IModel {
  public people: Array<IPerson>
  public onChanges: Array<any>

  constructor() {
      this.people = []
      this.onChanges = []
  }

  public subscribe(onChange: any) {
    this.onChanges.push(onChange)
  }

  public inform() {
    this.onChanges.forEach(function (cb) { cb() })
  }

  public update(people: Array<IPerson>) {
    this.people = people
    this.inform()
  }
}
