import axios from 'axios';
import config from "../config"

export const isLogin = () => {
    let apiUrl = config.url.api_auth_url // read from config    
    return axios.get(apiUrl + 'account');
}