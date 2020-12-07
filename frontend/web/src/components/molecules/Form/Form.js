import React, { Component } from 'react';
import LabelledInput from '../LabelledInput/LabelledInput';
import getCorridorInfo from '../../../api/getCorridorInfo'
import getVolumeInfo from '../../../api/getVolumeInfo';
class Form extends Component {
  constructor(props) {
    super(props);
    this.state = {
      fromCode: "CENTUS",
      fromIssuer: "GAKMVPHBET4T7DPN32ODVSI4AA3YEZX2GHGNNSBGFNRQ6QEVKFO4MNDZ",
      toCode: "USD",
      toIssuer: "GB2O5PBQJDAFCNM2U2DIMVAEI7ISOYL4UJDTLN42JYYXAENKBWY6OBKZ",
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    const target = event.target;
    const name = target.name;
    const value = target.value
    this.setState({
      [name]: value
    });

  }

  handleSubmit(event) {
    console.log("Submitting with params: ", this.state)
    if (this.state.fromCode !== ""  && this.state.fromIssuer !== "" && this.state.toCode !== "" && this.state.toIssuer !== "") {
      getCorridorInfo(this.props.baseUrl, this.state.fromCode, this.state.fromIssuer, this.state.toCode, this.state.toIssuer).then(
        response => {this.props.handler(response)})
    } else if (this.state.fromCode !== ""  && this.state.fromIssuer !== "" && this.state.toCode === "" && this.state.toIssuer === "") {
      getVolumeInfo(this.props.baseUrl, this.state.fromCode, this.state.fromIssuer, "true").then(
        response => {this.props.handler(response)})
    } else if (this.state.fromCode === ""  && this.state.fromIssuer === "" && this.state.toCode !== "" && this.state.toIssuer !== "") {
      getVolumeInfo(this.props.baseUrl, this.state.toCode, this.state.toIssuer, "").then(
        response => {this.props.handler(response)})
    } else {
      return {"results": "parameters are malformed"}
    }
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
            <LabelledInput title="From Code" id="fromCode" value={this.state.fromCode}/>
            <LabelledInput title="From Issuer" id="fromIssuer" value={this.state.fromIssuer}/>
            <LabelledInput title="To Code" id="toCode" value={this.state.toCode}/>
            <LabelledInput title="To Issuer" id="toIssuer" value={this.state.toIssuer}/>
            <input type="submit" value="Submit"></input>
      </form>
    );
  }
}

export default Form;