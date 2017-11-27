# Issues2Markdown

`Issues2Markdown` is a command line tool that you can use to render your exported Bitbucket issues as Markdown.

## TODO

[ ] Add ability to select which tickets to output or not output. Like based on priority

[ ] Add ability to go directly to Bitbucket and get the issue list


## Why Do I Need This?

Sometimes when working with remote teams they are reluctant to use an issue tracking tool. This make it easy for you to keep track of issues in your Bitbucket account and share them view Markdown or via email.


## How To Use

### Export Issues

First you need to export your issues from your Bitbucket project.


1. Open you Bitbucket project
2. Click `Settings`
3. Scroll down to `Issues`
4. Select `Import & Export`
5. Select the `Start Export` button
6. Download zip file containing your issues
7. Unzip file

This tool uses `./db-1.0.json` to create the Markdown document.

### Render Issues as Markdown


Example Usage:

`./Issues2Markdown --file ./db-1.0.json |tee output.md`

The output is rendered to standard out. You can pipe that to a file to save it. Here is another example that outputs the markdown to a file and opens it with your Markdown reader:

`./Issues2Markdown --file ./db-1.0.json |tee output.md; open output.md`