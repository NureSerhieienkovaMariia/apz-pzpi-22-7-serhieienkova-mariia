<template>
  <div class="form-container">
    <h2>{{ $t('register') }}</h2>
    <form @submit.prevent="register">
      <div>
        <label>{{ $t('name') }}</label>
        <input v-model="name" required />
      </div>
      <div>
        <label>{{ $t('surname') }}</label>
        <input v-model="surname" required />
      </div>
      <div>
        <label>{{ $t('email') }}</label>
        <input v-model="email" type="email" required />
      </div>
      <div>
        <label>{{ $t('password') }}</label>
        <input v-model="password" type="password" required />
      </div>
      <button type="submit">{{ $t('registerButton') }}</button>
    </form>
    <p>
      {{ $t('login') }}?
      <router-link to="/login">{{ $t('login') }}</router-link>
    </p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { signUp } from '../api/auth'
import { useRouter } from 'vue-router'

const router = useRouter()

const name = ref('')
const surname = ref('')
const email = ref('')
const password = ref('')

const register = async () => {
  try {
    const response = await signUp({
      name: name.value,
      surname: surname.value,
      email: email.value,
      password: password.value,
    })
    console.log('Успішна реєстрація:', response)

    router.push('/login')
  } catch (err) {
    console.error(err)
    alert('Помилка реєстрації')
  }
}
</script>
