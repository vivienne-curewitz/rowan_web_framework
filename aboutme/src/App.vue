<script setup lang="ts">
// No additional script logic needed for this static page
</script>

<template>
  <div class="about-me">
    <header class="header">
      <div class="profile-container">
        <img src="/IMG_2624.jpeg" alt="Profile Picture" class="profile-img" />
        <div class="header-text">
          <h1>[Your Name]</h1>
          <p class="tagline">[Professional Title | e.g., Senior Software Engineer]</p>
        </div>
      </div>
    </header>

    <main class="content">
      <section class="summary-section">
        <h2>Professional Summary</h2>
        <p>
          [A brief, high-level overview of your professional background, core strengths, and what
          drives you. This is your "elevator pitch" to potential employers or collaborators.]
        </p>
      </section>

      <section class="capabilities-section">
        <h2>Professional Capabilities</h2>
        <ul class="capabilities-list">
          <li>
            <strong>[Capability 1]</strong>: [Brief description of your expertise and impact in this
            area.]
          </li>
          <li>
            <strong>[Capability 2]</strong>: [Brief description of your expertise and impact in this
            area.]
          </li>
          <li>
            <strong>[Capability 3]</strong>: [Brief description of your expertise and impact in this
            area.]
          </li>
          <li>
            <strong>[Capability 4]</strong>: [Brief description of your expertise and impact in this
            area.]
          </li>
        </ul>
      </section>

      <section class="experience-section">
        <h2>Key Achievements</h2>
        <ul class="achievements-list">
          <li>[Achievement 1: e.g., Led the migration of a legacy system to a modern microservices architecture.]</li>
          <li>[Achievement 2: e.g., Developed a high-performance framework used by 50+ developers.]</li>
          <li>[Achievement 3: e.g., Optimized database queries, reducing latency by 40%.]</li>
        </ul>
      </section>
    </main>

    <footer class="footer">
      <div class="links">
        <a href="https://github.com/[your-username]" target="_blank" rel="noopener" class="link">
          GitHub
        </a>
        <span class="separator">|</span>
        <a href="https://linkedin.com/in/[your-username]" target="_blank" rel="noopener" class="link">
          LinkedIn
        </a>
      </div>
      <p class="copyright">&copy; {{ new Date().getFullYear() }} [Your Name]</p>
    </footer>
  </div>
</template>

<style>
body {
  margin: 0;
  background-image: url('/rowan_tree.jpg');
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  background-repeat: no-repeat;
  min-height: 100vh;
}
</style>

<style scoped>
:root {
  --primary-color: #1b4332;
  --secondary-color: #2d6a4f;
  --accent-color: #40916c;
  --text-color: #1b4332;
  --light-bg: rgba(248, 249, 250, 0.9);
  --white: rgba(255, 255, 255, 0.95);
  --shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
}

.about-me {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  color: #1b4332;
  line-height: 1.6;
  max-width: 800px;
  margin: 2rem auto;
  padding: 3rem;
  background-color: rgba(255, 255, 255, 0.92);
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(4px);
}

.header {
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 2px solid #e0e0e0;
}

.profile-container {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.profile-img {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  object-fit: cover;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15);
  border: 4px solid #ffffff;
}

.header-text h1 {
  margin: 0;
  font-size: 2.5rem;
  color: #1b4332;
}

.tagline {
  margin: 0.5rem 0 0;
  font-size: 1.2rem;
  color: #2d6a4f;
  font-weight: 500;
}

.content section {
  margin-bottom: 2.5rem;
}

h2 {
  font-size: 1.5rem;
  color: #1b4332;
  border-left: 4px solid #2d6a4f;
  padding-left: 1rem;
  margin-bottom: 1rem;
}

.capabilities-list, .achievements-list {
  list-style-type: none;
  padding: 0;
}

.capabilities-list li, .achievements-list li {
  margin-bottom: 1rem;
  padding: 1rem;
  background-color: rgba(248, 249, 250, 0.7);
  border-radius: 8px;
  transition: transform 0.2s;
  border-left: 2px solid transparent;
}

.capabilities-list li:hover {
  transform: translateX(5px);
  border-left: 2px solid #2d6a4f;
}

.footer {
  margin-top: 4rem;
  padding-top: 2rem;
  border-top: 2px solid #e0e0e0;
  text-align: center;
}

.links {
  margin-bottom: 1rem;
}

.link {
  color: #2d6a4f;
  text-decoration: none;
  font-weight: 600;
  transition: color 0.2s;
}

.link:hover {
  color: #1b4332;
}

.separator {
  margin: 0 1rem;
  color: #bdc3c7;
}

.copyright {
  font-size: 0.9rem;
  color: #707070;
}

@media (max-width: 600px) {
  .profile-container {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }

  .header-text h1 {
    font-size: 2rem;
  }

  .about-me {
    padding: 1.5rem;
    margin: 1rem;
  }
}
</style>
