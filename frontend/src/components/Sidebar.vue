<script setup lang="ts">
</script>

<template>
  <aside>
    <div class="logo">
      <span class="material-icons">science</span>
    </div>

    <h3>Menu</h3>
    <div class="menu">
      <router-link class="button" to="/">
        <span class="material-icons">factory</span>
        <span class="text">{{ $t('menu.process') }}</span>
      </router-link>
      <router-link class="button" to="/phases">
        <span class="material-icons">settings</span>
        <span class="text">{{ $t('menu.phases') }}</span>
      </router-link>
      <router-link class="button" to="/wifi">
        <span class="material-icons">wifi</span>
        <span class="text">{{ $t('menu.wifi') }}</span>
      </router-link>
      <router-link class="button" to="/pt100">
        <span class="material-icons">thermostat</span>
        <span class="text">{{ $t('menu.pt100') }}</span>
      </router-link>
      <router-link class="button" to="/ds">
        <span class="material-icons">thermostat</span>
        <span class="text">{{ $t('menu.ds') }}</span>
      </router-link>
      <router-link class="button" to="/heaters">
        <span class="material-icons">microwave</span>
        <span class="text">{{ $t('menu.heaters') }}</span>
      </router-link>
      <router-link class="button" to="/outputs">
        <span class="material-icons">output</span>
        <span class="text">{{ $t('menu.outputs') }}</span>
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

  color: var(--el-color-white);
  background-color: var(--sidebar-dark);

  .logo {
    margin-bottom: 1rem;
    margin-left: auto;
    margin-right: auto;

    .material-icons {
      font-size: 64px;
    }
  }

  h3,
  .button .text {
    opacity: 1;
  }

  h3 {
    color: var(--el-color-info);
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
        color: var(--el-color-white);
      }

      .text {
        color: var(--el-color-white);
      }

      &:hover,
      &.router-link-exact-active {
        background-color: var(--sidebar-dark-alt);

        .material-icons,
        .text {
          color: var(--el-color-primary);
        }
      }

      &.router-link-exact-active {
        border-right: 5px solid var(--el-color-primary);
      }
    }
  }

  @media (max-width: 768px) {
    position: fixed;
    z-index: 99;
  }
}
</style>
