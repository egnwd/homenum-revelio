import * as React from "react"
import { Person, PersonItem } from "./Person"
import { Model } from "../models/Model";
import "style/modules/table"

export interface PeopleTableProps { model: Model; }

export class PeopleTable extends React.Component<PeopleTableProps, {}> {
    render() {
      var people = this.props.model.people
      var rows = people.map(function (person) {
        return (
          <PersonItem
            mac={person.mac}
            name={person.name}
            status={person.status }
          />
        )
      }, this)

      return  (
        <table>
          <tbody>
            { rows }
          </tbody>
        </table>
      )
    }
}
