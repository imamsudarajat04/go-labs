#!/bin/bash

# Install Artisan CLI alias to shell profile
# This script will add the artisan alias to your shell profile

SHELL_PROFILE=""
CURRENT_SHELL=$(basename "$SHELL")

# Determine shell profile file
case "$CURRENT_SHELL" in
    "bash")
        SHELL_PROFILE="$HOME/.bashrc"
        ;;
    "zsh")
        SHELL_PROFILE="$HOME/.zshrc"
        ;;
    *)
        echo "❌ Unsupported shell: $CURRENT_SHELL"
        echo "Please manually add the alias to your shell profile"
        exit 1
        ;;
esac

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Check if alias already exists
if grep -q "alias artisan=" "$SHELL_PROFILE" 2>/dev/null; then
    echo "⚠️  Artisan alias already exists in $SHELL_PROFILE"
    echo "Please remove the existing alias manually if you want to update it"
    exit 0
fi

# Add alias to shell profile
echo "" >> "$SHELL_PROFILE"
echo "# Artisan CLI alias for Go Labs" >> "$SHELL_PROFILE"
echo "alias artisan='cd \"$SCRIPT_DIR\" && ./artisan'" >> "$SHELL_PROFILE"

echo "✅ Artisan alias added to $SHELL_PROFILE"
echo ""
echo "To use the alias, either:"
echo "1. Restart your terminal, or"
echo "2. Run: source $SHELL_PROFILE"
echo ""
echo "Then you can use: artisan make:model [name] -m"
