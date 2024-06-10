# Contributing to Git Commits Visualizer (gitcs)

Thank you for considering contributing to Git Commits Visualizer! We welcome contributions from the community to help improve and expand this project. To ensure a smooth collaboration, please follow the guidelines outlined below.

## Getting Started

1. **Fork the repository**: Click the "Fork" button at the top of this repository to create your own copy of the project.
2. **Clone the repository**: Clone your fork to your local machine.
3. **Create a new branch**: Create a new branch for your work with a descriptive name using `git switch -c feature-branch-name`.

## Making Changes

1. **Keep your fork up to date**: Regularly pull the latest changes from the main repository to keep your fork up to date. This helps avoid merge conflicts.
2. **Write clear and concise commit messages**: Each commit should have a meaningful message that describes what changes you made and why.
3. **Add tests for new features or bug fixes**: Ensure that any new feature or bug fix includes corresponding tests. This helps maintain the project's stability.

## Testing

We use a script called `setup-test.sh` to generate test data for running unit tests. Please follow these steps to run the tests:

1. **Run the setup script**: Execute the `setup-test.sh` script to generate the `/test_data` folder:
    ```sh
    ./setup-test.sh
    ```
2. **Run the tests**: After generating the test data, run the tests using:
    ```sh
    go test ./...
    ```

Ensure that all tests pass before submitting your changes.

## Submitting Your Contribution

1. **Push to your fork**: Push your changes to your fork on GitHub using `git push origin feature-branch-name`.
2. **Open a Pull Request**: Navigate to the main repository on GitHub and open a Pull Request (PR) from your fork's branch to the `main` branch of the main repository.
3. **Describe your changes**: Provide a detailed description of your changes in the PR. If your PR addresses an issue, mention the issue number in the description (e.g., `Fixes #123`).

## Need Help?

If you need any help or have questions, feel free to open an issue in the repository or reach out to the maintainers.

Thank you for contributing to Git Commits Visualizer!
