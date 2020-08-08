const mongoose = require('mongoose')


const mongoDB = 'mongodb://localhost:27017/blockflow'

mongoose.connect(mongoDB, {useNewUrlParser: true});

const db = mongoose.connection;

db.on('error', console.error.bind(console, "MongoDB connection error:"))

mongoose.Promise = global.Promise

module.exports = mongoose;

