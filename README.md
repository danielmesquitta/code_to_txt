# code_to_txt

## Overview

The code_to_txt cli tool designed to convert any codebase into a single text file. This can be particularly useful for AI model training or integration to the ChatGPT API to adhere to codebase standards or simplifying the process of code analysis and documentation.

## Features

Convert entire codebases to a single text file with customizable separators.
Ignore specified files and directories to tailor the output to your needs.
Command-line interface for easy integration into automation scripts.

## Installation

To install code_to_txt, ensure you have Go installed on your system. Then, run the following command:

```bash
go install https://github.com/danielmesquitta/code_to_txt
```

## Usage

To use the cli tool, run the compiled binary with the "code_to_txt" command and the necessary flags.

Example:

```bash
code_to_txt -p path/to/codebase
```

## Flags:

-h, --help help for code_to_txt
-i, --ignore strings Paths and files to ignore while converting to txt file (default [.git,.env])
-o, --output string Path to the output txt file (default "output.txt")
-p, --path string Path to the folder with code files (default "./")
-s, --separator string Separator to use between files in the output txt file, the %s will be replaced by the relative path of the file in the folder (default "=== %s ===\n")
