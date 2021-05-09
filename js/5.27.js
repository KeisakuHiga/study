function Triangle() {
  // private members
  var _base
  var _height

  // define 'base' private member by Object.defineProperties()
  Object.defineProperties(
    this, {
      base: {
        get: function () {
          return _base
        },
        set: function (base) {
          if (typeof base === 'number' && base > 0) {
            _base = base
          }
        }
      },
      height: {
        get: function () {
          return _height
        },
        set: function (height) {
          if (typeof height === 'number' && height > 0) {
            _height = height
          }
        }
      }
    }
  )
}
// basic prototype methods which won't access private members
Triangle.prototype.getArea = function () {
  return this.base * this.height / 2
}

var t = new Triangle()
t.base = 10
t.height = 5
console.log(t.base)  //10
console.log(t.height)  //5
console.log(t.getArea())  //25


