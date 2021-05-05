var Member = function (firstName, lastName) {
  console.log(this instanceof Member)
  if (!(this instanceof Member)) {
    return new Member(firstName, lastName)
  }
  this.firstName = firstName
  this.lastName = lastName
  this.getName = function () {
    return this.firstName + ' ' + this.lastName
  }
}
var member = new Member('first', 'last')
console.log(member.firstName)
console.log(member.lastName)
console.log(member.getName())


member.hello = function () {
  return 'hello'
}

console.log(member.hello())

var data = 'Global data'
var obj1 = { data: 'obj1 data' }
var obj2 = { data: 'obj2 data' }

function hoge() {
  // console.log(this)
  console.log(this.data)
}

hoge.call(data)
hoge.call(obj1)
hoge.call(obj2)

function callArgs() {
  // var args = Array.prototype.slice.call(arguments)
  let args = Array.from(arguments)
  console.log(args.join('/'))
}
callArgs('hahaha', 'hehehe')
callArgs('A', 'B', 'C')

var m = Member('Mike', 'Brown')
console.log(m)
// console.log(firstName)
console.log(m.firstName)

console.log('x'.repeat(25) + ' prototype ' + 'x'.repeat(25))

Member.prototype.changeName = function (newFirstName,newLastName) {
  this.firstName = newFirstName
  this.lastName = newLastName
}
var mike = new Member('mike', 'scot')
Member.prototype.hello = function () {
  console.log(`hello ${this.firstName} !`)
}
console.log(mike.hello())
console.log(mike.getName())

mike.changeName('foo', 'bar')
console.log(mike.getName())

console.log(member.getName())
member.changeName('foo', 'bar')
console.log(member.getName())

var Human = function (name, age) {
  this.name = name
  this.age = age
}
Human.prototype = {
  getAge: function () {
    return this.age
  },
  getName: function () {
    return this.name
  }
}

var me = new Human('me', 22)
console.log(me.getName())
console.log(me.getAge())



console.log('x'.repeat(25) + ' 静的プロパティ・静的メソッド ' + 'x'.repeat(25))
Human.version = '1.0'
Human.walk = function () {
  return 'walk walk walk'
}
Human.eat = function () {
  return 'eat eat eat'
}
console.log('Human version: ', Human.version)
console.log('Human ', Human.walk())
console.log('Human ', Human.eat())

var Japanese = function () {
  Human.call(this, '山田', 44)
}
Japanese.prototype = new Human()
Japanese.prototype.greeting = function () {
  console.log('こんにちわ')
}

var j = new Japanese()
console.log(j.name)
console.log(j.age)
j.greeting()

for (var key in j) {
  if (j.hasOwnProperty(key)) {
    console.log(key)
  }
}
