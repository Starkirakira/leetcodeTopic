function ArrayRankTransform(arr) {
    arr.sort(function (a, b) { return a - b; });
    console.log(arr);
    return 1;
}
ArrayRankTransform([10, 1, 3, 2]);
