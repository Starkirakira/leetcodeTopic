function DistanceBetweenBusStops(distance, start, destination) {
    if (start > destination) {
        var temp = start;
        start = destination;
        destination = temp;
    }
    var dis1 = 0;
    var dis2 = 0;
    for (var index = 0; index < distance.length; index++) {
        if (index >= start && index < destination) {
            dis1 += distance[index];
        }
        else {
            dis2 += distance[index];
        }
    }
    return Math.min(dis1, dis2);
}
