window.onload = function () {
    if (!'WebSocket' in window) {
        console.log("%cYour browser does not support WebSocket", " text-shadow: 0 1px 0 #ccc,0 2px 0 #c9c9c9,0 3px 0 #bbb,0 4px 0 #b9b9b9,0 5px 0 #aaa,0 6px 1px rgba(0,0,0,.1),0 0 5px rgba(0,0,0,.1),0 1px 3px rgba(0,0,0,.3),0 3px 5px rgba(0,0,0,.2),0 5px 10px rgba(0,0,0,.25),0 10px 10px rgba(0,0,0,.2),0 20px 20px rgba(0,0,0,.15);font-size:5em")
        return false
    } else {
        new LiteMain();
    }
}

function LiteMain() {
    let ws = new WebSocket("ws://127.0.0.1:8866/ws/onlineServer?weburl="+ getwebURL());
    new CreateView()
    const view = document.getElementById("onlineview");
    ws.onmessage = function (evt) {
        view.innerHTML = evt.data + "&nbsp;人正在瀏覽";
    }
    ws.onerror = function () {
        console.info("Connection WebSocket Error!")
    }
    ws.onclose = function () {
        console.info("WebSocket connection is closed!")
    }
    window.onbeforeunload = function () {
        ws.close();
    }
}

function CreateView() {
    var sty = document.createElement("style");
    var csscontent = document.createTextNode("#onlineview{bottom: 0px;text-align:center;color:#3A3A3A;font-size:13px}")
    sty.appendChild(csscontent)
    var d = document.createElement("div");
    d.id = "onlineview";
    d.style.bottom = "0px";
    document.body.appendChild(d);
}

function getwebURL() {
    let host = window.location.href.trim();
    -1 !== host.indexOf("#") && (host = host.substr(0, host.indexOf("#")));
    -1 !== host.indexOf("?") && (host = host.substr(0, host.indexOf("?")));
    "/" === host.substr(host.length - 1, 1) && (host = host.substr(0, host.length - 1));
    host_arr = host.split("/");
    web_url = host_arr[2];
    void 0 !== host_arr[3] && (web_url = host_arr[2] + "/" + host_arr[3]);
    web_url = String(web_url);
    return web_url;
}