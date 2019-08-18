
// const connectBtn = document.querySelector("#connect");
let isGaming = false;

function showGameResult(obj) {
    // console.log(obj)
    document.querySelector(".run").innerHTML = obj.run;
    let index = 1;
    [...document.querySelectorAll(".dice")].forEach(function (Element) {
        Element.setAttribute("src", "/static/img/game/dice/" + obj["d" + index] + ".jpg");
        index++;
    })
}

function connect() {
    let ws = new WebSocket("ws://localhost:8090/ws");

    ws.onmessage = (message) => {
        let obj = JSON.parse(message.data);
        // console.log(obj);    
        switch (obj.event) {
            case "202":
                // console.log("recived success");
                showGameResult(JSON.parse(obj.message));
                break;
            default:
                break;
        }
    }

    ws.onclose = function (evt) {      
        console.log("Connection close")
        setTimeout(function () {
            connect()
        }, 5000)
    };
}

connect();