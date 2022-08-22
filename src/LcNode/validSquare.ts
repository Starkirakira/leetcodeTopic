export function ValidSquare(p1: number[], p2: number[], p3: number[], p4: number[]): boolean {
    //determine side lengths are equal
    if (p1.toString() == p2.toString() && p2.toString() == p3.toString() && p3.toString() == p4.toString()) {
        return false
    }
    let side_arr: number[][] = [p1,p2,p3,p4]
    side_arr.sort(
        (a, b) => {
            return a[0] - b[0] 
        }
    )
    let side_arr_1 = [side_arr[0], side_arr[1]]
    let side_arr_2 = [side_arr[2], side_arr[3]]

    if (side_arr_1[0][0] == side_arr_1[1][0]) {
        side_arr_1.sort((a, b) => a[1] - b[1])
    }
    if (side_arr_2[0][0] == side_arr_2[1][0]) {
        side_arr_2.sort((a, b) => a[1] - b[1])
    }
    console.log(side_arr_1, side_arr_2)
    //p1 ==> side_arr[0]; p2 ==> side_arr[1]; p3 ==> side_arr[2]; p4 ==> side_arr[3];
    let a1 = side_arr_1[0]
    let a2 = side_arr_1[1]
    let a3 = side_arr_2[0]
    let a4 = side_arr_2[1]
    //The length of a1-a2
    let a1a2_x = Math.pow(a1[0]-a2[0],2)
    let a1a2_y = Math.pow(a1[1]-a2[1],2)
    //The length of a3-a4
    let a3a4_x = Math.pow(a3[0]-a4[0],2)
    let a3a4_y = Math.pow(a3[1]-a4[1],2)
    //The length of a1-a3
    let a1a3_x = Math.pow(a1[0]-a3[0],2)
    let a1a3_y = Math.pow(a1[1]-a3[1],2)
    //The length of a2-a4
    let a2a4_x = Math.pow(a2[0]-a4[0],2)
    let a2a4_y = Math.pow(a2[1]-a4[1],2)
    //The length of a1-a4 && a2-a3
    let a1a4_x = Math.pow(a1[0]-a4[0],2)
    let a1a4_y = Math.pow(a1[1]-a4[1],2)

    let a2a3_x = Math.pow(a2[0]-a3[0],2)
    let a2a3_y = Math.pow(a2[1]-a3[1],2)
    //special case
    
    if (a1a2_x+a1a2_y == a3a4_x + a3a4_y && a1a3_x + a1a3_y == a2a4_x + a2a4_y && a1a4_x + a1a4_y == a2a3_x + a2a3_y && a1a2_x + a1a2_y == a1a3_x + a1a3_y) {
        return true
    }
    return false
}

console.log(ValidSquare([0,0],[0,0],[0,0],[0,0]))
