const axios = require('axios')

axios({
    method: 'POST',
    url:'http://localhost:3000/girl',
    data: {
        "id": 6,
        "name": "大波妹",
        "size": "G罩杯"
      }
}).then(res => {
    console.log(res.data);
})

axios({
    method: 'delete',
    url:'http://localhost:3000/girl/6',
}).then(res => {
    console.log(res);
})