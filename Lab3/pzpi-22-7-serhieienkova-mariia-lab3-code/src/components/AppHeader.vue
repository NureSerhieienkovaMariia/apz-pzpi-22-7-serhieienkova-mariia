<template>
  <header class="app-header">
    <div class="left">
      <router-link to="/" class="nav-link">{{ $t('clinic') }}</router-link>
    </div>

    <div class="center" v-if="user">
      <router-link to="/patients" class="nav-link">{{ $t('patients') }}</router-link>
      <router-link to="/medicines" class="nav-link">{{ $t('medicines') }}</router-link>
      <router-link to="/diagnoses" class="nav-link">{{ $t('diagnoses') }}</router-link>
      <router-link to="/visits" class="nav-link">{{ $t('visits') }}</router-link>
      <router-link to="/notifications" class="nav-link">{{ $t('notifications') }}</router-link>
    </div>

    <div class="right">
      <select v-model="locale" @change="changeLanguage">
        <option value="ua">UA</option>
        <option value="en">EN</option>
      </select>

      <template v-if="user">
        <span class="greeting">
          {{ user.name }} {{ user.surname }}
        </span>
        <button @click="handleLogout" class="logout-button">{{ $t('logout') }}</button>
      </template>
    </div>
  </header>
</template>

<script setup>
import { useI18n } from 'vue-i18n'
import { useUser } from '../composables/useUser'
import { useRouter } from 'vue-router'

const router = useRouter()
const { user, clearUser } = useUser()

const { locale } = useI18n()

const changeLanguage = () => {
  localStorage.setItem('language', locale.value)
}

const handleLogout = () => {
  clearUser()
  localStorage.removeItem('accessToken')
  localStorage.removeItem('refreshToken')
  router.push('/login')
}
</script>

<style scoped>
.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: #4caf50;
  color: white;
}

.nav-link {
  margin: 0 10px;
  color: white;
  text-decoration: none;
}

.nav-link:hover {
  text-decoration: underline;
}

.greeting {
  margin-right: 10px;
}

.logout-button {
  padding: 4px 8px;
  background-color: #c14f4a;
  border: none;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

.logout-button:hover {
  background-color: #9c2626;
}
</style>
