function login(username, password) {
    var req = new XMLHttpRequest();
    var creds = { "password": password, "username": username };
    req.open("POST", "/api/v1/login", false);
    req.send(JSON.stringify(creds));
    console.log(req.responseText);
    return req.status;
}

function getHumidity() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/humidity", false);
    req.send();
    console.log(req.responseText);
    return JSON.parse(req.response).humidity;
}

function getPressure() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/pressure", false);
    req.send();
    console.log(req.responseText);
    return JSON.parse(req.response).pressure;
}

function getTemperature() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/temperature", false);
    req.send();
    console.log(req.responseText);
    return JSON.parse(req.response).temperature;
}

function generateTime(n) {
    var today = new Date();
    var time = today.getHours() + today.getMinutes() + today.getSeconds();
    var times = [];
    for (let index = n - 1; index >= 0; index--) {
        times.push(time - 5 * index);
    }
    return times;
}