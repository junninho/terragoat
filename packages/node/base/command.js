const { exec } = require('child_process');

function badCommand(input) {
  exec(`ls -l ${input}`);
}

function badEval(input) {
  eval(`console.log("${input}")`);
}

function badAssign(input) {
  let data = {};
  Object.assign(data, JSON.parse(input));
}

function badHtml(input) {
    return `<h1>${input}</h1>`;
} 