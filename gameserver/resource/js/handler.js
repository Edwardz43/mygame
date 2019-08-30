
// const connectBtn = document.querySelector("#connect");
let isGaming = false;

const COMMAND_RESULT = "202",
    COMMAND_NEW_RUN = "201";

function bgChange() {
    let count = 1;
    let oldClass = "bg1";
    setInterval(function () {
        // console.log("bgchange()");
        count = count % 5 + 1;
        document.getElementById("container").classList.replace(oldClass, "bg" + count);
        oldClass = "bg" + count;
    }, 10 * 1000)
}

function showGameResult(obj) {
    console.log(obj)
    detail = obj.game_detail
    document.querySelector("#run").innerHTML = obj.run;
    document.querySelector("#inn").innerHTML = obj.inn;
    let index = 1;
    [...document.querySelectorAll(".dice")].forEach(function (Element) {
        Element.setAttribute("src", "/static/img/game/dice/" + detail["d" + index] + ".jpg");
        index++;
    })
}

function startNewRun(obj) {
    let cd = obj.message;
    console.log(cd);

    let countdown = setInterval(function () {
        if (cd >= 0) {
            document.querySelector("#countdown").innerHTML = cd--;
        } else {
            clearInterval(countdown);
        }
    }, 1000)

}

function connect() {
    let ws = new WebSocket("ws://localhost:8090/ws");
    let counter = 5;

    ws.onmessage = (message) => {
        let obj = JSON.parse(message.data);
        switch (obj.event) {
            case COMMAND_RESULT:
                showGameResult(JSON.parse(obj.message));
                break;
            case COMMAND_NEW_RUN:
                startNewRun(obj);
                break;
            default:
                break;
        }
    }

    ws.onclose = function (evt) {
        if (counter >= 0) {
            console.log("Connection close")
            setTimeout(function () {
                counter--;
                connect();
            }, 5000)
        }

    };
}

connect();
bgChange();