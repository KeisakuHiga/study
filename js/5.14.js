var Member = function () { };
Member.prototype.sex = 'male'
var mem1 = new Member()
var mem2 = new Member()
console.log(mem1.sex, '|', mem2.sex)

mem2.sex = undefined
console.log(mem1.sex, '|', mem2.sex)


