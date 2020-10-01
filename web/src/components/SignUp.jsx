import React from 'react';
import axios from 'axios';
import './LoginForm.css'


export default class SignUp extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            email: '',
            password: '',
        };
    }

    handleSignUp = (event) => {
        event.preventDefault();

        let apiUrl = process.env.REACT_APP_PORT // read from config        
        let payload = {
            "email": this.state.email,
            "password": this.state.password
        }

        axios.post(apiUrl + 'accounts', payload).then(res => {
            this.props.onLogin(true)
            console.log("Sign up success!")
        }).catch(err => {
            console.log(err)
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
                <form className="loginForm" onSubmit={this.handleSignUp}>
                    <h1>Register An Account</h1>
                    <div>
                        <label>Email</label>
                        <input type="email" id="email" name='email' className="field" onChange={this.handleInput} />

                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <input type="password" id="password" name='password' className="field" onChange={this.handleInput} />

                    </div>

                    <input type="submit" value="SignUp" className="button block"></input>
                </form>
            </div>
        );
    }
}