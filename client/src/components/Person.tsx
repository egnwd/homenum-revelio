import * as React from "react"
import "style/modules/person"

export interface IPerson { mac: string ; name: string; status: boolean }

export class Person implements IPerson {
  public mac: string
  public name: string
  public status: boolean

  constructor(mac: string, name: string, status: boolean) {
    this.mac = mac
    this.name = name
    this.status = status
  }
}

export class PersonItem extends React.Component<IPerson, {}> {
  render() {
    var name = this.props.name
    var status_name = this.props.status ? "In" : "Out"
    
    return (
      <tr>
        <td>
          { name }
        </td>
        <td>
          { status_name }
        </td>
      </tr>
    )
  }
}
