var Area = function () { }

Area.version = '1.0'

Area.triangle = function (base, height) {
  return base * height
}

Area.diamond = function (width, height) {
  return width * height / 2
}

console.log(Area.version)
console.log(Area.triangle(2,3))
console.log(Area.diamond(2, 3))

var a = new Area()
console.log(a.triangle(2,3))
console.log(a.diamond(2,3))



