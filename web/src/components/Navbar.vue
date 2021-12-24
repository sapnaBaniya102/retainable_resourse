<script setup lang="ts">
import { isDark, toggleDark, collapsed, toggleSidebar, navbarWidth } from '~/composables'
const { t, availableLocales, locale } = useI18n()

const toggleLocales = () => {
  // change to some real logic
  const locales = availableLocales
  locale.value = locales[(locales.indexOf(locale.value) + 1) % locales.length]
}
</script>
<template>
  <div class="float-right sticky" :style="{ width: navbarWidth }">
    <button class="icon-btn mx-2 !outline-none float-right bg-gray-300 rounded-full h-10 w-13 mt-2" :title="t('button.toggle_dark')" @click="toggleDark()">
      <carbon-moon v-if="isDark" />
      <carbon-sun v-else />
    </button>
    <router-link to="/" active-class="active" tag="button">
      <div class="link-container text-2xl float-right mt-2 pt-1 pl-3 mr-2 bg-gray-300 rounded-full h-10 w-13 items-center justify-center">
        <carbon-search /> <span class="text-black" />
      </div>
    </router-link>
    <input class="float-right  px-2 py-2 mt-2 mr-2 w-60 border rounded-md focus:ring" type="text" name="text" placeholder="Search" required><br>
    <button class="collapse-icon !outline-none icon-btn ml-2 -mt-7 bg-gray-300 rounded-full h-10 w-13 text-1xl" :class="{ 'rotate-180': collapsed }" @click="toggleSidebar">
      <carbon-logout />
    </button>
  </div>
</template>
