var Member = function () { }

Member.prototype.sex = 'male'

var mem = new Member()
mem.sex = undefined

for (var key in mem) {
  console.log(key + ': ' + mem[key])
}
