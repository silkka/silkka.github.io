# Website Generator

A Go program that automates the generation of your static website from JSON data. This tool follows the workflow specified in `AI-Workflow.md` and generates all necessary HTML files.

## Features

- **JSON-based Content Management**: All content is managed through `posts.json`
- **Interactive CLI Tool**: Add new posts easily with the `add-post.go` tool
- **Automatic Post Management**: 
  - Latest post becomes featured post
  - Next 3 posts become recent posts
  - All posts listed in chronological order
- **Category Support**: Organizes posts by categories (maps, games, posts, etc.)
- **WCAG 2.1 AA Compliant**: Follows accessibility guidelines
- **HTMX Integration**: Maintains existing HTMX functionality

## Quick Start

### 1. Generate Website

```bash
./generate.sh generate
# or
go run website-generator.go posts.json
```

This will:
- Read posts from `posts.json`
- Generate individual HTML files in appropriate category directories
- Update `index.html` with featured and recent posts
- Update `posts/index.html` with all posts

### 2. Add New Post

```bash
./generate.sh add
# or
go run add-post.go
```

This will:
- Start an interactive session to collect post information
- Add the new post to the front of `posts.json` (making it featured)
- Automatically generate slug and category

## Input Format

### JSON Format

The program uses a `posts.json` file with this structure:

```json
[
  {
    "title": "Post Title",
    "description": "A brief description of your post content",
    "topic": "Maps",
    "date": "Dec 15, 2024",
    "image_url": "https://example.com/image.jpg",
    "external_links": [
      "https://steamcommunity.com/sharedfiles/filedetails/?id=1234567890",
      "https://github.com/silkka/example-map"
    ],
    "slug": "post-title",
    "category": "maps",
    "content": ""
  }
]
```

### Interactive CLI Tool

The `add-post.go` tool prompts for:

- **Post title** (required)
- **Description** (required)
- **Topic** (Maps/Games/Posts, defaults to Maps)
- **Date** (defaults to current date)
- **Image URL** (required)
- **External links** (Steam Workshop, GitHub, etc.)

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
├── posts.json              # Your content data
├── config.json             # Configuration file
├── website-generator.go    # Main generator
└── add-post.go            # CLI tool for adding posts
```

## Adding New Content

### Method 1: Interactive CLI (Recommended)
```bash
./generate.sh add
```
Follow the prompts to add your new post.

### Method 2: Direct JSON Editing
1. Edit `posts.json` to add your new post at the top
2. Run `./generate.sh generate`

### Method 3: Custom JSON File
1. Create your own JSON file with posts
2. Run `go run website-generator.go your-file.json`

## Typical Workflow

```bash
# 1. Add a new post
./generate.sh add

# 2. Generate the website
./generate.sh generate

# 3. Repeat as needed
```

## Benefits

- **No Manual HTML Editing**: All HTML is generated automatically
- **Easy Content Management**: Simple JSON format
- **Interactive Post Creation**: User-friendly CLI tool
- **Automatic Updates**: Featured/recent posts update automatically
- **Category Organization**: Posts are automatically organized by category
- **Accessibility**: Maintains WCAG 2.1 AA compliance
- **HTMX Support**: Preserves existing HTMX functionality

## Troubleshooting

### Common Issues

1. **Date Parsing Errors**: Ensure dates follow the format "Dec 15, 2024"
2. **Missing Images**: Check that image URLs are accessible
3. **Category Issues**: Verify topic names (Maps/Games/Posts)
4. **JSON Syntax**: Ensure `posts.json` has valid JSON syntax

### Debug Mode

Check the console output for parsing errors or validation issues.

## Future Enhancements

- RSS feed generation
- Sitemap generation
- Image optimization
- SEO meta tags
- Multiple language support
- Custom templates per category
- Post editing CLI tool
- Post deletion CLI tool 