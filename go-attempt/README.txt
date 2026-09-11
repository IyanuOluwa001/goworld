Instruction:
receive as arguments the name of a file containing a text that needs some modifications (the input)
and the name of the file the modified text should be placed in (the output). 

Needed modifications:
HexToDecimal
1. hex: replace the word before an hexadecimal number(hex) with the decimal version of the word.
1E - 30

BinToDecimal
2. bin: replace the word before a binary number(bin) with the decimal version of the word
e.g: 10 - 2

Capitalize
3. up: converts the word before (up) with the Uppercase version of it.
e.g: down - DOWN

4. low: converts the word before (low) with the lowercase version of it.
e.g: UP - up

5. cap: coverts the word before (cap) with the capitalized version of it.
e.g: bridge - Bridge

6. For (low), (up), (cap) if a number appears next to it, like so:
(low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

7. for . , ! ? : and ;  should not have space before it, but must have space after it

8. if we have two or more punctuation together, no formatting should apply

9a. for punctuation like ' ', no space after the first quotation and no space before the second quotation if one word is inside the quote.
9b. same with multiple words within the quote.

10. detect if a exist alone, detect if the word after it start with vowels a e i o u then turn a to an.