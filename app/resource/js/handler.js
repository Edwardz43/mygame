
// const connectBtn = document.querySelector("#connect");
let isGaming = false;
let ws;
let betBtnList = [];
let counter = 10;

const StatusMap = {
    1: "New Run",
    2: "Show Down",
    3: "Settlement"

}

const GameID = {
    1 : "dice",
    2 : "rullote",
    3 : "dragontiger"
}

const COMMAND_CONNECTED = "200",
    COMMAND_NEW_RUN = "201",
    COMMAND_SHOWDOWN = "202",
    COMMAND_RESULT = "203",
    COMMAND_BET = "204",
    COMMAND_COUNTDOWN = "205";

const btn_dice_big = document.getElementById("dice-big"),
    btn_dice_small = document.getElementById("dice-small"),
    btn_dice_odd = document.getElementById("dice-odd"),
    btn_dice_even = document.getElementById("dice-even");


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
    console.log(obj)
    detail = obj.game_detail
  
    
    switch(obj.game_type){
        case 1:
            document.querySelector("#dice-run").innerHTML = obj.run;
            document.querySelector("#dice-inn").innerHTML = obj.inn;
            [...document.querySelectorAll(".dice")].forEach(function (Element, Index) {
                Element.setAttribute("src", "/static/img/game/dice/" + detail["d" + (Index + 1)] + ".jpg");
                Index++;
            })
            break;
        case 3:
            document.querySelector("#dragontiger-run").innerHTML = obj.run;
            document.querySelector("#dragontiger-inn").innerHTML = obj.inn;           
            document.getElementById("d_card").setAttribute("src", "/static/img/game/dragontiger/" + detail["d_card"] + ".jpg");
            document.getElementById("t_card").setAttribute("src", "/static/img/game/dragontiger/" + detail["t_card"] + ".jpg");
            break;
    }
    
}

function connect() {
    
    //console.log("memberID=" + memberID)
    ws = new WebSocket("ws://localhost:8090/ws?memberID=" + memberID);

    ws.onmessage = (message) => {        
        // handle sticky packets
        // console.log(message.data)
        
        try{
            message.data.split('\n').forEach(element => {
                operation(JSON.parse(element))    
            })        
        }
        catch(e){
            console.log(e)
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

function operation(cmd) {
    switch (cmd.event) {
        case COMMAND_CONNECTED:
            // console.log(cmd)
            register();
            setTableStatus(cmd);
            break;
        case COMMAND_NEW_RUN:            
            showStatus("New Run", GameID[JSON.parse(cmd.message).game_type]);            
            countdown(cmd);
            // console.log(new Date().toLocaleString() + " New Run")
            break;
        case COMMAND_SHOWDOWN:
            showStatus("Show Down", GameID[JSON.parse(cmd.message).game_type]);            
            // console.log(new Date().toLocaleString() + " Show Down")
            showGameResult(JSON.parse(cmd.message));
            break;
        case COMMAND_RESULT:
            showStatus("Settlement", GameID[JSON.parse(cmd.message).game_type]);
            // console.log(new Date().toLocaleString() + " Settlement")
            break;
        case COMMAND_COUNTDOWN:
            countdown(cmd)
        default:
            break;
    }
}

function register() {
    console.log("send login")
    let data = { event: '200', message: '{"name":"edlo", "email":"test@example.com", "password":"8888"}' }
    ws.send(JSON.stringify(data))
}


function bet(game, run, inn, betArea) {
    console.log("bet")
    
    let msg = {
        game: parseInt(game) ,
        run: parseInt(run) ,
        inn: parseInt(inn),
        bet_area :betArea,
        amount:100
    } 

    let data = { 
        event: '301',
        message: JSON.stringify(msg)
    }

    let s = JSON.stringify(data);

    ws.send(s)
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
            let data = e.path[0].dataset;
            let game_type = GameID[data.game];            
            run = document.getElementById(game_type + "-run").innerText;
            inn = document.getElementById(game_type + "-inn").innerText;
            bet(data.game,run, inn, data.area);
        }
    })
}

function setTableStatus(data) {
    // let d = JSON.parse(data.message);
    // setGameInfo(d.run, d.inn, d.status, d.countdown)

    // let result = d.result;    
    // document.getElementById("d1").setAttribute("src", "/static/img/game/dice/" + result.d1 + ".jpg");
    // document.getElementById("d2").setAttribute("src", "/static/img/game/dice/" + result.d2 + ".jpg");
    // document.getElementById("d3").setAttribute("src", "/static/img/game/dice/" + result.d3 + ".jpg");
}

function countdown(data) {
    // console.log(data)
    let d = JSON.parse(data.message);      
    setGameInfo(d.game_type, d.run, d.inn, 1, d.countdown)
}

function setGameInfo(game_type, run, inn, status, cd) {
    document.querySelector("#"+ GameID[game_type] +"-run").innerHTML = run;
    document.querySelector("#"+ GameID[game_type] +"-inn").innerHTML = inn;
    document.querySelector("#"+ GameID[game_type] +"-countdown").innerHTML = cd;
    showStatus(StatusMap[status], GameID[game_type]);   
}

function showStatus(status, game_type) {
    document.querySelector("#" + game_type + "-status").innerHTML = status;
}


init();
connect();
bgChange();