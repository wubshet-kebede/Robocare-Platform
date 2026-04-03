<template>
  <div class="min-h-screen flex bg-slate-50">
    <div class="hidden lg:block w-1/2 relative overflow-hidden bg-slate-950">
      <video
        autoplay
        loop
        muted
        playsinline
        class="absolute inset-0 w-full h-full object-cover opacity-60"
      >
        <!-- <source src="/assets/robot-animation.mp4" type="video/mp4" /> -->
      </video>
      <div
        class="absolute inset-0 bg-gradient-to-tr from-teal-900/80 via-transparent to-teal-800/20"
      ></div>
      <div
        class="relative z-10 h-full flex flex-col justify-end p-20 text-white"
      >
        <h1
          class="text-7xl font-extrabold tracking-tighter leading-none italic"
        >
          ROBO<span class="text-teal-400">CARE.</span>
        </h1>
        <p
          class="text-xl text-slate-300 font-light max-w-sm mt-4 leading-relaxed"
        >
          Advanced autonomous logistics for the
          <span class="text-white font-medium"
            >University of Gondar Hospital.</span
          >
        </p>
        <div
          class="mt-12 flex gap-8 opacity-50 text-[10px] font-mono uppercase tracking-widest"
        >
          <div>Gondar, ET</div>
          <div>Lat: 12.6000° N</div>
          <div>Long: 37.4667° E</div>
        </div>
      </div>
    </div>

    <div
      class="w-full lg:w-1/2 flex items-center justify-center p-12 lg:p-24 bg-white"
    >
      <div class="w-full max-w-lg">
        <div class="mb-16">
          <h2 class="text-4xl font-extrabold text-slate-950 tracking-tight">
            Welcome Back!
          </h2>
          <p class="mt-4 text-base text-slate-600 leading-relaxed max-w-md">
            Sign in to access your dashboard and continue
            <span class="font-medium text-teal-700"
              >managing your hospital fleet.</span
            >
          </p>
        </div>

        <form class="space-y-8" @submit.prevent="handleLogin">
          <div class="space-y-2.5">
            <label
              for="email"
              class="block text-sm font-semibold text-slate-800"
              >Email</label
            >
            <div class="relative">
              <div
                class="absolute inset-y-0 left-0 pl-4.5 flex items-center pointer-events-none"
              >
                <svg
                  class="h-5 w-5 text-teal-600"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.206"
                  />
                </svg>
              </div>
              <input
                v-model="form.email"
                id="email"
                type="email"
                required
                class="block w-full pl-13 pr-4 py-3.5 bg-white border border-slate-200 rounded-xl text-sm placeholder-slate-400 focus:outline-none focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20 transition duration-150"
                placeholder="doctor@hospital.com"
              />
            </div>
          </div>

          <div class="space-y-2.5">
            <label
              for="password"
              class="block text-sm font-semibold text-slate-800"
              >Password</label
            >
            <div class="relative">
              <div
                class="absolute inset-y-0 left-0 pl-4.5 flex items-center pointer-events-none"
              >
                <svg
                  class="h-5 w-5 text-teal-600"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                  />
                </svg>
              </div>
              <input
                v-model="form.password"
                id="password"
                :type="passwordVisible ? 'text' : 'password'"
                required
                class="block w-full pl-13 pr-12 py-3.5 bg-white border border-slate-200 rounded-xl text-sm placeholder-slate-400 focus:outline-none focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20 transition duration-150"
                placeholder="Enter your password"
              />
              <button
                type="button"
                @click="passwordVisible = !passwordVisible"
                class="absolute inset-y-0 right-0 pr-4 flex items-center text-slate-400 hover:text-teal-600 transition"
              >
                <svg
                  v-if="passwordVisible"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L5.136 5.136M14.122 14.122l4.742 4.742"
                  />
                </svg>
                <svg
                  v-else
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                  />
                </svg>
              </button>
            </div>
            <div class="text-right">
              <a
                href="#"
                class="text-xs font-semibold text-teal-700 hover:text-teal-600 transition"
                >Forgot Password?</a
              >
            </div>
          </div>

          <button
            type="submit"
            :disabled="loading"
            class="w-full flex justify-center py-4 px-4 text-base font-semibold rounded-xl text-white bg-teal-950 hover:bg-teal-900 transition duration-150 disabled:opacity-50"
          >
            <span v-if="loading">Signing in...</span>
            <span v-else>Sign In</span>
          </button>
        </form>

        <div class="relative my-12">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-slate-100"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-3 bg-white text-slate-400 font-medium">OR</span>
          </div>
        </div>

        <div class="space-y-4">
          <button
            class="w-full flex items-center justify-center gap-3 py-3 px-4 border border-slate-200 rounded-xl text-sm font-medium text-slate-800 hover:bg-slate-50 transition"
          >
            <!-- <img src="/icons/google.svg" class="h-5 w-5" alt="Google" /> -->
            <Icon name="devicon:google" class="h-5 w-5" />
            Continue with Google
          </button>
          <button
            class="w-full flex items-center justify-center gap-3 py-3 px-4 border border-slate-200 rounded-xl text-sm font-medium text-slate-800 hover:bg-slate-50 transition"
          >
            <!-- <img src="/icons/apple.svg" class="h-5 w-5" alt="Apple" /> Continue -->
            <Icon name="ic:outline-apple" class="h-5 w-5" />
            with Apple
          </button>
        </div>

        <p class="mt-14 text-center text-sm text-slate-600">
          Don't have an Account?
          <a
            href="#"
            class="font-semibold text-teal-700 hover:text-teal-600 transition"
            >Sign Up</a
          >
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
const form = reactive({ email: "", password: "" });
const loading = ref(false);
const passwordVisible = ref(false);
definePageMeta({
  layout: false,
});

const handleLogin = async () => {
  loading.value = true;
  console.log("Submitting:", form.email);
  setTimeout(() => {
    loading.value = false;
  }, 1000);
};
</script>
<!-- <script setup>
definePageMeta({
  layout: false,
});
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
</script> -->
