interface IMessage {
  people: IPerson[]
}

interface IPerson {
  mac: string
  name: string
  status: boolean
}

interface IModel {
  people: Array<IPerson>
  onChanges: Array<any>
  subscribe: (onChange: any) => void
  inform: () => void
  update: (people: Array<IPerson>) => void
}
