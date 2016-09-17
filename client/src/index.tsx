import * as React from "react"
import * as ReactDOM from "react-dom"

import { PeopleTable } from "./components/PeopleTable"
import { Person } from "./components/Person"

import "./style/main"

var people = [new Person("id", "Elliot", true)]

ReactDOM.render(
    <PeopleTable people={people}/>,
    document.getElementById("main")
);
