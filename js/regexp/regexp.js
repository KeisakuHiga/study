// https://designsupply-web.com/media/knowledgeside/1640/
// 先頭の文字が0-9で０文字以上繰り返すパターンを探す
let pattern = /^.*/
let str = ''
let result = str.match(pattern)
console.log(`\np: ${pattern} \nr:`, result)

// 【名前】全角文字で入力されているか
pattern = /^[ぁ-んァ-ヶー一-龠]+$/
str = 'あっっ'
result = str.match(pattern)
console.log(`\np: ${pattern} \nr:`, result)

// 【ふりがな】全角ひらがなで入力されているか
pattern = /^[あ-ん゛゜ぁ-ぉゃ-ょー「」、]+$/
str = 'わいぁ'
result = str.match(pattern)
console.log(`\np: ${pattern} \nr:`, result)

// 【フリガナ】全角カタカナで入力されているか
// 限定させるには文頭（^）と文末（$）で挟んでおく
// 挟まなかったら、先頭から一部分一致した文字列がマッチしてしまう
pattern = /^[ア-ン゛゜ァ-ォャ-ョー「」、]+$/
str = 'ヘイへイ'
result = str.match(pattern)
console.log(`\np: ${pattern} \nr:`, result)

// 【フリガナ】半角カタカナで入力されているか
pattern = /^[ｱ-ﾝﾞﾟｧ-ｫｬ-ｮｰ｡｢｣､]+$/
str = 'パンﾊﾟﾝ'
result = str.match(pattern)
console.log(`\np: ${pattern} \nr:`, result)