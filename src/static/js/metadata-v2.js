// calculate size
$('.bdpinfo .size').append(CalSize(size));
$(document).ready(function () {
    // verify pan url
    verifyPanURL();
    // display or hidden password
    if ($('.password').text().length > 0) {
        $('.password-meta').css('display', 'block');
    }
    // distinguish pc or phone, do not display qrcode for phone
    var ispc = isPC();
    if (ispc == false || ispc == true) {
        $.cookie('codeScanned', true, {expires: 10, path: '/'});
    }
    // change btn color if cookie is valid
    if (checkIfQRCodeScanned() == true) {
        $('.btn').css('background-color', '#fa541c');
    }
});
function verifyPanURL() {
    var reqPath = "/pan/verifypanurl";
    $.ajax({
        type: "post",
        url: reqPath,
        data: {"panURL": url},
        dataType: "text",
        success: function (resp) {
            if (resp == "valid") {
                $('.verifyPanURL p').text("该链接有效，可以访问");
                $('.verifyPanURL p').css("color", "#41b035");
            } else {
                $('.verifyPanURL p').text("该链接已失效");
            }
        },
        error: function (resp) {
            $('.verifyPanURL p').text("该链接有效，可以访问");
            $('.verifyPanURL p').css("color", "#41b035");
        }
    })
}
function copy() {
    clipboard = new ClipboardJS('.copy', {
        target: function () {
            var e = document.querySelector('.password');
            return e;
        }
    });
}
// check if cookie contains codeScanned: true
function checkIfQRCodeScanned() {
    var key = "codeScanned";
    var codeScanned = $.cookie(key);
    if (codeScanned) {
        return true
    } else {
        return false
    }
}
function checkIfTicketExists() {
    var key = "ticket";
    var value = $.cookie(key);
    if (value) {
        return value
    } else {
        return false
    }
}
function setImageSrc(ticket) {
    $('#qrcodeImg').attr("src", "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket="+ticket);
}
function getTicketAndBuildWebSocket() {
    var url = "/weixin/getticket";
    $.ajax({
        url: url,
        success: function (resp) {
            var ticket = resp;
            // set ticket cookie
            var date = new Date();
            date.setTime(date.getTime() + (10 * 60 * 1000)); // expire in 10 minutes
            // console.log(date);
            $.cookie('ticket', ticket, {expires: date, path: '/'});
            // set image src
            setImageSrc(ticket);
            buildWebSocket(ticket);
        },
        error: function (resp) {
            // console.log(resp);
            // hide modal
            $('#myModal').modal('hide');
            // change btn color
            $('.btn').css('background-color', '#fa541c');
            $.cookie('codeScanned', true, {expires: 10, path: '/'});
        }
    })
}
function block() {
    if (checkIfQRCodeScanned() != true) {
        // check if ticket already exists
        var ticket = checkIfTicketExists();
        if (ticket) {
            setImageSrc(ticket);
            buildWebSocket(ticket)
        } else {
            getTicketAndBuildWebSocket();
        }
    } else {
        // jump to url if cookie is valid
        // var url = url;
        window.open(url);
    }
    return
}
function isPC() {
    var userAgentInfo = navigator.userAgent;
    var Agents = ["Android", "iPhone",
        "SymbianOS", "Windows Phone",
        "iPad", "iPod"];
    var flag = true;
    for (var v = 0; v < Agents.length; v++) {
        if (userAgentInfo.indexOf(Agents[v]) > 0) {
            flag = false;
            break;
        }
    }
    return flag;
}
function buildWebSocket(ticket) {
    var url = "ws://www.panghaozi.com/weixin/buildwebsocket";
    // var url = "ws://localhost:8080/weixin/buildwebsocket";
    var ws = new WebSocket(url);
    ws.onopen = function (event) {
        // console.log("websocket connected");
        // pop out a dialog
        $('#myModal').modal('show');
        ws.send(ticket);
    };
    ws.onclose = function (event) {
        // console.log("websocket closed");
    };
    ws.onerror = function (event) {
        console.log("websocket error");
        ws.close();
    };
    ws.onmessage = function (event) {
        if (event.data == "scanned") {
            // set cookie
            $.cookie('codeScanned', true, {expires: 10, path: '/'});
            $.cookie("ticket", "", {expires: -1, path: '/'});
            // hide modal
            $('#myModal').modal('hide');
            // change btn color
            $('.btn').css('background-color', '#fa541c');
        } else {
            // unexpected error
            console.log("unexpected error:" + event.data);
        }
        ws.close();
    }
}
