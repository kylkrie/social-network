<!-- src/lib/components/LandingPage.svelte -->
<script lang="ts">
  import { ArrowRight } from "lucide-svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Card from "$lib/components/ui/Card.svelte";
  import { startAuthLogin, startAuthRegister } from "$lib/stores";
  import ThemeToggle from "$lib/components/layout/ThemeToggle.svelte";

  function handleLogin() {
    startAuthLogin();
  }

  function handleSignUp() {
    startAuthRegister();
  }
</script>

<div
  class="flex flex-col min-h-screen bg-gradient-to-b from-background to-surface text-text"
>
  <header class="container mx-auto px-4 py-8">
    <div class="flex justify-between items-center">
      <div class="w-10"></div>
      <!-- Spacer to balance the ThemeToggle -->
      <h1 class="text-4xl font-bold text-center text-primary flex-grow">
        Yabro Social
      </h1>
      <div class="w-10 flex justify-end">
        <ThemeToggle />
      </div>
    </div>
    <p class="text-xl text-center mt-4 text-text-secondary">
      A Portfolio Project by Kyle Smith
    </p>
  </header>

  <main class="flex-grow container mx-auto px-4 py-2">
    <section class="text-center mb-12">
      <div class="flex justify-center space-x-4">
        <Button variant="default" size="lg" on:click={handleLogin}>
          Log In <ArrowRight class="ml-2 h-4 w-4" />
        </Button>
        <Button variant="outline" size="lg" on:click={handleSignUp}>
          Sign Up
        </Button>
      </div>
    </section>
    <section class="mb-12 text-center">
      <h2 class="text-3xl font-semibold mb-4 text-text">About Me</h2>
      <p class="text-xl text-text-secondary mb-4">
        Hi, I'm Kyle Smith, a passionate full-stack developer with expertise in
        modern web technologies. This project showcases my skills in building
        scalable and user-friendly web applications.
      </p>
    </section>

    <section class="mb-12 text-center">
      <h2 class="text-3xl font-semibold mb-4 text-text">Portfolio Showcase</h2>
      <p class="text-xl text-text-secondary mb-4">
        Yabro Social is a comprehensive social media platform built as a
        portfolio piece to demonstrate my proficiency in full-stack development,
        including:
      </p>
      <ul
        class="list-disc list-inside text-left max-w-md mx-auto text-text-secondary"
      >
        <li>Frontend development with Svelte and TailwindCSS</li>
        <li>Backend API design and implementation</li>
        <li>User authentication and authorization</li>
        <li>Real-time updates and notifications</li>
        <li>Responsive and accessible UI design</li>
      </ul>
    </section>

    <section class="mb-12">
      <h2 class="text-3xl font-semibold mb-4 text-center text-text">
        Key Features
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <Card title="User Profiles">
          <p class="text-text-secondary">
            Create and customize your profile, share your interests, and connect
            with others.
          </p>
        </Card>
        <Card title="Post Creation">
          <p class="text-text-secondary">
            Share your thoughts, images, and links with your network through
            easy-to-create posts.
          </p>
        </Card>
        <Card title="Interactions">
          <p class="text-text-secondary">
            Like, comment, and share posts to engage with content and other
            users.
          </p>
        </Card>
        <Card title="Direct Messaging">
          <p class="text-text-secondary">
            Communicate privately with other users through a secure messaging
            system.
          </p>
        </Card>
        <Card title="Notifications">
          <p class="text-text-secondary">
            Stay updated with real-time notifications for interactions and
            messages.
          </p>
        </Card>
        <Card title="Search & Discovery">
          <p class="text-text-secondary">
            Find new connections and interesting content through powerful search
            and discovery tools.
          </p>
        </Card>
      </div>
    </section>
  </main>

  <footer class="bg-primary text-button-text py-8">
    <div class="container mx-auto px-4 text-center">
      <p>&copy; 2024 Kyle Smith. All rights reserved.</p>
      <p class="mt-2">
        <a
          href="https://github.com/kylkrie/social-network"
          target="_blank"
          rel="noopener noreferrer"
          class="underline hover:text-primary-light"
        >
          View this project on GitHub
        </a>
      </p>
    </div>
  </footer>
</div>
