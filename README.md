# Parsed CC-CEDICT

This project provides a parsed CSV version of [CC-CEDICT](cc-cedict-parsed.csv)
so that software projects looking to use CEDICT do not have to include their
own parser.

To use CEDICT in your project, download the [parsed CSV](cc-cedict-parsed.csv).

## Formatting notes

The file is formatted as CSV. When multiple English definitions exist, they're
separated by the tab character.

## Todos

- Automate download, parsing, and pushing new release (if any updates) on a regular basis.
- Write documentation for consuming updates.
