import React, { Component } from 'react';
import getAssetInfo from '../../../api/getAssetInfo';

class LabelledInput extends Component {
  constructor(props) {
    super(props);
    this.state = {
      Choices: "",
    };
  }

  componentDidMount() {
    this.renderDropdownChoices()
  }

  renderDropdownChoices = async() => {
    this.props.assets.then(
      (response) => {
        console.log("resp", response)
        this.setState({
          Choices: response.results.map((asset, index) => 
            <option key={index} value={JSON.stringify(asset)}>
              {asset.code}:{asset.alias}
            </option>
          )})
      }
    )
  }
  render(props) {
    return (
      <div>
        <label htmlFor={this.props.id}>{this.props.title}</label>
        <select type="text" id={this.props.id} name={this.props.id} onChange={this.props.changeHandler}>
          <option value="">Please select an option.</option>
          {this.state.Choices}
        </select>
      </div>
    );
  }
}

export default LabelledInput;