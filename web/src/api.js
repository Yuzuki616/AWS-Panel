import axios from 'axios';

export default axios.create({
    baseURL: ``,
    withCredentials: true,
    validateStatus: function (status) {
        return status < 500;
    }
});