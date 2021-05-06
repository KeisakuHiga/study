var Member = function (firstName, lastName) {
  this.firstName = firstName
  this.lastName = lastName
}

Member.prototype = {
  getName: function () {
    return this.firstName + ' ' + this.lastName
  },
  toString: function () {
    return this.lastName + this.firstName
  }
}

var member = new Member('Keisaku', 'HIGA')
console.log(member.getName())
console.log(member.toString())

