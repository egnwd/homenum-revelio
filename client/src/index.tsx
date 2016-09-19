/// <reference path="sse.d.ts" />
/// <reference path="./protocol.d.ts"/>

import * as React from "react"
import * as ReactDOM from "react-dom"

import { PeopleTable } from "./components/PeopleTable"
import { Person } from "./components/Person"
import { Model } from "./models/Model"

import "./style/main"

let model = new Model()
let source = new EventSource("/updates")

source.addEventListener("message", function(event) {
    let data = JSON.parse((event as sse.IOnMessageEvent).data)
    let message = (data as IMessage)
    let people = message.people
    
    model.update(people)
})

var render = function() {
  ReactDOM.render(
      <PeopleTable model={model}/>,
      document.getElementById("main")
  )
}

model.subscribe(render)
render()
