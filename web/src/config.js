const dev = {
    url: {
        api_url: "http://localhost:8080/api/",
        api_auth_url: "http://localhost:8080/auth/api/",
        ws_url: "ws://localhost:8080/auth/"
    }
};

const prod = {
    url: {
        api_url: "https://peaceful-hamlet-17389.herokuapp.com/api/",
        api_auth_url: "https://peaceful-hamlet-17389.herokuapp.com/auth/api/",
        ws_url: "wss://peaceful-hamlet-17389.herokuapp.com/auth/"
    }
};

const config = process.env.REACT_APP_STAGE === 'production'
    ? prod
    : dev;

export default {
    ...config
};