function SolveEquation(equation: string): string {
   let x = 0, num = 0, n = equation.length 
   for (let i = 0,  op = 1; i < n;) {
        if(equation[i] == '+') {
            //Set operator property and let index plus 1
            op = 1; i++
        }else if (equation[i] == '-') {
            op = -1; i++
        } else if (equation[i] == '=') {
            //Let operator be reversed to plus, and let index plus 1
            x *= -1; num *= -1; op = 1; i++
        } else {
            //Use j and a while loop to mark the operation part's index
            let j = i
            //If the next string is not '+' or '-' or '=', let j plus 1, so the operation part is equation.substring(i,j), if it is the coefficient part, operation part is equation.substring(i,j-1)
            while(j < n && equation[j] != '+' && equation[j] != '-' && equation[j] != '=') j++
            if (equation[j-1] == 'x') x += (i < j -1 ? Number(equation.substring(i,j-1)) : 1) * op 
            else num += Number(equation.substring(i,j)) * op
            i = j
        }
   }

   if(x == 0) return (num == 0)?'Infinite solutions':'No solution' 

   return 'x=' + String(num / -x)

   

}
console.log(SolveEquation('-x-2x-3=4-2x'))