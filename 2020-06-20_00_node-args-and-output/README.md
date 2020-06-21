# Node args and output

In this very short script, I play around with using TypeScript and Node.

I also touch on simple input and output, which is so common to command line utilities.

Line 1 is a fix for the "TS2580: Cannot find name 'require'" error that line 3 would otherwise produce.  This solution is lifted from https://stackoverflow.com/a/12742371.  Line 1 declares a `require` function with a `string` parameter.  This compiles successfully.  It checks that the argument on line 3 is a `string` and then deletes line 1.  And Node understands the resulting code.

I believe an alternative solution would be to install @types/node with npm and use an `import` instead of a `require`.

On lines 6 and 15, I play with the backtick strings and string interpolation.  I tend to prefer interpolation to concatenation using plus signs, when adding any more than two strings.  I like that TypeScript includes this.

On lines 11 and 12, I use `process.argv` to check and use the command line arguments.  `process.argv` basically returns everything typed to run this script.  This includes both the executing program, `node`, and the executed script, `greeter.js`.  Because of this, I skip the first two values in the `process.argv` array.

I do wonder what would happen if I executed this in other ways, like with a Node shebang.

It is also worth noting that `process.argv[0]` is actually expanded to its absolute path.  To get the original value as entered, use `process.argv0`.

On line 15, I use `process.stdout` to print to the terminal.  This could have just as easily have been `console.log` and I wouldn't have had to add the newline character.  But I like using stdin, stdout, and stderr and appreciate being able to be very explicit even here where I am only using stdout.
