import React, { Component } from 'react';

class LabelledInput extends Component {
  render(props) {
    return (
      <div>
        <label htmlFor={this.props.id}>{this.props.title}</label>
        <input type="text" id={this.props.id} name={this.props.id} defaultValue={this.props.value}></input>
      </div>
    );
  }
}

export default LabelledInput;