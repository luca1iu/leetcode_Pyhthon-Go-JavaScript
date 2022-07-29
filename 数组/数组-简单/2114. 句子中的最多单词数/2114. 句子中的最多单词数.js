/**
 * @param {string[]} sentences
 * @return {number}
 */
 var mostWordsFound = function(sentences) {
    var maxLen = 0;
    for (i=0; i<sentences.length; i++){
        let len = sentences[i].split(" ").length;
        maxLen = maxLen >len? maxLen : len;
    }
    return maxLen
};