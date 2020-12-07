import React, { Component } from 'react';

class DataTable extends Component {
    render(props) {
        return (
            <div>
            <p>{this.props.data}</p>
            </div>
        );
    }
}

export default DataTable;