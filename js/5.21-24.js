var Animal = function () { }
var Hamster = function () { }
Hamster.prototype = new Animal()

var a = new Animal()
var h = new Hamster()

// i.constructor indicates the inherited Object.
console.log(a.constructor === Animal)  // true
console.log(a.constructor === Hamster) // false
console.log(h.constructor === Animal)  // true
console.log(h.constructor === Hamster) // false

// i.instanceof indicates if the instance is created from the Object.
console.log(a instanceof Animal)  // true
console.log(a instanceof Hamster) // false
console.log(h instanceof Animal)  // true
console.log(h instanceof Hamster) // true

// .isPrototypeOf(i) confirms that the instance refers which prototype.
console.log(Hamster.prototype.isPrototypeOf(a)) // false
console.log(Animal.prototype.isPrototypeOf(a))  // true
console.log(Hamster.prototype.isPrototypeOf(h)) // true
console.log(Animal.prototype.isPrototypeOf(h))  // true

// in checks if an instance has the specific member
var obj = { hoge: 'hoge', foo: function () { } }

console.log('hoge' in obj)  // true
console.log('foo' in obj)   // true
console.log('pyo' in obj)   // false
