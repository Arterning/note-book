//import {axios} from 'axios'
const axios = require('axios')


axios({
    method: 'POST',
    url:'http://localhost:3000/girl',
    data: {
        "id": 5,
        "name": "大波妹 静静",
        "size": "F 罩杯",
        "age": 23,
        "height": "168cm"
      }
}).then(res => {
    console.log(res.data);
})