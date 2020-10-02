import React from 'react';
import { useState, useEffect } from 'react';


import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";

import './App.css';
import LoginForm from './components/LoginForm'
import DashBoard from './components/DashBoard'
import Err404 from './components/Err404'
import SignUp from './components/SignUp'
import { isLogin } from './api/auth'

import ReactLoading from 'react-loading';
import Cookies from 'js-cookie'



const LoadingIcon = ({ type, color }) => (
  <ReactLoading className="preloader" type={type} color={color} height={'20%'} width={'20%'} />
);

function App() {
  const [login, setLogin] = useState(false);
  const [loading, setLoading] = useState(true);

  const handleLogin = (status) => {
    console.log("handleLogin: ", status)
    setLogin(status)
  };

  useEffect(() => {
    setLoading(true)
    async function fetchData() {
      try {
        const res = await isLogin();
        if (res.status === 200) {
          setLogin(true)
          Cookies.set('user', JSON.stringify(res.data))
        }
        else {
          setLogin(false)
          Cookies.remove('user')
        }
        setLoading(false)
      } catch (err) {
        console.log(err)
        setLogin(false)
        setLoading(false)
        Cookies.remove('user')
      }
    }

    // show loading icon 
    setTimeout(() => {
      fetchData()
    }, 500);
  }, [login]);

  if (loading === true) {
    return <LoadingIcon type={"bubbles"} color="#00FFFF" />
  }

  return (
    <Router>
      <Switch >
        <Route exact path="/" render={() => (
          login
            ? <Redirect to="/dashboard" />
            : <Redirect to="/login" />
        )} />

        <Route exact path="/dashboard" render={() => (
          login
            ? <DashBoard onLogin={handleLogin} />
            : <Redirect to="/login" />
        )} />

        <Route exact path="/login" render={() => (
          login
            ? <Redirect to="/dashboard" />
            : <LoginForm onLogin={handleLogin} />
        )} />


        <Route exact path="/signup" render={() => (
          login
            ? <Redirect to="/dashboard" />
            : <SignUp onLogin={handleLogin} />
        )} />

        <Route>
          <Err404 />
        </Route>
      </Switch >

    </Router>
  );
}

export default App;