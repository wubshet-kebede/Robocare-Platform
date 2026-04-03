<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-6">
    <!-- Backdrop -->
    <div
      class="absolute inset-0 bg-slate-950/80 backdrop-blur-md"
      @click="$emit('close')"
    ></div>

    <!-- Modal Content -->
    <div
      class="relative w-full max-w-lg bg-slate-900 border border-white/10 rounded-[3rem] p-12 shadow-2xl overflow-hidden"
    >
      <!-- Glow -->
      <div
        class="absolute -top-20 -right-20 w-64 h-64 bg-cyan-500/10 blur-[100px] rounded-full"
      ></div>

      <div class="relative z-10 space-y-8">
        <div class="text-center">
          <h3
            class="text-2xl font-black text-white uppercase tracking-tighter italic"
          >
            Generate Staff Access
          </h3>
          <p
            class="text-slate-500 text-[10px] font-bold uppercase tracking-[0.3em] mt-2"
          >
            Multi-Tenant Invitation Link
          </p>
        </div>

        <div class="space-y-6">
          <UiBaseInput
            v-model="form.name"
            label="Full Name"
            placeholder="Dr. Jane Doe"
            icon="lucide:user"
          />
          <UiBaseInput
            v-model="form.email"
            label="Institutional Email"
            placeholder="jane@uog.edu.et"
            icon="lucide:mail"
          />

          <div class="space-y-2">
            <label
              class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-500 ml-1"
              >Assign Access Role</label
            >
            <select
              v-model="form.role"
              class="w-full bg-slate-950/50 border border-white/5 rounded-2xl px-6 py-4 text-slate-400 text-sm focus:border-cyan-500/50 outline-none transition-all appearance-none cursor-pointer"
            >
              <option>Doctor</option>
              <option>Nurse</option>
              <option>Receptionist</option>
            </select>
          </div>
        </div>

        <div class="flex gap-4 pt-4">
          <UiBaseButton
            class="flex-1"
            variant="secondary"
            @click="$emit('close')"
            >Cancel</UiBaseButton
          >
          <UiBaseButton class="flex-1" variant="gradient" @click="submit"
            >Transmit Invite</UiBaseButton
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(["close", "send"]);
const form = reactive({
  name: "",
  email: "",
  role: "Doctor",
});

const submit = () => {
  if (form.name && form.email) {
    emit("send", { ...form });
  }
};
</script>
