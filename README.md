<br>

<h2 align="center">Task</h2>
<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

</div>

<p align="center">
  A command line to do list written in go.
  <br/>
  <br/>
  <a href="https://github.com/aarctanz/task/issues">Report Bug</a>
  .
  <a href="https://github.com/aarctanz/task/issues">Request Feature</a>
</p>


## Table Of Contents

- [Table Of Contents](#table-of-contents)
- [Features](#features)
- [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)


## Features

- Create, Read, Update, Delete task.
- Stores task locally using sqlite database.

## Built With



- [Go](https://golang.org)
- [Sqlite](https://www.sqlite.org/)

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Go installed locally

### Installation

Clone the repo:

```bash
git clone https://github.com/aarctanz/task
```

```bash
cd task
```

```bash
make
```

### Commands 

- `task` - Show all tasks
- `task -h` - Show list of all commands
- `task -a` - Add new task
- `task -u` - Update task to completed or not completed
- `task -d` - Delete a task


## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.


- Fork the Project
- Create your Feature Branch (git checkout -b feature/AmazingFeature)
- Commit your Changes (git commit -m 'Add some AmazingFeature')
- Push to the Branch (git push origin feature/AmazingFeature)
- Open a Pull Request

## License
Distributed under the [Apache License 2.0](https://github.com/aarctanz/task/blob/main/LICENSE). See LICENSE for more information.