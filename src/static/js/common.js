$('#panform').submit(function (event) {
    event.preventDefault();
    var keyword = $('#searchinput').val();
    window.location.href = "/pan/search/" + encodeURIComponent(keyword);
    return false;
});

function CalSize(size) {
    size = size / 1024;
    if (size < 1024) return size.toFixed(1) + ' KB';
    size = size / 1024;
    if (size < 1024) return size.toFixed(1) + ' MB';
    size = size / 1024;
    if (size < 1024) return size.toFixed(1) + ' G';
    size = size / 1024;
    if (size < 1024) return size.toFixed(1) + ' T';
};

$('#magnetform').submit(function (event) {
    event.preventDefault();
    var keyword = $('#searchinput').val();
    window.location.href = "/magnet/search/" + encodeURIComponent(keyword);
    return false;
});

// $('.navimenu a').click(function (event) {
//     // event.preventDefault();
//     // $('.navimenu a').each(function (index, item) {
//     //     item.style.color = "#0078ff";
//     // });
//     $(this).css("color", "#D23141").siblings().css("color", "#0078ff");
// });