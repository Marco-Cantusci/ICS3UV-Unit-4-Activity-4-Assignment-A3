/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-08
 * @fileoverview Is the word in the given sentence
 */

// declare variables
let exist: boolean = false;

// input
// sentence
const sentence: string = prompt("Please enter a sentence: ") || ("\n");

// word to be searched
const word: string = prompt("Enter a word to search in your sentence: ") || ("\n");

// find if the word is in the sentence 
for (let counter = 0; counter <= sentence.length - word.length; counter++) {
  let match: boolean = true;

  for (let notFound = 0; notFound < word.length; notFound++) {
    if (sentence[counter + notFound] !== word[notFound]) {
      match = false;
    }
  }

  if (match == true) {
    exist = true;
  }
}

// output
if (exist == true) {
  console.log(`Hooray, the word ${word} was found in the sentence: ${sentence}`);
} else {
  console.log(`Sorry, the word ${word} was not found in the sentence: ${sentence}`);
}

console.log("\nDone.");
