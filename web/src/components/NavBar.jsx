import React from 'react';


import './NavBar.css'

export default class NavBar extends React.Component {

    render() {
        return (
            <div>
                <header className="top-nav">
                    <h1>
                        <i className="logo">X</i>User Management Dashboard</h1>
                    <button className="button is-border" onClick={this.props.onLogout}>Logout</button>
                </header>
            </div>
        );
    }
}