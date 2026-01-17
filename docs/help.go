package docs

const HelpText = `
explainit — Understand any codebase faster

USAGE:
  explainit <command> [arguments]

COMMANDS:
  scan
      Scan and index the current repository

  files
      List all indexed files

  functions
      List all detected functions

  stats
      Show repository statistics

  calls <function>
      Show functions called by a function

  called-by <function>
      Show functions that call a function

  impact <function>
      Show impact of changing a function

  output help
      Explain the meaning of explainit outputs

FLAGS:
  -h, --help
      Show this help message
`
