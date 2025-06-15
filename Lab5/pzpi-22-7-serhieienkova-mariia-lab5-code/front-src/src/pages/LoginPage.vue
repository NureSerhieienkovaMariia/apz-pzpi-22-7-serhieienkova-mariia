<template>
  <div class="form-container">
    <h2>{{ $t('login') }}</h2>

    <div v-if="error" class="error">{{ error }}</div>

    <input v-model="email" :placeholder="$t('email')" />
    <input v-model="password" type="password" :placeholder="$t('password')" />

    <button @click="handleLogin">{{ $t('login') }}</button>

    <p>
      <router-link to="/register">{{ $t('register') }}</router-link>
    </p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { signIn, getCurrentUser } from '../api/auth'
import { useUser } from '../composables/useUser'

const router = useRouter()
const { setUser } = useUser()

const email = ref('')
const password = ref('')
const error = ref(null)

const handleLogin = async () => {
  error.value = null
  try {
    const data = await signIn({ email: email.value, password: password.value })
    localStorage.setItem('accessToken', data.access_jwt_token)
    localStorage.setItem('refreshToken', data.refresh_jwt_token)

    const currentUser = await getCurrentUser(data.access_jwt_token)
    setUser(currentUser)

    localStorage.setItem('user', JSON.stringify({
      accessToken: data.access_jwt_token,
      userId: data.user_id,
      userType: data.user_type
    }))

    router.push('/patients')
  } catch (err) {
    console.error(err)
    error.value = 'Помилка авторизації: ' + err.message
  }
}
</script>

<style scoped>
.form-container {
  max-width: 400px;
  margin: 30px auto;
  background-color: white;
  padding: 20px 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

input {
  display: block;
  margin-bottom: 12px;
  padding: 8px;
  width: 100%;
}

button {
  padding: 8px 12px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  width: 100%;
}

button:hover {
  background-color: #45a049;
}

.error {
  color: red;
  margin-bottom: 10px;
  text-align: center;
}
</style>
