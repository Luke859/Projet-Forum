console.log("Script signIn loaded")

let pseudo = document.getElementById('pseudo').value
let password = document.getElementById('password').value

const userRoutes = require('/routes/user');

signIn.use('/api/stuff', stuffRoutes);
signIn.use('/api/auth', userRoutes);

console.log(pseudo);
console.log(password);