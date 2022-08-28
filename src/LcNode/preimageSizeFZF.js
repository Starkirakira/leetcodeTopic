function preimageSizeFZF(k) {
    return help(k + 1) - help(k);
}
function help(k) {
    var r = 5 * k;
    var l = 0;
    while (l <= r) {
        var mid = Math.floor((l + r) / 2);
        if (zeta(mid) < k) {
            l = mid + 1;
        }
        else {
            r = mid - 1;
        }
    }
    return r + 1;
}
function zeta(x) {
    var res = 0;
    while (x != 0) {
        res += Math.floor(x / 5);
        x = Math.floor(x / 5);
    }
    return res;
}
