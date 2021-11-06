/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  serverRuntimeConfig: {
    APPOLO_URI: process.env.SERVER_APOLLO_URI || 'http://localhost:3080',
  },
  publicRuntimeConfig: {
    APOLLO_URI: process.env.PUBLIC_APOLLO_URI || 'http://localhost:3080',
  },
  images: {
    domains: [
      "placeimg.com",
    ]
  }
}
