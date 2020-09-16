
function CalSize(size) {
    if (size == 0) return '未知';
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
