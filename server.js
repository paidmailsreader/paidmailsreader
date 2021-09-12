const express = require('express')
const app = express()

app.get('/', (req, res) => {
    res.send('Key server PMR is running...')
})

app.get('/control([2-3]{0,1}).php', (req, res) => {
    if (req.params[0] === '2') {
        res.send('1t21Pz6tHt29uIt2ayHR6M8LKovW5fiddG+m0ds=')
    } else {
        res.send('5DD7D16FE42EBFF9F717B71559C05D13D9F3C89DDD28AF48902984053D')
    }
})

app.get('/update.txt', (req, res) => {
    res.sendFile(__dirname + '/update.txt')
})

app.listen(80, () => {
    console.log(`Server listening at http://localhost`)
})
