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
      console.log(`Finding volume between ${this.state.fromCode}->${this.state.toCode}`)

      getCorridorInfo(this.props.baseUrl, this.state.fromCode, this.state.fromIssuer, this.state.toCode, this.state.toIssuer).then(
        response => {this.props.handler(response)})

    } else if (this.state.fromCode !== ""  && this.state.fromIssuer !== "" && this.state.toCode === "" && this.state.toIssuer === "") {
      console.log(`Finding volume from ${this.state.fromCode}`)

      getVolumeInfo(this.props.baseUrl, this.state.fromCode, this.state.fromIssuer, "true").then(
        response => {this.props.handler(response)})
    } else if (this.state.fromCode === ""  && this.state.fromIssuer === "" && this.state.toCode !== "" && this.state.toIssuer !== "") {
      console.log(`Finding volume to ${this.state.toCode}`)

      getVolumeInfo(this.props.baseUrl, this.state.toCode, this.state.toIssuer, "").then(
        response => {this.props.handler(response)})
    } else {
       this.props.handler({"results": "parameters are malformed"})
    }

    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
            <LabelledInput title="From Code" id="fromCode" value={this.state.fromCode} changeHandler={this.handleChange}/>
            <LabelledInput title="From Issuer" id="fromIssuer" value={this.state.fromIssuer} changeHandler={this.handleChange}/>
            <LabelledInput title="To Code" id="toCode" value={this.state.toCode} changeHandler={this.handleChange}/>
            <LabelledInput title="To Issuer" id="toIssuer" value={this.state.toIssuer} changeHandler={this.handleChange}/>
            <input type="submit" value="Submit"></input>
      </form>
    );
  }
}

export default Form;