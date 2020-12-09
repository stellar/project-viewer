import React, { Component } from 'react';
import LabelledInput from '../LabelledInput/LabelledInput';
import getCorridorInfo from '../../../api/getCorridorInfo'
import getVolumeInfo from '../../../api/getVolumeInfo';
import getAssetInfo from '../../../api/getAssetInfo';
class Form extends Component {
  constructor(props) {
    super(props);
    this.state = {
      fromAsset: "",
      toAsset: "",
      assets: getAssetInfo(this.props.baseUrl),
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    const target = event.target;
    const name = target.name;
    const value = target.value
    console.log(target)
    this.setState({
      [name]: value
    });
  }

  handleSubmit(event) {
    console.log("Submitting with params: ", this.state)
    if (this.state.fromAsset !== ""  && this.state.toAsset !== "") {
      console.log(`Finding volume between ${this.state.fromAsset.code}->${this.state.toAsset.code}`)

      let fromAsset = JSON.parse(this.state.fromAsset)
      let toAsset = JSON.parse(this.state.toAsset)
      getCorridorInfo(this.props.baseUrl, fromAsset.code, fromAsset.issuer, toAsset.code, toAsset.issuer).then(
        response => {this.props.handler(response)})

    } else if (this.state.fromAsset !== ""  && this.state.toAsset === "") {
      let fromAsset = JSON.parse(this.state.fromAsset)

      console.log(`Finding volume from ${fromAsset.code}`)

      getVolumeInfo(this.props.baseUrl, fromAsset.code, fromAsset.issuer, "true").then(
        response => {this.props.handler(response)})
    } else if (this.state.fromAsset === ""  && this.state.toAsset !== "") {
      let toAsset = JSON.parse(this.state.toAsset)

      console.log(`Finding volume to ${toAsset.code}`)

      getVolumeInfo(this.props.baseUrl, toAsset.code, toAsset.issuer, "").then(
        response => {this.props.handler(response)})
    } else {
       this.props.handler({"results": "parameters are malformed"})
    }

    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
            <LabelledInput title="From" id="fromAsset" value={this.state.fromAsset} changeHandler={this.handleChange} assets={this.state.assets}/>
            <LabelledInput title="To" id="toAsset" value={this.state.toAsset} changeHandler={this.handleChange} assets={this.state.assets}/>
            <input type="submit" value="Submit"></input>
      </form>
    );
  }
}

export default Form;