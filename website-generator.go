package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Post represents a blog post with all its metadata
type Post struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Topic        string   `json:"topic"`
	Date         string   `json:"date"`
	ImageURL     string   `json:"image_url"`
	ExternalLinks []string `json:"external_links"`
	Slug         string   `json:"slug"`
	Category     string   `json:"category"`
	Content      string   `json:"content"`
}

// WebsiteData holds all the data needed to generate the website
type WebsiteData struct {
	FeaturedPost Post   `json:"featured_post"`
	RecentPosts  []Post `json:"recent_posts"`
	AllPosts     []Post `json:"all_posts"`
}

// Templates for HTML generation
const postTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} - Silkka's Blog</title>
    <link rel="stylesheet" href="../styles.css">
    <script src="/components/theme.js"></script>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
</head>

<body>
    <div hx-get="/components/navbar.html" hx-trigger="load" hx-swap="outerHTML"></div>

    <div class="blog-page">
        <div class="blog-post">
            <h1>{{.Title}}</h1>
            <date class="post-content">{{.Date}}</date>

            <img src="{{.ImageURL}}" alt="{{.Title}} Screenshot" class="featured-image">
            <div class="post-content">
                <h2>Description</h2>
                <p>{{.Description}}</p>

                {{if .ExternalLinks}}
                <!-- Optional CTA Section -->
                <div class="cta-section">
                    <h2>External Links</h2>
                    {{range .ExternalLinks}}
                    <a href="{{.}}" class="cta-button" target="_blank" rel="noopener noreferrer">View Resource</a>
                    {{end}}
                </div>
                {{end}}
            </div>
        </div>
    </div>
</body>

</html>`

const indexTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Silkka's blog</title>
    <link rel="stylesheet" href="styles.css">
    <script src="/components/theme.js"></script>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
</head>

<body>
    <div hx-get="/components/navbar.html" hx-trigger="load" hx-swap="outerHTML"></div>

    <main>
        <section class="featured-post">
            <h1>Latest Post</h1>
            <article class="post-card featured">
                <div class="highlighted-image">
                    <img src="{{.FeaturedPost.ImageURL}}" alt="{{.FeaturedPost.Title}} Screenshot">
                    <span class="date">{{.FeaturedPost.Date}}</span>
                </div>
                <h2>{{.FeaturedPost.Title}}</h2>
                <p>{{.FeaturedPost.Description}}</p>
                <a href="/{{.FeaturedPost.Category}}/{{.FeaturedPost.Slug}}.html" class="read-more" hx-get="/{{.FeaturedPost.Category}}/{{.FeaturedPost.Slug}}.html" hx-target="body" hx-push-url="true" aria-label="Read more about {{.FeaturedPost.Title}}">READ MORE</a>
            </article>
        </section>

        <section class="latest-posts">
            <h2>Recent Posts</h2>
            <div class="posts-grid">
                {{range .RecentPosts}}
                <article class="post-card">
                    <div class="post-image">
                        <img src="{{.ImageURL}}" alt="{{.Title}} Screenshot">
                        <span class="date">{{.Date}}</span>
                    </div>
                    <h3>{{.Title}}</h3>
                    <p>{{.Description}}</p>
                    <a href="/{{.Category}}/{{.Slug}}.html" class="read-more" aria-label="Read more about {{.Title}}">READ MORE</a>
                </article>
                {{end}}
            </div>
        </section>

        <div class="view-all-posts">
            <a href="/posts/index.html" class="read-more">VIEW ALL POSTS</a>
        </div>
    </main>
</body>

</html>`

const postsIndexTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>All Posts - My Technology Blog</title>
    <link rel="stylesheet" href="../styles.css">
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <script src="/components/theme.js"></script>
</head>

<body>
    <div hx-get="/components/navbar.html" hx-trigger="load" hx-swap="outerHTML"></div>

    <main class="featured-post">
        <h1>All Posts</h1>
        <div class="posts-grid">
            {{range .AllPosts}}
            <article class="post-card">
                <div class="post-image">
                    <img src="{{.ImageURL}}" alt="{{.Title}} Screenshot">
                    <span class="date">{{.Date}}</span>
                </div>
                <h2>{{.Title}}</h2>
                <p>{{.Description}}</p>
                <a href="/{{.Category}}/{{.Slug}}.html" class="read-more" aria-label="Read more about {{.Title}}">READ MORE</a>
            </article>
            {{end}}
        </div>
    </main>
</body>

