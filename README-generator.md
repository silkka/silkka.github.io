# Website Generator

A Go program that automates the generation of your static website from markdown files or JSON data. This tool follows the workflow specified in `AI-Workflow.md` and generates all necessary HTML files.

## Features

- **Markdown to HTML**: Converts markdown files to individual blog post HTML pages
- **JSON Support**: Can work with JSON files for easier content management
- **Automatic Post Management**: 
  - Latest post becomes featured post
  - Next 3 posts become recent posts
  - All posts listed in chronological order
- **Category Support**: Organizes posts by categories (maps, games, posts, etc.)
- **WCAG 2.1 AA Compliant**: Follows accessibility guidelines
- **HTMX Integration**: Maintains existing HTMX functionality

## Quick Start

### 1. Generate from Markdown Files

```bash
go run website-generator.go
```

This will:
- Read all `.md` files from the `md/` directory
- Generate individual HTML files in appropriate category directories
- Update `index.html` with featured and recent posts
- Update `posts/index.html` with all posts
- Create `posts.json` for easier future editing

### 2. Generate from JSON File

```bash
go run website-generator.go --json posts.json
```

This allows you to edit the JSON file directly and regenerate the website.

## Input Formats

### Markdown Format

Create markdown files in the `md/` directory following this structure:

```markdown
# Post Title

- Description:
    - A brief description of your post content

- Topic: Main topic/category of the post

- Date: Use format: Feb 24, 2021

- Image for the post: URL to featured image

- External links: 
    - https://example.com/link1
    - https://example.com/link2
```

### JSON Format

The program generates a `posts.json` file that you can edit directly:

```json
[
  {
    "title": "Post Title",
    "description": "A brief description of your post content",
    "topic": "Main topic/category",
    "date": "Feb 24, 2021",
    "image_url": "https://example.com/image.jpg",
    "external_links": [
      "https://example.com/link1",
      "https://example.com/link2"
    ],
    "slug": "post-title",
    "category": "maps",
    "content": ""
  }
]
```

## Configuration

Edit `config.json` to customize:

- Site title and description
- Category definitions
- Template paths
- Output settings
- Feature flags

## Workflow Integration

The generator follows the AI-Workflow.md specifications:

1. **Featured Post**: Latest post becomes the featured post
2. **Recent Posts**: Next 3 posts become recent posts
3. **Post Updates**: When adding new posts, the system automatically:
   - Makes the new post featured
   - Moves current featured to recent posts
   - Shifts existing recent posts down
   - Removes oldest from recent posts
4. **All Posts**: Updates the full post list in chronological order

## File Structure

After generation, your website will have:

```
/
├── index.html              # Main page with featured + recent posts
├── posts/
│   └── index.html          # All posts listing
├── maps/                   # Category directories
│   ├── post-slug-1.html
│   └── post-slug-2.html
├── games/
│   └── game-post.html
├── posts.json              # Generated for easy editing
└── config.json             # Configuration file
```

## Adding New Content

### Method 1: Markdown Files
1. Create a new `.md` file in the `md/` directory
2. Follow the markdown format above
3. Run `go run website-generator.go`

### Method 2: JSON Editing
1. Edit `posts.json` to add your new post
2. Run `go run website-generator.go --json posts.json`

### Method 3: Direct JSON Creation
1. Create a new JSON file with your posts
2. Run `go run website-generator.go --json your-file.json`

## Benefits

- **No Manual HTML Editing**: All HTML is generated automatically
- **Consistent Structure**: All posts follow the same template
- **Easy Content Management**: JSON format is easier to edit than HTML
- **Automatic Updates**: Featured/recent posts update automatically
- **Category Organization**: Posts are automatically organized by category
- **Accessibility**: Maintains WCAG 2.1 AA compliance
- **HTMX Support**: Preserves existing HTMX functionality

## Troubleshooting

### Common Issues

1. **Date Parsing Errors**: Ensure dates follow the format "Feb 24, 2021"
2. **Missing Images**: Check that image URLs are accessible
3. **Category Issues**: Verify category names match those in `config.json`

### Debug Mode

Add debug logging by modifying the Go code or check the console output for parsing errors.

## Future Enhancements

- RSS feed generation
- Sitemap generation
- Image optimization
- SEO meta tags
- Multiple language support
- Custom templates per category 