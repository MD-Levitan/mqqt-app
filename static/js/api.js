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
    if (n === null) {
        n = 0;
    } else {
        n = n.length;
    }
    console.log(n);
    var today = new Date();
    var time = today.getHours() + today.getMinutes() + today.getSeconds();
    var times = [];
    for (let index = n - 1; index >= 0; index--) {
        times.push(time - 5 * index);
    }
    console.log(times);
    return times;
}

function updateChart(ctx, upd_func, timeout) {
    setInterval(function (ctx, upd_func) {
        data = upd_func();
        if (data !== null) {
            ctx.labels.push(ctx.labels[ctx.labels.length - 1] + 5);
            ctx.data.datasets[0].push(data);
            ctx.update();
        }
    }, timeout, ctx, upd_func);
}

var updateTime = 5000000;

function addTemperatureChart(ctx) {
    var data = getTemperature();
    var chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: generateTime(data),
            datasets: [
                {
                    label: "Temperature",
                    backgroundColor: 'rgba(12, 199, 132, 0.2)',
                    //borderColor: 'rgba(12, 199, 132, 0.2)',
                    data: data,

                },
            ]
        },
        options: {
            responsive: true,
            title: {
                display: true,
                text: 'Smart House - Temperature'
            },
            tooltips: {
                mode: 'index',
                intersect: false,
            },
            hover: {
                mode: 'nearest',
                intersect: true
            },
            scales: {
                xAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Time'
                    }
                }],
                yAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Value'
                    }
                }]
            }
        }
    });
    //updateChart(chart, getTemperature, updateTime);
    return chart;
}

function addPressureChart(ctx) {
    var data = getPressure();
    var chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: generateTime(data),
            datasets: [
                {
                    label: "Pressure",
                    backgroundColor: 'rgba(122, 99, 67, 0.2)',
                    //borderColor: 'rgba(12, 199, 132, 0.2)',
                    data: data,

                },
            ]
        },
        options: {
            responsive: true,
            title: {
                display: true,
                text: 'Smart House - Pressure'
            },
            tooltips: {
                mode: 'index',
                intersect: false,
            },
            hover: {
                mode: 'nearest',
                intersect: true
            },
            scales: {
                xAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Time'
                    }
                }],
                yAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Value'
                    }
                }]
            }
        }
    });
    //updateChart(chart, getPressure, updateTime);
    return chart;
}

function addHumidityChart(ctx) {
    var data = getHumidity();
    var chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: generateTime(data),
            datasets: [
                {
                    label: "Humidity",
                    backgroundColor: 'rgba(23, 12, 207, 0.2)',
                    //borderColor: 'rgba(12, 199, 132, 0.2)',
                    data: data,

                },
            ]
        },
        options: {
            responsive: true,
            title: {
                display: true,
                text: 'Smart House - Humidity'
            },
            tooltips: {
                mode: 'index',
                intersect: false,
            },
            hover: {
                mode: 'nearest',
                intersect: true
            },
            scales: {
                xAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Time'
                    }
                }],
                yAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Value'
                    }
                }]
            }
        }
    });
    //updateChart(chart, getHumidity, updateTime);
    return chart;
}