import React from 'react';
export default class List extends React.Component {
    constructor(props) {
        super(props)
        this.state = {data: [], update: 0, brand: "", model: "", price: 0, status:0, milage:0, id:null}


    }
    componentWillMount(){
        let list = this.props.data;
        this.setState({data: list});
        console.log(list)
    }
    render() {
        return(
        <div>
            <table>
            <thead>
            <tr>
                <th>ID</th>
                <th>Бренд</th>
                <th>Модель</th>
                <th>Цена</th>
                <th>Статус</th>
                <th>Пробег</th>
                <th>Удалить</th>
            </tr>
            </thead>
                <tbody>
                {this.state.data.map((item)=>(
                    <tr key={item.id+this.state.update}>
                        <td>{item.id}</td>
                        <td>{item.brand}</td>
                        <td>{item.model}</td>
                        <td>{item.price}</td>
                        <td>{item.status}</td>
                        <td>{item.milage}</td>
                        <td onClick={()=>this.delProduct(item.id)}>X</td>
                    </tr>
                ))}


                </tbody>
            </table>
            <h1>TEST CREATE</h1>
            <form>
                <label>Бренд<input name={"brand"}  onChange={(event)=>this.changeBrand(event, "brand")}/></label>
                    <br/>
                <label>Модель<input name={"model"}  onChange={(event)=>this.changeBrand(event, "model")}/></label>
                    <br/>
                <label>Прайс<input name={"price"}  onChange={(event)=>this.changeBrand(event, "price")}/></label>
                    <br/>
                <label>Статус<input type={"radio"} name={"status"} value={"0"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    В пути  <input type={"radio"} name={"status"} value={"1"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    На складе<input type={"radio"} name={"status"} value={"2"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    Продан<input type={"radio"} name={"status"} value={"3"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    Снят с продажи
                </label>
                    <br/>
                <label>Пробег<input name={"milage"} onChange={(event)=>this.changeBrand(event, "milage")}/></label>
                    <br/>
                    <button onClick={(event)=>this.save(event)}>Сохранить</button>
            </form>
            <h1>TEST UPDATE</h1>
            <form>
                <label>ID<input name={"id"}  onChange={(event)=>this.changeBrand(event, "id")}/></label>
                <br/>
                <label>Бренд<input name={"brand"}  onChange={(event)=>this.changeBrand(event, "brand")}/></label>
                <br/>
                <label>Модель<input name={"model"}  onChange={(event)=>this.changeBrand(event, "model")}/></label>
                <br/>
                <label>Прайс<input name={"price"}  onChange={(event)=>this.changeBrand(event, "price")}/></label>
                <br/>
                <label>Статус<input type={"radio"} name={"status"} value={"0"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    В пути  <input type={"radio"} name={"status"} value={"1"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    На складе<input type={"radio"} name={"status"} value={"2"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    Продан<input type={"radio"} name={"status"} value={"3"} onChange={(event)=>this.changeBrand(event, "status")}/>
                    Снят с продажи
                </label>
                <br/>
                <label>Пробег<input name={"milage"} onChange={(event)=>this.changeBrand(event, "milage")}/></label>
                <br/>
                <button onClick={(event)=>this.update(event)}>Сохранить</button>
            </form>
        </div>
        )
    }

    delProduct(id) {
        fetch("/delete/"+id)
        this.updateList()

       }
  async  updateList() {
      let res  = await fetch("/list")
      let json = await res.json()
      this.setState({data: json.product_list, update:this.state.update+1})
    }

    changeBrand(e, state) {
        switch (state) {
            case "brand"       : this.setState({brand: e.target.value });
                break;
            case "model": this.setState({model: e.target.value });
                break;
            case "price" : this.setState({price: e.target.value });
                break;
            case "status"    : this.setState({status: e.target.value });
                break;
            case "milage"    : this.setState({milage: e.target.value });
                break;
            case "id"    : this.setState({id: e.target.value });
                break;
        }
    }
   async save(e) {
        e.preventDefault()
        delete this.state.data
        let body = this.state;
        let j = JSON.stringify(this.state)

        let res = await fetch("/new", {
            body: j,
            method: "POST"
        })
        let js = await res.json()
        alert(js.status)

        this.updateList()
    }
    async update(e) {
        e.preventDefault()
        delete this.state.data
        let body = this.state;
        let j = JSON.stringify(this.state)
        let res = await fetch("/put", {
            body: j,
            method: "POST"
        })
        let js = await res.json()
        alert(js.status)
        this.updateList()
    }


}