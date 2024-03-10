const numbers: number = 10;
let num1: number = 0;
let num2: number = 1;
let next: number;

console.log("Fibonacci Series in ts:");

for (let i = 1; i <= numbers; i++) {
  console.log(num1);
  next = num1 + num2;
  num1 = num2;
  num2 = next;
}
