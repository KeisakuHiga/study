function closure(init) {
  var counter = init
  return function () {
    // returns after-incremented number
    return ++counter
  }
}
var myClosure = closure(1)
console.log(myClosure())
console.log(myClosure())
console.log(myClosure())