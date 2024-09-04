import { writable } from "svelte/store";

type Theme = "light" | "dark";

const defaultTheme: Theme = "dark";

// Initialize the store with the saved theme or the default
const initialTheme: Theme =
  (localStorage.getItem("theme") as Theme | null) || defaultTheme;

export const theme = writable<Theme>(initialTheme);

// Function to apply theme
function applyTheme(newTheme: Theme) {
  document.documentElement.classList.toggle("dark-theme", newTheme === "dark");
  localStorage.setItem("theme", newTheme);
}

// Subscribe to changes and update HTML class and localStorage
theme.subscribe(applyTheme);

export function toggleTheme(): void {
  theme.update((currentTheme) => (currentTheme === "light" ? "dark" : "light"));
}
