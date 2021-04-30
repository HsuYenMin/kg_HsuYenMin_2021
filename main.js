var myArgs = process.argv.slice(2);

/* check input format */
if (myArgs.length < 2) {
  console.log("usage: scripts node main.js [num...]\n");
  process.exit();
}

/* the mapping */
const digitcharToPhone = {'0': "Zero", '1': "One", '2': "Two", '3': "Three",
			'4': "Four", '5': "Five", '6': "Six", '7': "Seven",
			'8': "Eight", '9': "Nine"};

/* convert a decimal character to a phone string */
function d2p(digit) {
  if (digit < '0' || digit > '9'){
    /* if it contains a character other than [0-9], exit */
    console.log("usage: scripts node main.js [num...]");
    console.log("[num...] should only contain decimal digits\n");
    process.exit();
  }
  return digitcharToPhone[digit];
}

/* convert a decimal string to a string of phones */
function toPhone(num) {
  return [...num].map(d2p).join('');
}

var phones = myArgs.map(toPhone);

/* show results to the standard output */
console.log(phones.join(','));
