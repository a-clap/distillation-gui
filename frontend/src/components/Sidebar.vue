<script setup>
import {ref} from "vue";

const is_expanded = ref(false)
const ToggleMenu = () => {
    is_expanded.value = !is_expanded.value
}
</script>

<template>
    <aside>
        <div class="logo">
            <img src="../assets/images/logo-universal.png" alt="Vue">
        </div>

        <h3>Menu</h3>
        <div class="menu">
            <router-link class="button" to="/">
                <span class="material-icons">home</span>
                <span class="text">Home</span>
            </router-link>
            <router-link class="button" to="/about">
                <span class="material-icons">settings</span>
                <span class="text">About</span>
            </router-link>

        </div>
    </aside>
</template>

<style lang="scss" scoped>

aside {
  display: flex;
  flex-direction: column;
  width: var(--sidebar-width);
  overflow: hidden;
  min-height: 100vh;
  padding: 1rem;

  color: var(--light);
  background-color: var(--dark);

  .logo {
    margin-bottom: 1rem;
    margin-left: auto;
    margin-right: auto;

    img {
      width: 2rem;
    }
  }

  h3, .button .text {
    opacity: 1;
  }

  h3 {
    color: var(--grey);
    font-size: 0.875rem;
    margin-bottom: 0.5rem;
    margin-left: auto;
    margin-right: auto;
    text-transform: uppercase;
  }

  .menu {
    margin: 0 -1rem;

    .button {
      display: flex;
      align-items: center;
      text-decoration: none;

      padding: 0.5rem 1rem;

      .material-icons {
        font-size: 2rem;
        margin-right: 1rem;
        color: var(--light);
      }

      .text {
        color: var(--light);
      }

      &:hover, &.router-link-exact-active {
        background-color: var(--dark-alt);

        .material-icons, .text {
          color: var(--primary);
        }
      }

      &.router-link-exact-active {
        border-right: 5px solid var(--primary);
      }
    }
  }

  @media (max-width: 768px) {
    position: fixed;
    z-index: 99;
  }
}
</style>
