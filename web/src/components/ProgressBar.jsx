import axios from 'axios';
import React from 'react';
import Cookies from 'js-cookie'

import './ProgressBar.css'

export default class ProgressBar extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            ws: null,
            plan: 'Startup',
            users: 0,
            capacity: 100,
            addition: ''
        };
    }

    componentDidMount() {
        this.fetchData();
        this.connect();
    }

    componentWillUnmount() {
        if (this.state.ws)
            this.state.ws.close();
    }

    fetchData = () => {
        let apiUrl = process.env.REACT_APP_AUTH_PORT // read from config
        let userCookie = JSON.parse(Cookies.get('user'))
        let tenantID = userCookie['tenant_id']

        axios.get(apiUrl + 'tenants/' + tenantID).then(res => {
            console.log(res.data)
            this.setState({ capacity: res.data['capacity'], users: res.data['users'] })
        }).catch(err => console.log(err));
    }

    connect = () => {
        var ws = new WebSocket("ws://localhost:8088/auth/ws/tenantUsers"); // read from config
        let that = this; // cache the this
        var connectInterval;

        ws.onopen = () => {
            console.log("connect websocket main component");
            ws.onmessage = this.update;
            this.setState({ ws: ws });

            that.timeout = 500; // reset timer to 250 on open of websocket connection 
            clearTimeout(connectInterval); // clear Interval on on open of websocket connection
        };

        ws.onclose = e => {
            if (that.timeout > 3000) {
                console.log('stop retry')
                clearTimeout(connectInterval);
                return
            }

            console.log(
                `Socket is closed. Reconnect will be attempted in ${Math.min(
                    10000 / 1000,
                    (that.timeout + that.timeout) / 1000
                )} second.`,
                e.reason
            );

            that.timeout = that.timeout + that.timeout; //increment retry interval
            connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)); //call check function after timeout
        };

        ws.onerror = err => {
            console.error(
                "Socket encountered error: ",
                err,
                "Closing socket"
            );

            ws.close();
        };
    }

    check = () => {
        const { ws } = this.state;
        if (!ws || ws.readyState === WebSocket.CLOSED) this.connect(); //check if websocket instance is closed, if so call `connect` function.
    };

    update = (event) => {
        if (event.data) {

            let num_user = parseInt(event.data, 10)
            if (Number.isInteger(num_user)) {
                this.setState({
                    users: num_user,
                })

            } else
                console.log("Not integer: ", num_user)

            if (this.state.users >= this.state.capacity) // or fullError
                this.props.onMsg('warning');
        }
    }

    handleUpgrade = (event) => {
        if (this.state.capacity === 1000) {
            this.props.onMsg('Already upgrade')
            return
        }

        let apiUrl = process.env.REACT_APP_AUTH_PORT // read from config
        axios.post(apiUrl + 'tenants/plan').then(res => {
            this.setState({ capacity: res.data['capacity'], plan: 'Enterprise' })
            this.props.onMsg('success');
        }).catch(err => console.log(err));
    }

    getPlan = (cap) => {
        let dict = {
            100: "Startup",
            1000: "Enterprise"
        }

        return dict[cap]
    }

    addUser = () => {
        let apiUrl = process.env.REACT_APP_AUTH_PORT // read from config

        let payload = {
            "users": parseInt(this.state.addition, 10),
        }

        axios.post(apiUrl + 'tenants/users', payload).then(res => {
            if (res.data['full'])
                this.props.onMsg('warning');
            else
                this.props.onMsg()
        }).catch(err => {
            console.log(err)
        });
    }

    handleUserInput = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        });
    }

    render() {
        return (
            <div className="plan">
                <header>{this.getPlan(this.state.capacity)} Plan - ${this.state.capacity}/Month</header>

                <div className="plan-content">
                    <div className="progress-bar">
                        <div className="progress-bar-usage" style={{ width: (this.state.users / this.state.capacity * 100) + '%' }}></div>
                    </div>

                    <h3>Users: {this.state.users}/{this.state.capacity}</h3>
                </div>

                <footer>
                    <button className="button is-success" onClick={this.handleUpgrade}>Upgrade to Enterprise Plan</button>

                    <input
                        className="inputField"
                        type="number"
                        name='addition'
                        onChange={this.handleUserInput}
                        placeholder="0"
                    />

                    <button className="button " onClick={this.addUser}>Add user</button>

                </footer>

            </div>
        );
    }
}