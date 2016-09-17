/// <reference path="sse.d.ts" />

import * as React from "react"
import * as ReactDOM from "react-dom"

import { IMessage } from "./protocol";
import { PeopleTable } from "./components/PeopleTable"
import { Person } from "./components/Person"

import "./style/main"

var people = [new Person("id", "Elliot", true),
              new Person("id", "Jonathan", true),
              new Person("id", "Florian", false)]

let source = new EventSource("/updates")

source.addEventListener("message", function(event) {
    let data = JSON.parse((event as sse.IOnMessageEvent).data)
    let message = (data as IMessage)

    console.log("Got message: " + message.people)
})

ReactDOM.render(
    <PeopleTable people={people}/>,
    document.getElementById("main")
);
