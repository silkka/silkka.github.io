# Website Generator - Solution Summary

## Problem Solved

You mentioned that "It's a big hassle to update this static website with new content." The manual process of:
- Creating individual HTML files for each post
- Updating the main index.html with featured/recent posts
- Updating posts/index.html with all posts
- Managing the post flow manually

## Solution Built

I created a **Go program** (`website-generator.go`) that automates the entire website generation process according to your `AI-Workflow.md` specifications.

## Key Features

### 🔄 **Automatic Workflow Management**
- **Featured Post**: Latest post automatically becomes featured
- **Recent Posts**: Next 3 posts automatically become recent posts  
- **Post Flow**: When adding new posts, everything updates automatically
- **Chronological Order**: All posts sorted by date (newest first)

### 📝 **Multiple Input Formats**
1. **Markdown Files** (existing format)
   ```bash
   ./generate.sh
   # or
   go run website-generator.go
   ```

2. **JSON Files** (easier to edit)
   ```bash
   ./generate.sh json
   # or
   go run website-generator.go --json posts.json
   ```

3. **Watch Mode** (auto-regenerate on changes)
   ```bash
   ./generate.sh watch
   ```

### 🎯 **WCAG 2.1 AA Compliant**
- Maintains accessibility standards
- Proper alt text for images
- Semantic HTML structure
- HTMX integration preserved

## Files Created

1. **`website-generator.go`** - Main Go program
2. **`generate.sh`** - Easy-to-use shell script
3. **`config.json`** - Configuration file
4. **`README-generator.md`** - Comprehensive documentation
5. **`example-new-post.json`** - Example of adding new posts
6. **`posts.json`** - Generated from your markdown files

## How to Use

### Quick Start
```bash
# Generate from existing markdown files
./generate.sh

# Edit posts.json and regenerate
./generate.sh json

# Watch for changes (requires fswatch)
./generate.sh watch
```

### Adding New Content

**Method 1: Edit posts.json**
1. Open `posts.json`
2. Add your new post at the top (it will become featured)
3. Run `./generate.sh json`

**Method 2: Create new markdown file**
1. Create `md/your-new-post.md`
2. Follow the existing format
3. Run `./generate.sh`

**Method 3: Create custom JSON**
1. Create your own JSON file
2. Run `go run website-generator.go --json your-file.json`

## Benefits Achieved

✅ **No more manual HTML editing**  
✅ **Consistent post structure**  
✅ **Automatic featured/recent post management**  
✅ **Easy content management with JSON**  
✅ **Maintains existing design and functionality**  
✅ **WCAG 2.1 AA compliance preserved**  
✅ **HTMX functionality preserved**  
✅ **Category organization**  

## Example Workflow

1. **Add new post**: Edit `posts.json` and add new post at the top
2. **Generate website**: Run `./generate.sh json`
3. **Result**: 
   - New post becomes featured
   - Previous featured moves to recent posts
   - All HTML files updated automatically
   - No manual editing required

## Technical Details

- **Language**: Go (fast, reliable, single binary)
- **Templates**: HTML templates with Go templating
- **Data Flow**: Markdown → JSON → HTML
- **Sorting**: Automatic date-based sorting
- **Categories**: Automatic category organization
- **Slugs**: Automatic URL-friendly slug generation

## Future Enhancements

The program is designed to be easily extensible for:
- RSS feed generation
- Sitemap generation  
- Image optimization
- SEO meta tags
- Multiple language support
- Custom templates per category

---

**Result**: You now have a fully automated website generation system that eliminates the hassle of manual updates while maintaining all your existing functionality and design standards. 