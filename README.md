# txtfyer CLI tool

## Overview

The txtfyer cli tool designed to convert any codebase into a single text file. This can be particularly useful for AI model training or integration to the ChatGPT API to adhere to codebase standards or simplifying the process of code analysis and documentation.

## Features

- Convert entire codebases to a single text file with customizable separators;
- Ignore specified files and directories to tailor the output to your needs;
- Command-line interface for easy integration into automation scripts.

## Installation

To install txtfyer, ensure you have Go installed on your system. Then, run the following command:

```bash
go install github.com/danielmesquitta/txtfyer@latest
```

## Usage

To use the cli tool, run the compiled binary with the "txtfyer" command and the necessary flags.

Example:

```bash
txtfyer -p path/to/codebase
```

## Flags:

| Shorthand | Command     | Description                                                                                                                                                   |
| --------- | ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -h        | --help      | help for txtfyer                                                                                                                                              |
| -i        | --ignore    | strings Paths and files to ignore while converting to txt file (default [.git,.env])                                                                          |
| -o        | --output    | string Path to the output txt file (default "output.txt")                                                                                                     |
| -p        | --path      | string Path to the folder with code files (default "./")                                                                                                      |
| -s        | --separator | string Separator to use between files in the output txt file, the %s will be replaced by the relative path of the file in the folder (default "=== %s ===\n") |
