function Triangle() {
  // private members
  var _base
  var _height
  var _checkArgs = function (val) {
    return (typeof val === 'number' && val > 0)
  }

  // privileged methods
  this.setBase = function (base) {
    if(_checkArgs(base)) { _base = base }
  }
  this.getBase = function () { return _base }
  this.setHeigh = function (height) {
    if (_checkArgs(height)) { _height = height }
  }
  this.getHeight = function () { return _height }
}
// basic prototype methods which won't access private members
Triangle.prototype.getArea = function () {
  return this.getBase() * this.getHeight() / 2
}

var t = new Triangle()
t._base = 10
t._height = 5
console.log(t.getArea())  //NaN

t.setBase(10)
t.setHeigh(5)
console.log(t.getBase())  //10
console.log(t.getHeight())  //5
console.log(t.getArea())  //25


