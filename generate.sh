#!/bin/bash

# Website Generator Script
# Usage: ./generate.sh [generate|add|help]

set -e

echo "🌐 Website Generator"
echo "=================="

case "${1:-generate}" in
    "generate")
        echo "📄 Generating website from posts.json..."
        if [ -f "posts.json" ]; then
            go run website-generator.go posts.json
            echo "✅ Website generated successfully!"
        else
            echo "❌ posts.json not found. Use './generate.sh add' to create your first post."
            exit 1
        fi
        ;;
    "add")
        echo "📝 Adding new post..."
        go run add-post.go
        echo "✅ Post added! Run './generate.sh generate' to update the website."
        ;;
    "help"|"-h"|"--help")
        echo "Usage: ./generate.sh [generate|add|help]"
        echo ""
        echo "Commands:"
        echo "  generate  - Generate website from posts.json (default)"
        echo "  add       - Add a new post interactively"
        echo "  help      - Show this help message"
        echo ""
        echo "Examples:"
        echo "  ./generate.sh           # Generate website"
        echo "  ./generate.sh generate  # Generate website"
        echo "  ./generate.sh add       # Add new post"
        echo ""
        echo "Workflow:"
        echo "  1. ./generate.sh add    # Add new post"
        echo "  2. ./generate.sh        # Generate website"
        echo "  3. Repeat as needed"
        ;;
    *)
        echo "❌ Unknown command: $1"
        echo "Use './generate.sh help' for usage information."
        exit 1
        ;;
esac 