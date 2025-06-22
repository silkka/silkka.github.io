package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

// getCurrentDate returns the current date in the format "Jan 2, 2006"
func getCurrentDate() string {
	return time.Now().Format("Jan 2, 2006")
}

// promptForInput prompts the user for input and returns the response
func promptForInput(prompt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", prompt)
	if defaultValue != "" {
		fmt.Printf(" (default: %s): ", defaultValue)
	} else {
		fmt.Print(": ")
	}
	
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	if input == "" && defaultValue != "" {
		return defaultValue
	}
	return input
}

// promptForMultilineInput prompts for multiline input until user enters "END"
func promptForMultilineInput(prompt string) []string {
	fmt.Println(prompt)
	fmt.Println("Enter each link on a new line. Type 'END' when finished:")
	
	var links []string
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "END" {
			break
		}
		
		if input != "" {
			links = append(links, input)
		}
	}
	
	return links
}

// loadPostsFromJSON loads posts from JSON file
func loadPostsFromJSON(filename string) ([]Post, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty slice
			return []Post{}, nil
		}
		return nil, err
	}

	var posts []Post
	err = json.Unmarshal(data, &posts)
	return posts, err
}

// savePostsToJSON saves posts to JSON file
func savePostsToJSON(posts []Post, filename string) error {
	data, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// addPostToFront adds a new post to the front of the posts slice
func addPostToFront(posts []Post, newPost Post) []Post {
	return append([]Post{newPost}, posts...)
}

func main() {
	postsFile := "posts.json"
	
	fmt.Println("🌐 Add New Post")
	fmt.Println("==============")
	
	// Load existing posts
	posts, err := loadPostsFromJSON(postsFile)
	if err != nil {
		log.Fatal("Error loading posts:", err)
	}
	
	fmt.Printf("Found %d existing posts\n\n", len(posts))
	
	// Get post information interactively
	title := promptForInput("Post title", "")
	if title == "" {
		fmt.Println("❌ Title is required!")
		os.Exit(1)
	}
	
	description := promptForInput("Description", "")
	if description == "" {
		fmt.Println("❌ Description is required!")
		os.Exit(1)
	}
	
	topic := promptForInput("Topic (Maps/Games/Posts)", "Maps")
	date := promptForInput("Date", getCurrentDate())
	imageURL := promptForInput("Image URL", "")
	
	if imageURL == "" {
		fmt.Println("❌ Image URL is required!")
		os.Exit(1)
	}
	
	// Get external links
	fmt.Println("\nExternal links:")
	externalLinks := promptForMultilineInput("Enter external links (Steam Workshop, GitHub, etc.)")
	
	// Create new post
	newPost := Post{
		Title:        title,
		Description:  description,
		Topic:        topic,
		Date:         date,
		ImageURL:     imageURL,
		ExternalLinks: externalLinks,
		Slug:         generateSlug(title),
		Category:     determineCategory(topic),
		Content:      "",
	}
	
	// Add to front of posts (will become featured)
	posts = addPostToFront(posts, newPost)
	
	// Save updated posts
	if err := savePostsToJSON(posts, postsFile); err != nil {
		log.Fatal("Error saving posts:", err)
	}
	
	fmt.Printf("\n✅ Post added successfully!\n")
	fmt.Printf("📝 Title: %s\n", newPost.Title)
	fmt.Printf("📅 Date: %s\n", newPost.Date)
	fmt.Printf("🏷️  Category: %s\n", newPost.Category)
	fmt.Printf("🔗 Slug: %s\n", newPost.Slug)
	fmt.Printf("📊 Total posts: %d\n", len(posts))
	
	fmt.Printf("\n🚀 To generate the website, run:\n")
	fmt.Printf("   go run website-generator.go %s\n", postsFile)
} 