function DistanceBetweenBusStops(distance: number[], start: number, destination: number): number {
    //If the starting point is bigger than the ending point, then swap their positions
    if (start > destination) {
        const temp = start
        start = destination
        destination = temp
    }
    //Two opposite directions result
    let dis1 = 0
    let dis2 = 0

    for (let index = 0; index < distance.length; index++) {
        if (index >= start && index < destination) {
            dis1 += distance[index]
        } else {
            dis2 += distance[index]
        }
        
    }
    return Math.min(dis1,dis2)

}