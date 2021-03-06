import React from 'react';
import axios from 'axios';

import MsgBoard from './MsgBoard'
import NavBar from './NavBar'
import ProgressBar from './ProgressBar'

import { eraseCookie } from '../api/cookie'
import config from "../config"

export default class DashBoard extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            sessionToken: '',
            msg: ''
        };
    }

    handleLogout = () => {
        let apiUrl = config.url.api_auth_url // read from config

        axios.post(apiUrl + 'logout').then(res => {
            this.props.onLogin(false)
            eraseCookie('sid')
        }).catch(err => console.log(err));
    }

    handleMsg = (message) => {
        this.setState({ msg: message });
    }

    render() {
        return (
            <>
                <NavBar onLogout={this.handleLogout} />
                <MsgBoard msg={this.state.msg} />
                <ProgressBar onMsg={this.handleMsg} />
            </>
        );
    }
}