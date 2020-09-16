function render() {
    var currentUrl = window.location.href;
    var keyword = getParamFromUrl("keyword", currentUrl);
    var method = getParamFromUrl("method", currentUrl);
    // if keyword is whitespaces, then donot render
    if (isNull(keyword) == true) {
        return
    }
    $.ajax({
        type: "get",
        url: "/pan/search/search",
        data: {"keyword": keyword, "method": method},
        beforeSend: function() {
            // var randomNumber = parseInt(Math.random() * 3 + 1);
            // $(".result").append('<img src="/static/gif/' + randomNumber + '.gif">');
            $(".result").append('<img src="/static/gif/loading.gif">');
        },
        success: function(response) {
            resources = response;
            renderResult(response);
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            $(".result").empty();
            $(".result").append("获取数据失败 " + XMLHttpRequest.responseText);
        }
    });
}
// check if str is "" or whitespaces
function isNull(str) {
    if ( str == "" ) {
        return true;
    }
    var regu = "^[ ]+$";
    var re = new RegExp(regu);
    return re.test(str);
}

function getParamFromUrl(name, url) {
    name = name.replace(/[\[]/,"\\\[").replace(/[\]]/,"\\\]");
    var regexS = "[\\?&]"+name+"=([^&#]*)";
    var regex = new RegExp( regexS );
    var results = regex.exec( url );
    if( results == null )
        return "";
    else
        return decodeURIComponent(results[1].replace(/\+/g, " "));
}

function renderResult(response) {
    // return if empty
    if (response === '' || response === null) {
        $(".totalCount").empty();
        $(".result").empty();
        $(".notfound-box").css("display", "block");
        return
    } else {
        var totalCount = response.length;
        $(".totalCount").empty();
        $(".totalCount").append(
            '<p>共搜索出 ' + '<span>' + totalCount + '</span>' + ' 条结果</p>'
        );
    }
    // pagination
    var pageSize = 20;
    var itemTotalCount = response.length;
    var pageNum = Math.ceil(itemTotalCount/pageSize);
    $('.pagination').pagination(itemTotalCount, {
        current_page: 0,
        ellipse_text: '...',
        num_edge_entries: 1,
        num_display_entries: 2,
        items_per_page: pageSize,
        prev_text:"上一页",
        next_text:"下一页",
        prev_show_always:true,
        next_show_always:true,
        show_firstEnd:false,
        first_text:"首页",
        end_text:"末页",
        callback: function (currentPageNum) {
            // get meta data
            var start = currentPageNum * pageSize;
            var end = start + pageSize;
            var data = response.slice(start, end);
            $(".result").empty();
            $(".result").append(
                '<div class="item-list">' +
                renderItemList(data) +
                '</div>'
            );
        }}
    );
}

function renderItemList(response) {
    var itemHtml = '';
    if (response === '' || response === null) {
        return itemHtml;
    }
    $.each(response, function (index, item) {
        var id = item["id"];
        var category = item["category"];
        var title = item["title"];
        var keyword = $("#searchinput").val();
        var re = new RegExp(keyword, "gi");
        title = title.replace(re, "<span class='highlight'>" + keyword + "</span>");
        var ctime = item["cTime"];
        var size = item["size"];
        var resource = item["resource"];
        if (category == "video") {
            var icon = "video.png";
        } else if (category == "document") {
            var icon = "document.png";
        } else if (category == "audio" || category == "picture") {
            var icon = "file.png";
        } else if (category == "seed") {
            var icon = "seed.png";
        } else if (category == "folder") {
            var icon = "folder.png";
        } else if (category == "archive") {
            var icon = "archive.png";
        } else {
            var icon = "file.png";
        }
        // handle password
        // if (password != "") {
        //     var passwordHtml = '<li class="password">' + '密码：' + password + '</li>';
        // } else {
        //     var passwordHtml = '';
        // }
        // handle size
        var realSize = CalSize(size);
        itemHtml += '<div class="item-box">' +
            '<div class="icon-box">' +
            '<img id="icon" src="/static/img/' + icon + '">' +
            '</div>' +
            '<div class="info-box">' +
            '<div class="title">' +
            '<a target="_blank" href="/pan/search/resource/' + id + '">' + title + '</a>' +
            '</div>' +
            '<ul class="attribute">' +
            '<li>' + '上传时间：' + ctime + '</li>' +
            '<li>' + '大小：' + realSize + '</li>' +
            // '<li>' + '渠道：' + resource + '</li>' +
            '</ul>' +
            '</div>' +
            '</div>';
    });
    return itemHtml;
}