# Contributing to go-odoo

Thank you for your interest in contributing to `go-odoo`, a Golang wrapper for the Odoo API! We welcome contributions from the community to help improve this library and make it more robust for interacting with Odoo instances. Whether you're fixing bugs, adding features, improving documentation, or suggesting enhancements, your efforts are appreciated.

This document outlines the process and guidelines for contributing to the project. Please read it carefully to ensure a smooth collaboration experience.

## How to Contribute

### Getting Started

* **Fork the Repository:** Click the "Fork" button on the [GitHub repository](https://github.com/Helvethink/go-odoo) to create your own copy of the project.
* **Clone Your Fork:** Clone your forked repository to your local machine:

```bash
git clone https://github.com/<your-username>/go-odoo.git
cd go-odoo
```

* **Set Upstream Remote:** Add the main repository as an upstream remote to keep your fork in sync:

```bash
git remote add upstream https://github.com/Helvethink/go-odoo.git
```

 ### Development Environment

* **Requirements:** Ensure you have Go installed (version 1.x or later, as specified in the project’s `go.mod` file if available).
* **Dependencies:** Install project dependencies:

```bash
go mod tidy
```

* **Testing Setup:** If you plan to test against an Odoo instance, have access to an Odoo server (e.g., version 11 as mentioned in the original repo, or your target version). Update the `ClientConfig` in your test code with appropriate credentials (`Admin`, `Password`, `Database`, `URL`).

### Finding Issues to Work On

* Check the [Issues tab](https://github.com/Helvethink/go-odoo/issues) for open issues labeled `help wanted` or `good first issue`.
* If you have an idea for a new feature or improvement, open an issue to discuss it with the maintainers before starting work.

### Making Changes

* **Create a Branch:** Work on a new branch for each contribution:

```bash
git checkout -b feature/<your-feature-name>
```

* **Code Style:** Follow Go conventions (e.g., use gofmt for formatting and go vet for static analysis).
* **Commit Messages:** Write clear, concise commit messages:
  * Use the present tense (e.g., "Add support for Odoo 14 API" instead of "Added support").
  * Reference related issues (e.g., "Fixes #42").

Testing: Add or update tests in the *_test.go files to cover your changes. Run tests locally:

```bash
go test ./...
```

### Submitting Your Contribution

* **Push Your Changes:** Push your branch to your fork:

```bash
git push origin feature/<your-feature-name>
```

* Open a Pull Request (PR):
  * Go to the [Pull Requests](https://github.com/Helvethink/go-odoo/pulls) tab in the original repository.
  * Click "New Pull Request" and select your branch.
  * Provide a detailed description of your changes, including:
     * What you changed and why.
     * How to test your changes.
     * Any related issue numbers (e.g., "Closes #42").
* **Code Review:** Be responsive to feedback from maintainers. Make requested changes by pushing additional commits to your branch.

### Syncing with Upstream

* Keep your fork up to date with the main repository:

```bash
git fetch upstream
git rebase upstream/master
git push --force
```

## Guidelines

* **Scope:** This project focuses on providing a Golang wrapper for Odoo’s API. Contributions should align with this goal (e.g., improving model generation, adding support for newer Odoo versions, enhancing usability).
* **Model Generation:** If adding new models, follow the existing pattern in `ir_model.go` and `ir_model_fields.go`. Regenerate models if necessary, as outlined in the README.
* **Compatibility:** Note that model structures may differ across Odoo versions. Specify the Odoo version your changes target in your PR.
* **Documentation:** Update the README or add inline comments for significant changes, especially if they affect usage.
* **License:** Contributions are assumed to be under the same license as the project (check the repository for the specific license, likely MIT or similar, unless specified otherwise).

## Reporting Bugs

* Use the [Issues tab](https://github.com/Helvethink/go-odoo/issues) to report bugs.
* Include:
  * Odoo version tested against.
  * Steps to reproduce the issue.
  * Expected vs. actual behavior.
  * Any error messages or logs.

## Questions or Support

* For general questions, open an issue with the `question` label.
* Reach out to the maintainers via GitHub for clarification on contribution processes.

## Code of Conduct

We aim to foster an inclusive and respectful community. Please be kind and considerate in all interactions.
