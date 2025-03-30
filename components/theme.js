// Apply theme immediately to prevent flash
const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
    document.documentElement.setAttribute('data-theme', savedTheme);
}

function toggleTheme() {
    const root = document.documentElement;
    const currentTheme = root.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? '' : 'dark';

    root.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
}

// Apply saved theme on page load (in case it was changed in another tab)
document.addEventListener('DOMContentLoaded', () => {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
        document.documentElement.setAttribute('data-theme', savedTheme);
    }
});