
// const connectBtn = document.querySelector("#connect");
let isGaming = false;
let timmer;
let ws;
let betBtnList = [];

const StatusMap = {
    1: "New Run",
    2: "Show Down",
    3: "Settlement"

}

const COMMAND_CONNECTED = "200",
    COMMAND_NEW_RUN = "201",
    COMMAND_SHOWDOWN = "202",
    COMMAND_RESULT = "203",
    COMMAND_BET = "204";

const btn_dice_big = document.getElementById("dice-big"),
    btn_dice_small = document.getElementById("dice-small"),
    btn_dice_odd = document.getElementById("dice-odd"),
    btn_dice_even = document.getElementById("dice-even");

function showStatus(status) {
    document.querySelector("#status").innerHTML = status;
}

function bgChange() {
    let count = 1;
    let oldClass = "bg1";
    setInterval(function () {     
        count = count % 2 + 1;
        document.getElementById("container").classList.replace(oldClass, "bg" + count);
        oldClass = "bg" + count;
    }, 10 * 1000)
}

function showGameResult(obj) {
    //console.log(obj)
    detail = obj.game_detail
    document.querySelector("#run").innerHTML = obj.run;
    document.querySelector("#inn").innerHTML = obj.inn;
    let index = 1;
    [...document.querySelectorAll(".dice")].forEach(function (Element) {
        Element.setAttribute("src", "/static/img/game/dice/" + detail["d" + index] + ".jpg");
        index++;
    })
}

function startNewRun(cd) {
    // let cd = 10;
    //console.log(cd);
    timmer = function () {
        if (cd >= 0) {
            document.querySelector("#countdown").innerHTML = cd--;
            setTimeout(timmer, 1000);
        }
    }
    timmer();
}

function connect() {
    let counter = 5;
    //console.log("memberID=" + memberID)
    ws = new WebSocket("ws://localhost:8090/ws?memberID=" + memberID);

    ws.onmessage = (message) => {
        //console.table(message.data)
        let obj = JSON.parse(message.data);
        //console.log(new Date().toLocaleString() + " " + obj.event)
        switch (obj.event) {
            case COMMAND_CONNECTED:
                console.log(obj)
                register();
                getTableStatus(obj);
                break;
            case COMMAND_NEW_RUN:
                showStatus("New Run");
                startNewRun(obj.message);
                // console.log(new Date().toLocaleString() + " New Run")
                break;
            case COMMAND_SHOWDOWN:
                showStatus("Show Down");
                // console.log(new Date().toLocaleString() + " Show Down")
                showGameResult(JSON.parse(obj.message));
                break;
            case COMMAND_RESULT:
                showStatus("Settlement");
                // console.log(new Date().toLocaleString() + " Settlement")
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

function register() {
    console.log("send login")
    let data = { event: '200', message: '{"name":"edlo", "email":"test@example.com", "password":"8888"}' }
    ws.send(JSON.stringify(data))
}


function bet(game, betArea) {
    console.log("bet")
    let data = { event: '301', message: '{"game":' + game + ', "bet-area":"' + betArea + '", "amount":100}' }
    ws.send(JSON.stringify(data))
}

function init() {

    let btnElementList = document.getElementsByClassName("bet-btn")

    // window.a = a;
    Array.from(btnElementList).map(element => {
        element.onmouseenter = function (e) {
            e.path[0].classList.add("btn-toggle");
        }

        element.onmouseleave = function (e) {
            e.path[0].classList.remove("btn-toggle");
        }

        element.onclick = function (e) {
            // window.e = e.path[0]
            let data = e.path[0].dataset
            bet(data.game, data.area);
        }
    })

    // Betting
    // let betBtn = {};

    // betBtn.child = document.getElementById("dice-big");

    // window.b = betBtn;
}

function getTableStatus(data) {
    console.log("set table status")

    let d = JSON.parse(data.message)
    console.log(d)

    document.querySelector("#run").innerHTML = d.Run;
    document.querySelector("#inn").innerHTML = d.Inn;
    showStatus(StatusMap[d.Status])
    startNewRun(d.Countdown - 1)
}

init();
connect();
bgChange();