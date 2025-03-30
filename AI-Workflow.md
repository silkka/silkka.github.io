# Design

- Pages should follow WCAG 2.1 AA
- General idea for the theme: Traffic Sign Theme - Information Blue with Accessibility Improvements

# Rules for blogposts

- Always use date format Feb 24, 2021 and always include the year
- User will provide a blogpost with structure md/template.md
- Use that data to create a blogpost and use templates/post.html as the base
- User will tell which category the post belongs
    - Place the blogposts under that category
    - Update the index.html in that category

# Front Page Structure

- Featured Post Section
    - Display the latest post with a large highlighted image
    - Show full title, date, and a longer excerpt
    - Include a "READ MORE" link with proper htmx attributes

- Recent Posts Section
    - Display exactly three recent posts in a grid layout
    - Each post card should include:
        - Featured image
        - Date
        - Title
        - Short excerpt
        - "READ MORE" link with htmx attributes

- Post Update Flow
    - When adding a new post:
        1. Make it the featured post at the top
        2. Move the current featured post to the first position in recent posts
        3. Shift existing recent posts down
        4. Remove the oldest post from recent posts
    - Update posts/index.html to include the new post in the full list