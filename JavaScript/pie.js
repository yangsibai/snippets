// Draw pie chart

var pie = function (ctx, w, h, total, list) {
    var radius = h / 2 - 5;
    var centerX = w / 2;
    var centerY = h / 2;
    var lastEnd = 0;
    list.forEach(function (item) {
        ctx.fillStyle = item.color;
        ctx.beginPath();
        ctx.moveTo(centerX, centerY);
        ctx.arc(centerX, centerY, radius, lastEnd, lastEnd + (Math.PI * 2 * (item.count / total)), false);
        ctx.lineTo(centerX, centerY);
        ctx.fill();
        lastEnd += Math.PI * 2 * (item.count / total);
    });
};
var colors = {
    "0": "#ECD078",
    "1": "#D95B43",
    "2": "#00AA00"
};
