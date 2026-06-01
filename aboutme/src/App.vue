<script setup lang="ts">
// No additional script logic needed for this static page
</script>

<template>
  <div class="about-me">
    <header class="header">
      <div class="profile-container">
        <img src="/IMG_2624.jpeg" alt="Profile Picture" class="profile-img" />
        <div class="header-text">
          <h1>Vivienne Curewitz</h1>
          <p class="tagline">Senior Software Engineer</p>
        </div>
      </div>
    </header>

    <main class="content">
      <section class="summary-section">
        <h2>Professional Summary</h2>
        <p>
          I am an energetic and experienced Software Engineer with a passion for reliablility and
          sustainability. I aim to build a more sustainable future by creating software that is
          efficient, scalable, and environmentally conscious. With over 10 years of classical
          software development experience and constant pursuit of cutting edge agentic AI tools and
          techniques, I am well suited to help you move your organization into the future.
        </p>
      </section>

      <section class="capabilities-section">
        <h2>Professional Capabilities</h2>
        <ul class="capabilities-list">
          <li>
            <strong>Technical Leadership</strong>: Leadership is more than tracking KPIs, task
            decomposition, and begging your workers to use AI. Leadership is developing a
            relationship with your team and your customer; understanding needs, capabilities, and
            motivations, and using that understanding to achieve mutually beneficial outcomes.
          </li>
          <li>
            <strong>Software Architecture and Design</strong>: Software is not code. It is a living
            and evolving system that exists in the context of developers, users, and hardware. Good
            software is maintainable and extensible while being fast and efficient. It is intuitive
            to use but enables complex controls and interactions. It is a journey, and design is
            guiding that journey towards new features without compromising what already exists.
          </li>
          <li>
            <strong>DevOps and Infrastructure</strong>: I design my software to be lightweight and
            efficient, and so I must design my infrastructure to be equally efficient. I have
            experience with Azure, AWS, and on-prem systems, and I am always looking for ways to
            stretch compute further and reduce costs.
          </li>
        </ul>
      </section>

    </main>

    <footer class="footer">
      <div class="links">
        <a href="https://github.com/vivienne-curewitz"" target="_blank" rel="noopener" class="link">
          GitHub
        </a>
        <span class="separator">|</span>
        <a href="https://www.linkedin.com/in/vivienne-curewitz/" target="_blank" rel="noopener" class="link">
          LinkedIn
        </a>
      </div>
      <p class="copyright">&copy; {{ new Date().getFullYear() }} Vivienne Curewitz</p>
    </footer>
  </div>
</template>

<style>
body {
  margin: 0;
  background-image: url('/birch-tree-clipart-lg.png');
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  background-repeat: no-repeat;
  min-height: 100vh;
  background-color: #f0f4f0; /* Light green-grey fallback */
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
