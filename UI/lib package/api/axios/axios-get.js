//import {axios} from 'axios'
const axios = require('axios')

axios({
    method: 'GET',
    url:'http://localhost:3000/girl'
}).then(res => {
    console.log(res.data);
})

