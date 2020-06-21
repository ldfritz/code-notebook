declare function require(name:string);

const process = require("process");

function greeter(person: string): string {
	return `Hello, ${person}!`;
}

let user: string = "Luke";

if (process.argv.length > 2) {
	user = process.argv.slice(2).join(" ");
}

process.stdout.write(`${greeter(user)}\n`);
