import * as React from "react"
import { Person, PersonItem } from "./Person"
import "style/modules/table"

export interface PeopleTableProps { people: Person[]; }

export class PeopleTable extends React.Component<PeopleTableProps, {}> {
    render() {
      var people = this.props.people
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
