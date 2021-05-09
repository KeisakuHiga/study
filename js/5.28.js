var Wings = Wings || {};

// クラスのコンストラクターを名前空間（Wings）を指定してつくる
Wings.Member = function (first, last) {
  this.first = first
  this.last = last
}

Wings.Member.prototype = {
  getName: function () {
    return this.first + ' '+ this.last
  }
}

var wMem = new Wings.Member('keisaku', 'higa')
console.log(wMem.getName())