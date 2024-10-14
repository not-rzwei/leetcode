# Go LeetCode Go

Welcome to my personal project for my LeetCode solutions in Go! This repository serves as a playground for reviewing DSA and exploring Go features that I rarely use in my work. It's all about adventure and learning!

## Requirements

- Visual Studio Code
- Go 1.21+
- LeetCode VSCode extension

## Project Structure

The project is organized in a way that allows for multiple solutions to the same problem using build tags. Each solution is encapsulated within its own file, and you can switch between them as needed.

### Generate Solution File
To generate a solution with the same file name and structure, copy the configuration in the `.vscode/leetcode.json` to your user setting JSON.

### Build Tags

To test a different solution for a problem, you will need to change the `buildTags` in the `.vscode/settings.json` file. After updating the build tags, run the following command to ignore any changes to this file:

```bash
git update-index --assume-unchanged .vscode/settings.json
```

Example tag to test bruteforce solutions
```json
{
    "go.buildTags": "bruteforce",
    ...
}
```