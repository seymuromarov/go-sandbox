var numbers = 10;
var num1 = 0;
var num2 = 1;
var next;
console.log("Fibonacci Series in js:");
for (var i = 1; i <= numbers; i++) {
  console.log(num1);
  next = num1 + num2;
  num1 = num2;
  num2 = next;
}
