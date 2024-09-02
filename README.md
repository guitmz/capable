<p align="center">
    <h1 align="center">Capable</h1>
</p>
<p align="center">
    <em>Interact with <a href="https://capacities.io/">Capacities.io</a> API from your terminal.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/GitHub%20Actions-2088FF.svg?style=default&logo=GitHub-Actions&logoColor=white" alt="GitHub%20Actions">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=default&logo=Go&logoColor=white" alt="Go">
</p>
<br>

#####  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Getting Started](#-getting-started)
    - [ Installation](#-installation)
    - [ Usage](#-usage)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

The CLI tool is designed to interact with the Capacities.io API, enabling users to send daily entries and weblinks to the Capacities platform directly from the command line. Capacities.io is a platform focused on personal knowledge management and note-taking, allowing users to create, organize, and connect information efficiently.

---

##  Features

|    |   Feature         | Description |
|----|-------------------|---------------------------------------------------------------|
| 📝 | **Daily Notes**  | Append text to the daily note for the current date. This is useful for journaling or quickly adding thoughts. Supports inputs from the command line or `stdin` with Markdown formatting and automatic timestamping. |
| 🔗 | **Weblinks**  | Add a weblink to Capacities from your command line. |

---

##  Getting Started

### Requirements

Capable requires `CAPACITIES_SPACE_ID` and `CAPACITIES_API_TOKEN` environment variables to be defined for authentication. Please refer to the official documentation for more information.

###  Installation

The installation options are the following:

- Download the pre-compiled binary from the releases page
- Use Homebrew `brew install guitmz/tools/capable`
- Or follow the guide below to build the project from source:

1. Clone the app repository:
```sh
❯ git clone https://github.com/guitmz/capable.git
```

2. Navigate to the project directory:
```sh
❯ cd capable
```

3. Install the required dependencies:
```sh
❯ go build -o capable
```

###  Usage

To run the project, execute the following command:

```sh
❯ ./capable
```

---

##  Project Roadmap

- [X] **`Weblinks`**: <strike>Implement sending weblinks.</strike>
- [X] **`Daily Notes`**: <strike>Implement adding to daily notes.</strike>
- [ ] **`Search`**: Implement search feature.

---

##  Contributing

Contributions are welcome! Please submit a pull request or open an issue, clearly describing the changes and their motivations.

---

##  License

This project is protected under the [MIT](https://choosealicense.com/licenses/mit/) License.
---

##  Acknowledgments

- [Capacities.io API documentation](https://docs.capacities.io/developer/api)
- [Capacities.io API reference](https://api.capacities.io/docs/)

---
