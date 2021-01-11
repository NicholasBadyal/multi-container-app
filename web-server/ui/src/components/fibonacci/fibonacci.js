import React, { Component } from 'react';
import axios from "axios";

class Fibonacci extends Component {
    state = {
        indices: [],
        values: {},
        currIndex: '',
    }

    componentDidMount() {
        this.fetchIndices();
        this.fetchValues();
    }

    async fetchValues() {
        const values = await axios.get('/api/fibonacci/values')
        this.setState({
            values: values.data
        });
    }

    async fetchIndices() {
        const indices = await axios.get('/api/fibonacci/indices')
        this.setState({
            indices: indices.data
        });
    }

    handleSubmit = async event => {
        event.preventDefault()

        await axios.post('/api/fibonacci/index', {
            index: this.state.currIndex
        });
        this.setState({ index: '' });
    }
}