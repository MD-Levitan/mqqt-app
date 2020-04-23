function getHumidity() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/humidity", false);
    req.send();
    return JSON.parse(req.response).humidity;
}

function getPressure() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/pressure", false);
    req.send();
    return JSON.parse(req.response).pressure;
}

function getDevice() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/device", false);
    req.send();
    return JSON.parse(req.response);
}

function getTemperature() {
    var req = new XMLHttpRequest();
    req.open("GET", "/api/v1/temperature", false);
    req.send();
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
            ctx.data.labels.push(ctx.data.labels[ctx.data.labels.length - 1] + 5);
            ctx.data.datasets.forEach((dataset) => {
                dataset.data.push(data[data.length - 1]);
            });
            ctx.update();
        }
    }, timeout, ctx, upd_func);
}

var updateTime = 5000;

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
    updateChart(chart, getTemperature, updateTime);
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
    updateChart(chart, getPressure, updateTime);
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
    updateChart(chart, getHumidity, updateTime);
    return chart;
}

function addDevice(ctx) {

    var dev = getDevice();

    if (dev["udev"] == null && dev["adev"] == null) { return; }

    if (dev["udev"] == null) {
        ctx.children["status"].innerHTML = dev["adev"].status;
        ctx.children["info"].innerHTML = dev["adev"].info;
    } else {
        ctx.children["status"].innerHTML = dev["udev"].status;
        ctx.children["info"].innerHTML = dev["udev"].info;
    }

    var temp = getTemperature();
    var pres = getPressure();
    var hum = getHumidity();

    ctx.children["temperature"].innerHTML = temp[temp.length - 1];
    ctx.children["pressure"].innerHTML = pres[pres.length - 1];
    ctx.children["humidity"].innerHTML = hum[hum.length - 1] + "%";

}

function addAdminDevice(ctx) {

    var dev = getDevice();

    if (dev["udev"] == null && dev["adev"] == null) { return; }

    if (dev["adev"] == null) {
        ctx.innerHTML = "No permissions"
    } else {
        ctx.children["status"].innerHTML = dev["adev"].status;
        ctx.children["info"].innerHTML = dev["adev"].info;
        ctx.children["battery"].innerHTML = dev["adev"].battery + "%";
        ctx.children["time"].innerHTML = dev["adev"].time;
        ctx.children["model"].innerHTML = dev["adev"].model;
    }

    var temp = getTemperature();
    var pres = getPressure();
    var hum = getHumidity();

    ctx.children["temperature"].innerHTML = temp[temp.length - 1];
    ctx.children["pressure"].innerHTML = pres[pres.length - 1];
    ctx.children["humidity"].innerHTML = hum[hum.length - 1] + "%";


}

function uploadFirmware(filename) {
    alert("Upload funconality doesn't work yet!")
}