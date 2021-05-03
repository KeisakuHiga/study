// new args in ES2015
function getTriangle(base = 2, height = 4) {
  return base * height / 2
}
console.log("getTriangle", getTriangle(4))

function getTriangle2(base = 2, height) {
  return base * height / 2
}
console.log("getTriangle2", getTriangle2(4))

function multi(a, b = a) {
  return a * b
}

console.log("multi", multi(10, 5))
console.log("multi", multi(10))
console.log("multi", multi(2, 0))
console.log("multi", multi(2, null))
console.log("multi", multi(2, false))
console.log("multi", multi(2, undefined))

const nums = [-1, 22, 83, -14, 5]

function sum(start, ...nums) {
  console.log("nums: ", nums)
  console.log("type of nums: ", typeof nums)
  for (let num of nums) {
    if (typeof num !== 'number') {
      throw new Error('wrong params type')
    }
    start += num
  }
  return start
}

try {
  console.log(sum(0, ...nums))
} catch (err) {
  console.error("Error: ", err.message)
}

console.log(Math.max.apply(null, nums))
console.log(...[-1, 22, 83, -14, 5])
