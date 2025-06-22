#!/bin/bash

# Website Generator Script
# Usage: ./generate.sh [markdown|json|watch]

set -e

echo "🌐 Website Generator"
echo "=================="

case "${1:-markdown}" in
    "markdown")
        echo "📝 Generating from markdown files..."
        go run website-generator.go
        echo "✅ Website generated from markdown files!"
        ;;
    "json")
        echo "📄 Generating from JSON file..."
        if [ -f "posts.json" ]; then
            go run website-generator.go --json posts.json
            echo "✅ Website generated from posts.json!"
        else
            echo "❌ posts.json not found. Run with 'markdown' first to generate it."
            exit 1
        fi
        ;;
    "watch")
        echo "👀 Watching for changes in md/ directory..."
        echo "Press Ctrl+C to stop watching"
        
        # Check if fswatch is available
        if command -v fswatch >/dev/null 2>&1; then
            fswatch -o md/ | while read f; do
                echo "🔄 Change detected, regenerating..."
                go run website-generator.go
                echo "✅ Website updated!"
            done
        else
            echo "❌ fswatch not found. Install it with: brew install fswatch"
            echo "   Or use 'markdown' or 'json' mode instead."
            exit 1
        fi
        ;;
    "help"|"-h"|"--help")
        echo "Usage: ./generate.sh [markdown|json|watch]"
        echo ""
        echo "Commands:"
        echo "  markdown  - Generate from markdown files in md/ directory (default)"
        echo "  json      - Generate from posts.json file"
        echo "  watch     - Watch md/ directory for changes and auto-regenerate"
        echo "  help      - Show this help message"
        echo ""
        echo "Examples:"
        echo "  ./generate.sh           # Generate from markdown"
        echo "  ./generate.sh json      # Generate from JSON"
        echo "  ./generate.sh watch     # Watch for changes"
        ;;
    *)
        echo "❌ Unknown command: $1"
        echo "Use './generate.sh help' for usage information."
        exit 1
        ;;
esac 