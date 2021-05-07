var Animal = function (name, age) {
  this.name = name
  this.age = age
}
Animal.prototype = {
  walk: function () {
    console.log('toko...toko...');
  },
	bark: function() {
    console.log('one!one!')
  }
}

var SuperAnimal = function () { }
SuperAnimal.prototype = {
  walk: function () {
    console.log('wowowowowowowowo')
  }
}

var Dog = function () {
  Animal.call(this, 'John', 3)
}
Dog.prototype = new Animal()
var d1 = new Dog()
d1.walk() // toko...toko...

Dog.prototype = new SuperAnimal()
var d2 = new Dog()
// Dogオブジェクトのprototypeにセットされている継承元が上書きされているから
d2.walk() // wowowowowowowowo

// Dogオブジェクトのprototypeにセットされている継承元が上書きされているように見えるけど...
// インスタンスが生成された時点で参照するprototypeは固定される！その後のprototypeの上書きに影響されない！
d1.walk() // toko...toko...
d1.bark() // one!one!
console.log(d2.name)
// d2はプロトタイプを経由してSuperAnimalオブジェクトを暗黙参照しているが、
// SuperAnimalオブジェクトにbarkメソッドが定義されていないからバグる
d2.bark() // d2.bark is not a function