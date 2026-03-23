<template>
  <section id="contact" class="py-24 bg-slate-950 border-t border-slate-900">
    <div class="container mx-auto px-6">
      <div class="grid lg:grid-cols-2 gap-16 items-start">
        <!-- Left Side: Contact Info & Branding -->
        <div class="space-y-8">
          <div>
            <h2
              class="text-cyan-500 font-black text-sm uppercase tracking-[0.3em] mb-4"
            >
              Get In Touch
            </h2>
            <h3 class="text-4xl font-bold text-white mb-6 tracking-tight">
              Inquire About the <br />
              <span class="text-slate-500 font-display italic"
                >Framework Deployment.</span
              >
            </h3>
            <p class="text-slate-400 text-lg leading-relaxed max-w-md">
              Are you interested in testing our IoT-SLAM framework in your
              medical facility? Reach out to our research team at the University
              of Gondar.
            </p>
          </div>

          <div class="space-y-6">
            <div class="flex items-center gap-4 group">
              <div
                class="w-12 h-12 bg-slate-900 border border-slate-800 rounded-xl flex items-center justify-center group-hover:border-cyan-500/50 transition-colors"
              >
                <Icon name="mdi:location" class="text-xl text-cyan-500" />
              </div>
              <div>
                <div class="text-white font-bold text-sm">Location</div>
                <div class="text-slate-500 text-sm font-medium">
                  Institute of Technology, Gondar, Ethiopia
                </div>
              </div>
            </div>

            <div class="flex items-center gap-4 group">
              <div
                class="w-12 h-12 bg-slate-900 border border-slate-800 rounded-xl flex items-center justify-center group-hover:border-emerald-500/50 transition-colors"
              >
                <Icon
                  name="ic:baseline-email"
                  class="text-xl text-emerald-500"
                />
              </div>
              <div>
                <div class="text-white font-bold text-sm">Email</div>
                <div class="text-slate-500 text-sm font-medium">
                  wubshet2829@gmail.com
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Side: Professional Form -->
        <div
          class="bg-slate-900/40 border border-slate-800 p-8 md:p-10 rounded-[2.5rem] backdrop-blur-xl shadow-2xl"
        >
          <form @submit.prevent class="space-y-6">
            <div class="grid md:grid-cols-2 gap-6">
              <div class="space-y-2">
                <label
                  class="text-[10px] font-black uppercase tracking-widest text-slate-500 ml-1"
                  >Full Name</label
                >
                <input
                  type="text"
                  placeholder="John Doe"
                  class="w-full bg-slate-950 border border-slate-800 rounded-xl px-5 py-4 text-white text-sm focus:border-cyan-500 outline-none transition-all"
                />
              </div>
              <div class="space-y-2">
                <label
                  class="text-[10px] font-black uppercase tracking-widest text-slate-500 ml-1"
                  >Email Address</label
                >
                <input
                  type="email"
                  placeholder="john@hospital.com"
                  class="w-full bg-slate-950 border border-slate-800 rounded-xl px-5 py-4 text-white text-sm focus:border-cyan-500 outline-none transition-all"
                />
              </div>
            </div>

            <div class="space-y-2">
              <label
                class="text-[10px] font-black uppercase tracking-widest text-slate-500 ml-1"
                >Inquiry Type</label
              >
              <select
                class="w-full bg-slate-950 border border-slate-800 rounded-xl px-5 py-4 text-slate-400 text-sm focus:border-cyan-500 outline-none transition-all appearance-none"
              >
                <option>Deployment Request</option>
                <option>Technical Collaboration</option>
                <option>Academic Inquiry</option>
              </select>
            </div>

            <div class="space-y-2">
              <label
                class="text-[10px] font-black uppercase tracking-widest text-slate-500 ml-1"
                >Message</label
              >
              <textarea
                rows="4"
                placeholder="How can we help your facility?"
                class="w-full bg-slate-950 border border-slate-800 rounded-xl px-5 py-4 text-white text-sm focus:border-cyan-500 outline-none transition-all resize-none"
              ></textarea>
            </div>

            <button
              class="w-full py-4 bg-gradient-to-r from-cyan-600 to-emerald-600 hover:from-cyan-500 hover:to-emerald-500 text-slate-950 font-black rounded-xl transition-all transform hover:scale-[1.02] shadow-xl shadow-cyan-500/10"
            >
              SEND MESSAGE
            </button>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
