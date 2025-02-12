# kubectl-protected

`kubectl-protected` is a `kubectl` plugin that verifies if the current context points to a protected cluster, helping prevent accidental actions.

## Installation

To install the `kubectl-protected` plugin, you have two options: compile from source or download from the releases page.

### Option 1: Compile from Source

1. **Clone the repository:**

    ```sh
    git clone https://github.com/bivas/kubectl-protected.git
    cd kubectl-protected
    ```

2. **Build the plugin:**

    ```sh
    go build -o kubectl-protected cmd/protected/main.go
    ```

3. **Move the binary to a directory in your PATH:**

    ```sh
    mv kubectl-protected /usr/local/bin/
    ```

4. **Verify the installation:**

    ```sh
    kubectl protected --help
    ```

### Option 2: Download from Releases

1. **Visit the [releases page](https://github.com/bivas/kubectl-protected/releases) of the project.**

2. **Download the appropriate binary for your operating system.**

3. **Move the binary to a directory in your PATH:**

    ```sh
    mv kubectl-protected /usr/local/bin/
    ```

4. **Verify the installation:**

    ```sh
    kubectl protected --help
    ```

## Usage

To use the `kubectl-protected` plugin, you need to create a YAML file that lists the protected clusters. By default, the plugin looks for this file at `$HOME/.kube/protected.yaml`.

### Example `protected.yaml`:

```yaml
protected:
  - cluster1
  - .*pattern.*
```

### Running the plugin

To check if the current Kubernetes cluster is protected, run:

```sh
kubectl protected
```

### Command-line options

- `--silence-error-on-protected`: Exit without error if the cluster is protected.
- `--protected-file-path`: Path to the file containing the list of protected clusters (default: `$HOME/.kube/protected.yaml`).

### Example

```sh
kubectl protected --protected-file-path /path/to/protected.yaml --silence-error-on-protected
```

This command will check if the current cluster is protected based on the patterns defined in the specified YAML file and will not exit with an error if the cluster is protected.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```