</html>`

// generateSlug creates a URL-friendly slug from a title
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")
	slug = strings.ReplaceAll(slug, "'", "")
	slug = strings.ReplaceAll(slug, ".", "")
	return slug
}

// determineCategory determines the category based on topic
func determineCategory(topic string) string {
	switch strings.ToLower(topic) {
	case "maps", "map":
		return "maps"
	case "games", "game":
		return "games"
	case "posts", "post":
		return "posts"
	default:
		return "maps" // default category
	}
}

// parseDate parses the date string and returns a time.Time
func parseDate(dateStr string) (time.Time, error) {
	// Handle different date formats
	formats := []string{
		"Jan 2, 2006",
		"Jan 02, 2006",
		"2 Jan, 2006",
		"02 Jan, 2006",
	}
	
	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}
	
	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

// loadPostsFromJSON loads posts from JSON file
func loadPostsFromJSON(inputFile string) ([]Post, error) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	var posts []Post
	err = json.Unmarshal(data, &posts)
	return posts, err
}

// generatePostHTML generates the HTML for a single post
func generatePostHTML(post Post, outputDir string) error {
	tmpl, err := template.New("post").Parse(postTemplate)
	if err != nil {
		return err
	}

	// Create category directory if it doesn't exist
	categoryDir := filepath.Join(outputDir, post.Category)
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		return err
	}

	// Generate HTML file
	outputFile := filepath.Join(categoryDir, post.Slug+".html")
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, post)
}

// generateIndexHTML generates the main index.html
func generateIndexHTML(data WebsiteData, outputDir string) error {
	tmpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}

	outputFile := filepath.Join(outputDir, "index.html")
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

// generatePostsIndexHTML generates the posts/index.html
func generatePostsIndexHTML(data WebsiteData, outputDir string) error {
	tmpl, err := template.New("postsIndex").Parse(postsIndexTemplate)
	if err != nil {
		return err
	}

	// Create posts directory if it doesn't exist
	postsDir := filepath.Join(outputDir, "posts")
	if err := os.MkdirAll(postsDir, 0755); err != nil {
		return err
	}

	outputFile := filepath.Join(postsDir, "index.html")
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

// updateWebsiteData updates the website data structure according to the workflow
func updateWebsiteData(posts []Post) WebsiteData {
	if len(posts) == 0 {
		return WebsiteData{}
	}

	// First post becomes featured
	featuredPost := posts[0]
	
	// Next 3 posts become recent posts
	var recentPosts []Post
	if len(posts) > 1 {
		end := 4
		if len(posts) < 4 {
			end = len(posts)
		}
		recentPosts = posts[1:end]
	}

	return WebsiteData{
		FeaturedPost: featuredPost,
		RecentPosts:  recentPosts,
		AllPosts:     posts,
	}
}

// savePostsAsJSON saves all posts as JSON
func savePostsAsJSON(posts []Post, outputFile string) error {
	data, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputFile, data, 0644)
}

func main() {
	outputDir := "."
	postsFile := "posts.json"
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run website-generator.go <posts.json>")
		fmt.Println("Example: go run website-generator.go posts.json")
		os.Exit(1)
	}
	
	postsFile = os.Args[1]
	
	// Load posts from JSON file
	posts, err := loadPostsFromJSON(postsFile)
	if err != nil {
		log.Fatal("Error loading posts from JSON:", err)
	}

	// Sort posts by date (newest first)
	sort.Slice(posts, func(i, j int) bool {
		dateI, errI := parseDate(posts[i].Date)
		dateJ, errJ := parseDate(posts[j].Date)
		
		if errI != nil || errJ != nil {
			return false
		}
		
		return dateI.After(dateJ)
	})

	// Update website data according to workflow
	websiteData := updateWebsiteData(posts)

	// Generate individual post HTML files
	for _, post := range posts {
		if err := generatePostHTML(post, outputDir); err != nil {
			log.Printf("Error generating HTML for %s: %v", post.Title, err)
		}
	}

	// Generate main index.html
	if err := generateIndexHTML(websiteData, outputDir); err != nil {
		log.Fatal("Error generating index.html:", err)
	}

	// Generate posts/index.html
	if err := generatePostsIndexHTML(websiteData, outputDir); err != nil {
		log.Fatal("Error generating posts/index.html:", err)
	}

	fmt.Printf("Website generated successfully!\n")
	fmt.Printf("- Featured post: %s\n", websiteData.FeaturedPost.Title)
	fmt.Printf("- Recent posts: %d\n", len(websiteData.RecentPosts))
	fmt.Printf("- Total posts: %d\n", len(websiteData.AllPosts))
} 