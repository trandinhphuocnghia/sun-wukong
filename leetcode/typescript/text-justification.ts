/**
 * Problem: https://leetcode.com/problems/text-justification/description/?envType=study-plan-v2&envId=top-interview-150.
 * My Approach:
 * from each word in Words:
 * Step1, Find the words (array) can create the string (1):
 *  IF it in previous string so continue,ELSE check every next word that can create string with number of character < maxWidth.
 * Step2, Add the space:
 *  IF contain the last word:
 *      left justify.
 *  ELSE:
 *      Justify-between, but if have extraSpace add all to the first words.
 * @param words
 * @param maxWidth
 * @returns
 */
function fullJustify(words: string[], maxWidth: number): string[] {
  const result: string[] = [];
  let passedIndex = -1;

  for (let i = 0; i < words.length; i++) {
    if (i <= passedIndex) {
      continue; // Skip already processed words
    }

    // Get possible words
    let wordLength = words[i].length;
    const streakOfWord = [words[i]];
    let containLast = false;

    for (let y = i + 1; y < words.length; y++) {
      if (wordLength + words[y].length + streakOfWord.length <= maxWidth) {
        wordLength += words[y].length;
        streakOfWord.push(words[y]);
        passedIndex = y;

        if (y == words.length - 1) {
          containLast = true;
        }
      } else {
        break;
      }
    }

    // Justify it by spacing
    const streakLength = wordLength;
    const space = maxWidth - streakLength;
    const gaps = streakOfWord.length - 1;

    if (containLast || gaps === 0) {
      // Left-justify if it's the last line or there's only one word
      const text =
        streakOfWord.join(" ") + " ".repeat(maxWidth - (streakLength + gaps));
      result.push(text);
    } else {
      // Fully justify
      const minSpace = Math.floor(space / gaps);
      const extraSpace = space % gaps;

      const text = streakOfWord.reduce((acc, curr, index) => {
        if (index === streakOfWord.length - 1) {
          return acc + curr;
        }
        const spacesToAdd = minSpace + (index < extraSpace ? 1 : 0);
        return acc + curr + " ".repeat(spacesToAdd);
      }, "");

      result.push(text);
    }
  }

  return result;
}
