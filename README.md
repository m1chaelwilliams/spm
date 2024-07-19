# SPM | Simple Project Manager

My personal terminal-based project manager.

# Why?

I'm a person that likes to start projects every other day. I have lists and lists of ideas that don't make it off the ground. As a result, it's hard to keep track of what I am currently focused on. I lose where I put a project, or I forget about it all together. __SPM__ was built to fix that. Each project is tracked in the terminal, where I am most comfortable. That means I don't need to open File Explorer to look through pages of nonsense to find what I want.

# Features

|Feature|Status|
|-------|------|
|`add`|✅|
|`remove`|✅|
|`list`|✅|
|`copypath`|✅|
|`contains`|✅|
|`help`|✅|
|`lookup`|✅|
|`init`|✅|
|`load`|✅|

For more information, run `tpm --help` to get detailed instructions for each command

# Installation

1. Clone the repository:
```shell
git clone https://github.com/m1chaelwilliams/tpm.git
```
2. Add it to your __PATH__ environment variables
3. Spin up the database:
```shell
spm spinup
```

# Example Usage

All of the commands are intuitive with useful shortcuts to make things faster. For example, adding a project can look like:

```shell
# option 1
spm add -name="example" -path="C:\example\path"
# option 2
spm add . -name="example"
# option 3 (must be in the directory to start)
spm add .
```

There are also options for loading projects from __JSON__:

```shell
# create a skeleton json project file in cwd
spm init
# load that json file into the database
spm load
```

Here's what the JSON file could look like:

```json
{
	"name":"",
	"path":"",
	"metadata":{}
}
```

# Copyright

This repository is licensed under [MIT](./LICENSE).

# Feature Request

This is my own personal project manager tool. So, if a feature is to be added, I must like it. Please feel free to request them though!