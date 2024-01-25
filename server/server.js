const express = require('express');
const app = express();
const process = require(`process`);

const PORT = process.argv[2];
const SERVER_NUMBER = process.argv[3];

app.get("/",(req,res)=>{
    console.log(`Recieved Request!`);
    res.send(`Hello From Server ${SERVER_NUMBER}`);
})

app.listen(PORT, ()=>{
    console.log(`[SERVER ${SERVER_NUMBER}]`)
    console.log(`Listening on PORT: ${PORT}`);
})