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

    fetchData()
    // show loading icon 
    // setTimeout(() => {
    //   fetchData()
    // }, 100);

  }, [login]);

  if (loading === true) {
    return <LoadingIcon type={"bubbles"} color="#00FFFF" />
  }

  return (
    <Router>
      <Switch >
        <Route exact path="/" render={() => (
          login
            ? <Redirect to={process.env.PUBLIC_URL + "/dashboard"} />
            : <Redirect to={process.env.PUBLIC_URL + "/login"} />
        )} />


        <Route exact path={process.env.PUBLIC_URL + "/"} render={() => (
          login
            ? <Redirect to={process.env.PUBLIC_URL + "/dashboard"} />
            : <Redirect to={process.env.PUBLIC_URL + "/login"} />
        )} />

        <Route path={process.env.PUBLIC_URL + "/dashboard"} render={() => (
          login
            ? <DashBoard onLogin={handleLogin} />
            : <Redirect to={process.env.PUBLIC_URL + "/login"} />
        )} />

        <Route path={process.env.PUBLIC_URL + "/login"} render={() => (
          login
            ? <Redirect to={process.env.PUBLIC_URL + "/dashboard"} />
            : <LoginForm onLogin={handleLogin} />
        )} />


        <Route path={process.env.PUBLIC_URL + "/signup"} render={() => (
          login
            ? <Redirect to={process.env.PUBLIC_URL + "/dashboard"} />
            : <SignUp onLogin={handleLogin} />
        )} />

        <Route component={Err404} />

      </Switch >

    </Router >
  );
}

export default App;