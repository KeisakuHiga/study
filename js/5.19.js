var Animal = function (name, age) {
  this.name = name
  this.age = age
}
Animal.prototype = {
  walk: function () {
    console.log('toko...toko...');
  }
}

var Dog = function () {
  // 継承してる
  Animal.call(this, 'John', 3)
}
// DogオブジェクトのプロトタイプとしてAnimalオブエクトのインスタンスをセット
Dog.prototype = new Animal()
Dog.prototype.bark = function () {
  console.log('one! one!')
}

var d = new Dog()
// Animalオブジェクトのprototypeにセットしたwalk()がDogインスタンスからでも呼び出せる
// プロパティ・メソッドの検索順番
// Dogインスタンス自身のメンバー
//  -> Dogオブジェクトのprototype
//  -> Animalインスタンスのメンバー
//  -> Animalオブジェクトのprototype
//  -> （最上位）Object.prototype
d.walk()
d.bark()
console.log(d.name)


