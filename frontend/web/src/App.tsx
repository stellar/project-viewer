import React, { Component } from "react";
import "./App.css";
import Form from "./components/molecules/Form/Form";
import DataTable from "./components/molecules/DataTable/DataTable";

class App extends Component {
  // @ts-ignore
  constructor(props: any) {
    super(props);
    this.state = {
      data: "",
    };
    this.queryResultHandler = this.queryResultHandler.bind(this);
  }

  // @ts-ignore
  queryResultHandler(response) {
    this.setState({
      data: JSON.stringify(response.results),
    });
    console.log("From parent, now our state is", this.state);
  }

  render() {
    return (
      <div className="App">
        <Form
          handler={this.queryResultHandler}
          baseUrl="http://localhost:8080"
        />
        {/* @ts-ignore */}
        <DataTable data={this.state.data} />
      </div>
    );
  }
}

export default App;
