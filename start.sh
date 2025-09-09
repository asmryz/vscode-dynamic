#!/bin/bash
# Install Microsoft C/C++ extension at runtime (license-compliant)
if ! code-server --list-extensions | grep -q "ms-vscode.cpptools"; then
    echo "Installing Microsoft C/C++ extension..."
    code-server --install-extension ms-vscode.cpptools || true
fi

# Start code-server normally
exec code-server "$@"
