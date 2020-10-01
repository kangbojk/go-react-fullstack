import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';

import axios from 'axios';

// send request with cookie
axios.defaults.withCredentials = true

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
