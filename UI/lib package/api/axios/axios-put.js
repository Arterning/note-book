const axios = require('axios')

axios({
    method: 'PUT',
    url:'http://localhost:3000/girl/4',
    data: {
        "id": 4,
        "name": "大波妹 花花",
        "size": "E 罩杯",
        "age": 23,
        "height": "168cm",
        "address": "chang sha"
      }
}).then(res => {
    console.log(res.data);
})