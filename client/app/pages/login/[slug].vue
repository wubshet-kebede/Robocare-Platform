<template>
  <div class="min-h-screen flex bg-slate-50">
    <div
      class="hidden lg:flex w-1/2 bg-blue-700 relative overflow-hidden items-center justify-center"
    >
      <div class="z-10 text-white p-12 max-w-lg">
        <h1 class="text-5xl font-bold mb-6 italic">Robocare.</h1>
        <p class="text-xl text-blue-100 leading-relaxed">
          The future of hospital automation. Managing your fleet of medical
          robots starts here.
        </p>
      </div>
      <div
        class="absolute -bottom-24 -left-24 w-96 h-96 bg-blue-600 rounded-full mix-blend-multiply filter blur-3xl opacity-70"
      ></div>
      <div
        class="absolute -top-24 -right-24 w-96 h-96 bg-blue-500 rounded-full mix-blend-multiply filter blur-3xl opacity-70"
      ></div>
    </div>

    <div class="w-full lg:w-1/2 flex items-center justify-center p-8">
      <div class="max-w-md w-full space-y-8">
        <div class="text-center">
          <img class="mx-auto h-12 w-auto" src="/logo.png" alt="Robocare" />
          <h2 class="mt-6 text-3xl font-extrabold text-slate-900">
            Welcome back
          </h2>
          <p class="mt-2 text-sm text-slate-600">
            Please enter your details to sign in.
          </p>
        </div>

        <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
          <div class="space-y-4">
            <div>
              <label
                for="email"
                class="block text-sm font-medium text-slate-700"
                >Email Address</label
              >
              <input
                v-model="form.email"
                id="email"
                type="email"
                required
                class="mt-1 block w-full px-4 py-3 bg-white border border-slate-300 rounded-xl text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition duration-200"
                placeholder="doctor@hospital.com"
              />
            </div>

            <div>
              <div class="flex items-center justify-between">
                <label
                  for="password"
                  class="block text-sm font-medium text-slate-700"
                  >Password</label
                >
                <a
                  href="#"
                  class="text-xs font-semibold text-blue-600 hover:text-blue-500"
                  >Forgot password?</a
                >
              </div>
              <input
                v-model="form.password"
                id="password"
                type="password"
                required
                class="mt-1 block w-full px-4 py-3 bg-white border border-slate-300 rounded-xl text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition duration-200"
                placeholder="••••••••"
              />
            </div>
          </div>

          <div class="flex items-center">
            <input
              id="remember-me"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-slate-300 rounded"
            />
            <label for="remember-me" class="ml-2 block text-sm text-slate-700"
              >Keep me logged in</label
            >
          </div>

          <button
            type="submit"
            :disabled="loading"
            class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-semibold rounded-xl text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition duration-200 disabled:opacity-50"
          >
            <span v-if="loading">Signing in...</span>
            <span v-else>Sign In</span>
          </button>
        </form>

        <p class="mt-10 text-center text-sm text-slate-500">
          Not a member yet?
          <span class="font-semibold text-slate-900 italic"
            >Contact your administrator</span
          >
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
const form = reactive({
  email: "",
  password: "",
});
const loading = ref(false);

const handleLogin = async () => {
  loading.value = true;
  // Call your Go Backend here!
  console.log("Logging in with:", form.email);
  setTimeout(() => {
    loading.value = false;
  }, 1000);
};
</script>
