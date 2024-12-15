# [GitHub Activity CLI](https://roadmap.sh/projects/github-user-activity)

A command-line tool that fetches and displays GitHub user activity events.

## Features

- Retrieves recent GitHub events for a specified username
- Displays event type, repository name, URL, and creation timestamp
- Simple command-line interface

## Installation

```sh
go build -o <location_to_store_binary>/github-activity main.go
```

## Usage

```sh
github-activity <username>
```


## Example Output
```sh
Event: PushEvent
Repository: user/repo-name
Repo URL: https://github.com/user/repo-name
Created At: 2023-07-21T15:30:00Z
```
