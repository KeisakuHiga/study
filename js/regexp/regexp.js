// https://designsupply-web.com/media/knowledgeside/1640/
/**
 * str.match(pattern)のログを出す関数
 * @param {RegExp} pattern 
 * @param {String} str 
 */
const regexpMatchPattern = (pattern, str) => {
  const result = str.match(pattern)
  console.log(`\np: ${pattern} \nr:`, result)
}

// 先頭の文字が0-9で０文字以上繰り返すパターンを探す
let pattern = /^.*/
let str = ''
regexpMatchPattern(pattern, str)

// 【名前】全角文字で入力されているか
pattern = /^[ぁ-んァ-ヶー一-龠]+$/
str = 'あっっ'
regexpMatchPattern(pattern, str)

// 【ふりがな】全角ひらがなで入力されているか
pattern = /^[あ-ん゛゜ぁ-ぉゃ-ょー「」、]+$/
str = 'わいぁ'
regexpMatchPattern(pattern, str)

// 【フリガナ】全角カタカナで入力されているか
// 限定させるには文頭（^）と文末（$）で挟んでおく
// 挟まなかったら、先頭から一部分一致した文字列がマッチしてしまう
pattern = /^[ア-ン゛゜ァ-ォャ-ョー「」、]+$/
str = 'ヘイへイ'
regexpMatchPattern(pattern, str)

// 【フリガナ】半角カタカナで入力されているか
pattern = /^[ｱ-ﾝﾞﾟｧ-ｫｬ-ｮｰ｡｢｣､]+$/
str = 'パンﾊﾟﾝ'
regexpMatchPattern(pattern, str)

// 【【メールアドレス】アットマークを含む正しい形式で入力されているか
pattern = /^([a-zA-Z0-9])+([a-zA-Z0-9\._-])*@([a-zA-Z0-9_-])+([a-zA-Z0-9\._-]+)+$/
str = '0@__'
regexpMatchPattern(pattern, str)


// 【【電話番号】半角数字で入力されているか
pattern = /^[0-9]+$/
str = '72131'
regexpMatchPattern(pattern, str)

// 【電話番号】半角ハイフンを含んだ1〜4桁・1〜4桁・3〜4桁の半角数字の形式で入力されているか
pattern = /^0\d{1,4}-\d{1,4}-\d{3,4}/
str = '07-2-131'
regexpMatchPattern(pattern, str)

