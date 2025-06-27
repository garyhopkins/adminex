// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  typescript: {
    typeCheck: true
  },

  nitro: {
    routeRules: {
        '/api/**': {
            proxy: {
                to: 'http://localhost:8888/api/**' // make sure this is an ENV driven variable if production does not match
            }
        }
    }
},
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  ssr: false,
  ui: {
	  colorMode: false,
  },

  modules: [
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/test-utils',
    '@nuxt/ui'
  ]
})
