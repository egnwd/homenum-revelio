export interface IMessage {
  people: IPerson[]
}

export interface IPerson {
  mac: string
  name: string
  status: boolean
}
