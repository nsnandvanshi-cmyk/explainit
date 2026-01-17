package docs

const OutputHelp = `
EXPLAINIT OUTPUT GUIDE

1. Files indexed
   Number of files discovered during scan.

2. Functions detected
   Total functions found using pattern analysis.

3. Function calls
   Shows which functions are called inside another function.

4. Called by
   Shows which functions depend on a given function.

5. Impact analysis
   Functions that may break if the target function changes.

NOTES:
- Results are best-effort (static analysis).
- No code execution is performed.
- Simple logic is used for portability and speed.
`
