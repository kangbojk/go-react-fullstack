import axios from 'axios';

export const isLogin = () => {
    let apiUrl = process.env.REACT_APP_AUTH_PORT // read from config    
    return axios.get(apiUrl + 'account');
}