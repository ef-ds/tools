# cmdrunner

This tool runs all the commands provided in the --input parameter and outputs the results to the console or into the supplied --output file.

Each command is limited to a single line.

Lines that starts with '//' (comment) are not executed, but their content is written either to the console or the the specified output file.

The result of each command is output with "```" before and after the result.
