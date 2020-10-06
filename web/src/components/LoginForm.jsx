import React from 'react';
import axios from 'axios';
import './LoginForm.css'

import config from "../config"

import { Link } from "react-router-dom";

export default class loginForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            email: '',
            password: '',
            errorLogin: false
        };
    }

    handleLogin = (event) => {
        console.log(this.state.email, this.state.password);
        event.preventDefault();

        let apiUrl = config.url.api_url // read from config
        console.log("server url: ", apiUrl)
        let payload = {
            "email": this.state.email,
            "password": this.state.password
        }

        axios.post(apiUrl + 'login', payload).then(res => {
            this.props.onLogin(true)
            console.log("Login success!")
        }).catch(err => {
            console.log(err)
            this.setState({ errorLogin: true })
        });
    }

    handleInput = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        });
    }

    render() {
        return (
            <div>
                {this.state.errorLogin &&
                    <div className="alertMsg">
                        Incorrect username or password.
                    </div>
                }
                <form className="loginForm" onSubmit={this.handleLogin}>
                    <h1>Sign Into Your Account</h1>
                    <div>
                        <label>Email</label>
                        <input type="email" id="email" name='email' className="field" onChange={this.handleInput} />

                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <input type="password" id="password" name='password' className="field" onChange={this.handleInput} />

                    </div>

                    <input type="submit" value="Login to my Dashboard" className="button block"></input>
                </form>


                <div className="signup">
                    <Link to={process.env.PUBLIC_URL + "/signup"}>Create an account</Link>
                </div>

            </div>
        );
    }
}